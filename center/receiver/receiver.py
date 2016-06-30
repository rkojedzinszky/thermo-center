
import struct
import re
import logging
from Crypto.Cipher import AES
from django.conf import settings
from twisted.internet import reactor
from center.receiver import RadioBase, radio, sensorvalue
from center import models
from center.carbon import PickleClient

logger = logging.getLogger(__name__)

class StringOrdIter(object):
    def __init__(self, s):
        self._i = iter(s)

    def next(self):
        return ord(next(self._i))

class Receiver(RadioBase):
    def __del__(self):
        print 'Receiver destroyed'

    def run(self):
        print 'Receiver.run()'

        self._cc = PickleClient(settings.CARBON_PICKLE_ENDPOINT)

        self._aes = AES.new(''.join(chr(int(c, base=16)) for c in re.findall(r'[0-9a-f]{2}', self._config.aes_key)))

        self._radio.setup_basic()
        self._radio.xfer2(self._config.config_bytes())
        self._radio.setup_for_rx()
        self._radio.wcmd(radio.Radio.CommandStrobe.SRX)

        self.enable_interrupt()

    def oninterrupt(self):
        data_len = self._radio.status(radio.Radio.StatusReg.RXBYTES)
        data = self._radio.read_rxfifo(data_len)
        if data_len != 18:
            return

        try:
            self._receive_packet(data)
        except Exception as e:
            logger.error('error processing packet: %s' % str(e))

    def _receive_packet(self, packet):
        metrics = [sensorvalue.RSSI(packet[16]), sensorvalue.LQI(packet[17] & 0x7f)]
        packet = self._aes.decrypt(''.join(chr(c) for c in packet[:16]))

        network, seq, length, id_ = struct.unpack('<HLBB', packet[:8])
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

        si = StringOrdIter(rest)
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
            models.Sensor.objects.get(id=id_).feed(seq, metrics, carbon=self._cc)
        except models.Sensor.DoesNotExist:
            logger.warn('Unknown device id: %02x' % id_)
