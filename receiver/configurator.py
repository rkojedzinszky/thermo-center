
import asyncio
import logging
from receiver import (
        base, radio
)
from configurator import api_pb2 as cfg_pb2

DISCOVERY_PACKET_TIMEOUT = 22

logger = logging.getLogger(__name__)


class Configurator(base.Base):
    name = 'configurator'

    def __init__(self, args, radio, task):
        super().__init__(args=args, radio=radio)

        self.task = task

    def _prepare_config_reply(self, sensor_id, config):
        packet = [
                54, # total length
                0,  # reply address
                # configuration is from here
                0,  # crc
                sensor_id,  # id
                config.network & 0xff, # network id lsb
                config.network >> 8    # network id msb
                ]
        packet.extend(config.aes_key)
        packet.extend(config.radio_config)
        packet.extend([0xff] * (packet[0] - len(packet)))
        packet[0] -= 1 # fixup for total length

        return packet

    async def _task_acquire(self):
        return await self.loop.run_in_executor(None, self.configurator.TaskAcquire, self.task)

    async def _task_discovery_received(self):
        await self.loop.run_in_executor(None, self.configurator.TaskDiscoveryReceived, self.task)

    async def _task_finished(self, error=None):
        await self.loop.run_in_executor(None,
                self.configurator.TaskFinished,
                cfg_pb2.TaskFinishedRequest(task_id=self.task.task_id, error=error)
        )

    async def arun(self):
        task = await self._task_acquire()

        replypacket = self._prepare_config_reply(task.sensor_id, task.config)

        await self.radio.setup_basic()
        await self.radio.setup_for_conf()

        deadline = self.loop.time() + DISCOVERY_PACKET_TIMEOUT

        seen = None
        while True:
            # Wait for a packet with timeout
            try:
                sensor_id = await asyncio.wait_for(self._wait_discovery_packet(), timeout=deadline - self.loop.time())
            except asyncio.TimeoutError:
                break

            if sensor_id & 128:
                logger.info('Received discovery packet from %d', sensor_id)
            else:
                logger.info('Received reconfiguration request from %d', sensor_id)

            if (sensor_id & 128) == 0 and sensor_id != task.sensor_id:
                logging.warning('Received unexpected reconfiguration discovery packet')
                continue

            if seen is None:
                replypacket[1] = sensor_id
                seen = sensor_id

            if seen != sensor_id:
                logging.warning('Ignoring unexpected discovery packet from %d, expecting %d', sensor_id, seen)
                continue

            await self._send_replypacket(replypacket)
            logging.info('Replied to %d', seen)

            # Just do logging now
            await self._task_discovery_received()

            # Wait another DISCOVERY_PACKET_TIMEOUT to check if the
            # sensor has received our reply, thus not sending discovery
            # packets
            deadline = self.loop.time() + DISCOVERY_PACKET_TIMEOUT

        error = None
        if seen is None:
            error = 'No discovery received'

        await self._task_finished(error=error)

        logger.info('Exiting discovery loop')

    async def _wait_discovery_packet(self):
        while True:
            self.radio.wcmd(radio.Radio.CommandStrobe.SRX)
            await self.gpio.waitforinterrupt()

            data_len = self.radio.status(radio.Radio.StatusReg.RXBYTES)
            data = self.radio.read_rxfifo(data_len)
            if data_len != 3 or data[0] != 2 or data[1] != 0:
                self.radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                continue

            return data[2]

    async def _send_replypacket(self, packet):
            self.radio.write_txfifo(packet)
            await self.radio.wait_sidle()  # wait until it is sent
