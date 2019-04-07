
import struct
import re
import logging
import time
import asyncio
from django.db import connection
from django.db.utils import OperationalError
from Crypto.Cipher import AES
from django.conf import settings
from django.core import exceptions
from center.receiver import RadioBase, radio, sensorvalue
from center import models
from center.carbon import LineClient, PickleClient
from center.receiver import pid
from center.receiver.tbf import TokenBucketFilter

PIDCONTROL_INT_ABS_MAX = 100.0

WATCHDOG_TIMEOUT = 300

logger = logging.getLogger(__name__)

class PIDControlValue(sensorvalue.Value):
    metric = 'pidcontrol'

class Receiver(RadioBase):
    name = 'receiver'

    class InterruptStorm(RuntimeError):
        pass

    async def _setup_radio(self):
        await self._radio.setup_basic()
        self._radio.xfer2(self._config.config_bytes())
        await self._radio.setup_for_rx()
        self._radio.wcmd(radio.Radio.CommandStrobe.SRX)
        self._tbf.reset()

    async def main(self):
        self._ccs = [] # Carbon cache clients
        self._aes = AES.new(bytes([int(c, base=16) for c in re.findall(r'[0-9a-f]{2}', self._config.aes_key)]))
        self._tbf = TokenBucketFilter(settings.INTERRUPT_MAX_RATE, settings.INTERRUPT_MAX_BURST, self.loop.time)
        await self._setup_radio()

        if settings.CARBON_LINE_RECEIVER_ENDPOINT[0]:
            self._ccs.append(LineClient(self.loop, settings.CARBON_LINE_RECEIVER_ENDPOINT, maxsize=settings.CARBON_QUEUE_MAXSIZE))

        if settings.CARBON_PICKLE_RECEIVER_ENDPOINT[0]:
            self._ccs.append(PickleClient(self.loop, settings.CARBON_PICKLE_RECEIVER_ENDPOINT, maxsize=settings.CARBON_QUEUE_MAXSIZE))

        for cc in self._ccs:
            cc.start()

        while True:
            try:
                packets = await asyncio.wait_for(self.receive_many(), timeout=WATCHDOG_TIMEOUT)
            except asyncio.TimeoutError:
                logger.warn('Watchdog timeout, resetting radio')
                self._setup_radio()
            except Receiver.InterruptStorm:
                logger.warn('Interrupt storm detected, resetting radio')
                await self._setup_radio()
            else:
                for packet in packets:
                    self.receive(packet)

    async def stop(self):
        for cc in self._ccs:
            cc.stop()

        await super().stop()

    async def receive_many(self):
        packets = []
        while len(packets) == 0:
            await self.waitforinterrupt()

            logger.debug('INT TBF Capacity: {}'.format(self._tbf.capacity))
            if not self._tbf.feed(1):
                raise Receiver.InterruptStorm

            data_len = self._radio.status(radio.Radio.StatusReg.RXBYTES)
            logger.debug('CC1101.RXBYTES=%d' % data_len)

            if data_len & 0x80:
                logger.warn('CC1101 RX_OVERFLOW')
                self._radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                await self._radio.wait_sidle()
                self._radio.wcmd(radio.Radio.CommandStrobe.SRX)
                continue

            if data_len > 54:
                logger.warn('CC1101 suspicious RXBYTES')
                await self._radio.sidle()
                self._radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                self._radio.wcmd(radio.Radio.CommandStrobe.SRX)
                continue

            while data_len >= 18:
                # read one packet from radio
                packets.append(self._radio.read_rxfifo(18))
                data_len -= 18

        return packets

    def receive(self, data):
        start = time.time()

        try:
            self._receive_packet(data)
        except OperationalError as e:
            logger.error('database error during processing packet: %s' % str(e))
            connection.close()
        except Exception as e:
            logger.error('error processing packet: %s' % str(e))

        end = time.time()

        logger.debug('Packet processed in %f seconds' % (end - start))

    def _receive_packet(self, packet):
        metrics = [sensorvalue.RSSI(packet[16]), sensorvalue.LQI(packet[17] & 0x7f)]
        packet = self._aes.decrypt(bytes(packet[:16]))

        network, seq, length, id_ = struct.unpack('<HLBB', packet[:8])

        logger.debug('packet header: network=%04x seq=%08x len=%02x id=%02x' % (network, seq, length, id_))

        if length < 8:
            logger.error('Invalid packet data received, short length: %d' % length)
            return
        if length > 16:
            logger.error('Invalid packet data received, large length: %d' % length)
            return

        rest = packet[8:length]

        if network != self._config.network_id:
            logger.warn('Received packet for invalid network: %d' % network)
            return

        si = iter(rest)
        try:
            while True:
                try:
                    metrics.append(sensorvalue.SensorValueParser.parse(next(si), next(si)))
                except sensorvalue.SensorValueParser.InvalidType:
                    pass
        except StopIteration:
            pass

        mh = {m.metric: m for m in metrics}
        try:
            t = mh['Temperature']
            h = mh['Humidity']
            if isinstance(t, sensorvalue.HTU21DTemperature) and isinstance(h, sensorvalue.HTU21DHumidty):
                h.temp_compensate(t)
        except KeyError:
            pass

        try:
            s = models.Sensor.objects.select_related('control').get(id=id_)
        except models.Sensor.DoesNotExist:
            logger.warn('Unknown device id: %02x' % id_)
            return

        if hasattr(s, 'control'):
            pcp = s.control

            try:
                pc = self._pidmap[id_]
            except KeyError:
                pc = self._pidmap.setdefault(id_, pid.PID())

            target_temp = pcp.get_target_temp()
            if target_temp is not None:
                error = target_temp - mh['Temperature'].value()
                pc.feed(error=error, intabsmax=PIDCONTROL_INT_ABS_MAX)
                pcv = pc.value(kp=pcp.kp, ki=pcp.ki, kd=pcp.kd)
                logger.debug('%s: pid control=%f', s, pcv)
                metrics.append(PIDControlValue(pcv))

        s.feed(seq, metrics, carbons=self._ccs, mqtt=self._mqtt)
