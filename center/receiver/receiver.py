
import struct
import re
from Crypto.Cipher import AES
from twisted.internet import reactor
from center.receiver import RadioBase, radio
from center import models

class Receiver(RadioBase):
    def __del__(self):
        print 'Receiver destroyed'

    def run(self):
        print 'Receiver.run()'
        reactor.callLater(10, self._main.startconfigurator)

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

        self._receive_packet(data)

    def _receive_packet(self, packet):
        rssi = (struct.unpack('b', struct.pack('B', packet[16]))[0] - 148 ) / 2.0
        lqi = packet[17] & 127
        packet = self._aes.decrypt(''.join(chr(c) for c in packet[:16]))

        network, seq, length, id_ = struct.unpack('<HLBB', packet[:8])
        rest = packet[8:length]

        print 'network=%d, devid=%d, length=%d, seq=%d, rest=%s' % (network, id_, length, seq, " ".join("{:02x}".format(ord(c)) for c in rest))
        print ' RSSI=%d LQI=%d' % (rssi, lqi)
