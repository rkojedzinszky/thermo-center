
import logging
from twisted.internet import reactor
from django.utils import timezone
from center.receiver import RadioBase, radio
from center.models import Sensor

logger = logging.getLogger(__name__)

class Configurator(RadioBase):
    def _prepare_config_reply(self):
        self._configpacket = [
                54, # total length
                0,  # reply address
                # configuration is from here
                0,  # crc
                0,  # id
                self._config.network_id & 0xff, # network id lsb
                self._config.network_id >> 8    # network id msb
                ]
        self._configpacket.extend(self._config.aes_bytes())
        self._configpacket.extend(self._config.config_bytes())
        self._configpacket.extend([0xff] * (self._configpacket[0] - len(self._configpacket)))
        self._configpacket[0] -= 1 # fixup for total length

    def _gen_next_id(self):
        self._next_id = (set(range(1, 128)) - set([s.pk for s in Sensor.objects.all()])).pop()

    def run(self):
        self._prepare_config_reply()
        self._gen_next_id()

        self._radio.setup_basic()
        self._radio.setup_for_conf()
        self.enable_interrupt()

        self._radio.wcmd(radio.Radio.CommandStrobe.SRX)

        logger.debug('Configurator initialized')
        reactor.callLater(15, self._main.startreceiver)

    def oninterrupt(self):
        data_len = self._radio.status(radio.Radio.StatusReg.RXBYTES)
        data = self._radio.read_rxfifo(data_len)
        if data_len != 3 or data[0] != 2 or data[1] != 0:
            # re-enter RX mode
            self._radio.sidle()
            self._radio.wcmd(radio.Radio.CommandStrobe.SFRX)
            self._radio.wcmd(radio.Radio.CommandStrobe.SRX)
            return

        id_ = data[2]
        if id_ & 0x80:
            id_ = self._next_id

        logger.info('Received autoconfig request from %02x, sending reply with id=%02x' % (data[2], id_))

        self._configpacket[1] = data[2]
        self._configpacket[3] = id_

        self._radio.write_txfifo(self._configpacket)

        try:
            sensor = Sensor.objects.get(id=id_)
        except Sensor.DoesNotExist:
            sensor = Sensor(id=id_)

        sensor.last_ts = timezone.now()
        sensor.save()

        self._gen_next_id()
