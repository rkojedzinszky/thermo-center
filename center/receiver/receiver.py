
import struct
import re
import logging
import time
import threading
import Queue
from Crypto.Cipher import AES
from django.conf import settings
from django.db import connection
from django.core import exceptions
from twisted.internet import reactor
from center.receiver import RadioBase, radio, sensorvalue
from center import models
from center.carbon import PickleClient
from center.receiver import pid

PIDCONTROL_INTERVAL = 10 * 60

logger = logging.getLogger(__name__)

class PIDControlValue(sensorvalue.Value):
    metric = 'pidcontrol'

class StringOrdIter(object):
    def __init__(self, s):
        self._i = iter(s)

    def next(self):
        return ord(next(self._i))

class Receiver(RadioBase):
    name = 'receiver'

    def setpidmap(self, pidmap):
        self._pidmap = pidmap

    def start(self):
        self._cc = PickleClient(settings.CARBON_PICKLE_ENDPOINT)

        self._aes = AES.new(''.join(chr(int(c, base=16)) for c in re.findall(r'[0-9a-f]{2}', self._config.aes_key)))

        self._radio.setup_basic()
        self._radio.xfer2(self._config.config_bytes())
        self._radio.setup_for_rx()
        self._radio.wcmd(radio.Radio.CommandStrobe.SRX)

        self.enable_interrupt()

        self._icnt = 0

        self._q = Queue.Queue()
        self._t = threading.Thread(target=self._receive_thread)
        self._t.start()

    def stop(self):
        self._q.put(None)
        self._t.join()
        pass

    def oninterrupt(self):
        logger.debug('Receiver.oninterrupt (#%d)' % self._icnt)
        self._icnt += 1
        while True:
            data_len = self._radio.status(radio.Radio.StatusReg.RXBYTES)
            logger.debug('CC1101.RXBYTES=%d' % data_len)

            if data_len & 0x80:
                logger.warn('CC1101 RX_OVERFLOW')
                self._radio.wcmd(radio.Radio.CommandStrobe.SFRX)
                self._radio.wcmd(radio.Radio.CommandStrobe.SRX)
                return

            # we read all available full packets
            data_len -= data_len % 18
            if data_len == 0:
                break

            data = self._radio.read_rxfifo(data_len)

            while len(data) > 0:
                p = data[:18]
                data = data[18:]

                self._q.put(p)

    def _receive_thread(self):
        while True:
            p = self._q.get()
            if p is None:
                break

            start = time.time()

            try:
                self._receive_packet(p)
            except Exception as e:
                logger.error('error processing packet: %s' % str(e))

            end = time.time()

            logger.debug('Packet processed in %f seconds' % (end - start))

        connection.close()

    def _receive_packet(self, packet):
        metrics = [sensorvalue.RSSI(packet[16]), sensorvalue.LQI(packet[17] & 0x7f)]
        packet = self._aes.decrypt(''.join(chr(c) for c in packet[:16]))

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
            s = models.Sensor.objects.select_related('pidcontrolparams').get(id=id_)
        except models.Sensor.DoesNotExist:
            logger.warn('Unknown device id: %02x' % id_)
            return

        try:
            pc = self._pidmap[id_]
        except KeyError:
            pc = self._pidmap.setdefault(id_, pid.PID(PIDCONTROL_INTERVAL))

        pc.feed(mh['Temperature'].value())
        target_temp = s.get_target_temp()
        if target_temp is not None:
            try:
                pcp = s.pidcontrolparams
            except exceptions.ObjectDoesNotExist:
                kp, ki, kd = 1.0, 1.0, 1.0
            else:
                kp, ki, kd = pcp.kp, pcp.ki, pcp.kd

            pcv = pc.value(target_temp, kp=kp, ki=ki, kd=kd)
            logger.debug('%s: pid control=%f', s, pcv)
            metrics.append(PIDControlValue(pcv))

        s.feed(seq, metrics, carbon=self._cc)
