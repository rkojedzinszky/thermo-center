
import asyncio
import logging
from receiver import (
        base, radio
)

DISCOVERY_PACKET_TIMEOUT = 22

logger = logging.getLogger(__name__)


class Configurator(base.Base):
    name = 'configurator'

    def __init__(self, args, radio, params):
        super().__init__(args=args, radio=radio)

        self.params = params

    def _prepare_config_reply(self, config):
        packet = [
                54, # total length
                0,  # reply address
                # configuration is from here
                0,  # crc
                0,  # id
                config.network & 0xff, # network id lsb
                config.network >> 8    # network id msb
                ]
        packet.extend(config.aes_key)
        packet.extend(config.radio_config)
        packet.extend([0xff] * (packet[0] - len(packet)))
        packet[0] -= 1 # fixup for total length

        return packet

    async def arun(self):
        print ('Configuring sensor {}'.format(self.params.sensor_id))
        config = self._read_config()

        replypacket = self._prepare_config_reply(config)

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

            logger.info('Received discovery packet from %d', sensor_id)

            if seen is None:
                if self.params.sensor_id == -1:  # only existing
                    if sensor_id < 128:
                        replypacket[3] = sensor_id
                    else:
                        logging.warning('Received unexpected initial discovery packet')
                        continue
                else:
                    if sensor_id < 128:
                        logging.warning('Received unexpected reconfiguration discovery packet')
                        continue
                    else:
                        replypacket[3] = self.params.sensor_id

                replypacket[1] = sensor_id
                seen = sensor_id

            if seen != sensor_id:
                logging.warning('Ignoring unexpected discovery packet from %d, expecting %d', sensor_id, seen)
                continue

            await self._send_replypacket(replypacket)
            logging.info('Replied to %d', seen)

            # Wait another DISCOVERY_PACKET_TIMEOUT to check if the
            # sensor has received our reply, thus not sending discovery
            # packets
            deadline = self.loop.time() + DISCOVERY_PACKET_TIMEOUT

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
