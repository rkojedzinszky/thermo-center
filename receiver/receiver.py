
import asyncio
import grpc
import logging
import struct
import json
import time
from receiver import (
        base, radio, config
)
from aggregator import api_pb2 as agg_msg, api_pb2_grpc as agg_grpc

WATCHDOG_TIMEOUT = 300

logger = logging.getLogger(__name__)


class Receiver(base.Base):
    name = 'receiver'

    def _create_astub(self):
        """ Creates stub to aggregator """
        channel = grpc.insecure_channel('{}:{}'.format(self.args.aggregator_host, self.args.aggregator_port),
                                        (
                    ('grpc.keepalive_time_ms', 10000),
                    ('grpc.keepalive_timeout_ms', 1000),
                )
            )

        return agg_grpc.AggregatorStub(channel)

    async def _setup_radio(self):
        await self.radio.setup_basic()
        self.radio.xfer2(self.config.radio_config)
        await self.radio.setup_for_rx()
        self.radio.wcmd(radio.Radio.CommandStrobe.SRX)

    # This handles one cycle:
    # 1. resets, configures radio
    # 2. Enters receiving loop
    async def arun(self):
        self.config = await self.loop.run_in_executor(None, self._read_config)
        self.astub = await self.loop.run_in_executor(None, self._create_astub)

        await self._setup_radio()
        self.gpio.resettbf()

        while True:
            try:
                packets = await asyncio.wait_for(self.receive_many(), timeout=WATCHDOG_TIMEOUT)
            except asyncio.TimeoutError:
                logger.warning('Watchdog timeout, resetting radio')
                await self._setup_radio()
            else:
                for packet in packets:
                    await self.receive(packet)

    async def receive_many(self):
        packets = []
        while len(packets) == 0:
            await self.gpio.waitforinterrupt()

            data_len = self.radio.status(radio.Radio.StatusReg.RXBYTES)
            logger.debug('CC1101.RXBYTES=%d', data_len)

            if data_len & 0x80:
                logger.warning('CC1101 RX_OVERFLOW')
                self.radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                await self.radio.wait_sidle()
                self.radio.wcmd(radio.Radio.CommandStrobe.SRX)
                continue

            if data_len > 54:
                logger.warning('CC1101 suspicious RXBYTES')
                await self.radio.sidle()
                self.radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                self.radio.wcmd(radio.Radio.CommandStrobe.SRX)
                continue

            while data_len >= 18:
                # read one packet from radio
                packets.append(self.radio.read_rxfifo(18))
                data_len -= 18

        return packets

    async def receive(self, data):
        start = time.time()

        try:
            await self._receive_packet(data)
        except Exception:
            logger.error('error processing packet', exc_info=True)

        end = time.time()

        logger.debug('Packet processed in %f seconds', end - start)

    async def _receive_packet(self, packet):
        rssi = struct.unpack('b', bytes([packet[16]]))[0] / 2.0 - 74
        lqi = packet[17] & 0x7f
        packet = self.config.cipher.decrypt(bytes(packet[:16]))

        network, seq, length, id_ = struct.unpack('<HLBB', packet[:8])

        logger.debug('packet header: network=%04x seq=%08x len=%02x id=%02x',
            network, seq, length, id_)

        if length < 8:
            logger.error(
                'Invalid packet data received, short length: %d', length)
            return
        if length > 16:
            logger.error(
                'Invalid packet data received, large length: %d', length)
            return

        if network != self.config.network:
            logger.warning('Received packet for invalid network: %d', network)
            return

        raw = bytes(packet[8:length])

        sensorpacket = agg_msg.SensorPacket(
                id=id_,
                seq=seq,
                rssi=rssi,
                lqi=lqi,
                raw=raw
                )

        await self.loop.run_in_executor(None, self.astub.FeedSensorPacket, sensorpacket)
