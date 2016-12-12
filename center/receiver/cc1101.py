
class CC1101(object):
    """ Represents a CC1101 radio """
    class ConfReg(object):
        IOCFG2 = 0x00
        IOCFG1 = 0x01
        IOCFG0 = 0x02
        FIFOTHR = 0x03
        SYNC1 = 0x04
        SYNC0 = 0x05
        PKTLEN = 0x06
        PKTCTRL1 = 0x07
        PKTCTRL0 = 0x08
        ADDR = 0x09
        CHANNR = 0x0A
        FSCTRL1 = 0x0B
        FSCTRL0 = 0x0C
        FREQ2 = 0x0D
        FREQ1 = 0x0E
        FREQ0 = 0x0F
        MDMCFG4 = 0x10
        MDMCFG3 = 0x11
        MDMCFG2 = 0x12
        MDMCFG1 = 0x13
        MDMCFG0 = 0x14
        DEVIATN = 0x15
        MCSM2 = 0x16
        MCSM1 = 0x17
        MCSM0 = 0x18
        FOCCFG = 0x19
        BSCFG = 0x1A
        AGCCTRL2 = 0x1B
        AGCCTRL1 = 0x1C
        AGCCTRL0 = 0x1D
        WOREVT1 = 0x1E
        WOREVT0 = 0x1F
        WORCTRL = 0x20
        FREND1 = 0x21
        FREND0 = 0x22
        FSCAL3 = 0x23
        FSCAL2 = 0x24
        FSCAL1 = 0x25
        FSCAL0 = 0x26
        RCCTRL1 = 0x27
        RCCTRL0 = 0x28
        FSTEST = 0x29
        PTEST = 0x2A
        AGCTEST = 0x2B
        TEST2 = 0x2C
        TEST1 = 0x2D
        TEST0 = 0x2E
        PATABLE = 0x3E

    class CommandStrobe(object):
        SRES = 0x30
        SFSTXON = 0x31
        SXOFF = 0x32
        SCAL = 0x33
        SRX = 0x34
        STX = 0x35
        SIDLE = 0x36
        SWOR = 0x38
        SPWD = 0x39
        SFRX = 0x3A
        SFTX = 0x3B
        SWORRST = 0x3C
        SNOP = 0x3D

    class StatusReg(object):
        PARTNUM = 0x30
        VERSION = 0x31
        FREQEST = 0x32
        LQI = 0x33
        RSSI = 0x34
        MARCSTATE = 0x35
        WORTIME1 = 0x36
        WORTIME0 = 0x37
        PKTSTATUS = 0x38
        VCO_VC_DAC = 0x39
        TXBYTES = 0x3A
        RXBYTES = 0x3B
        RCCTRL1_STATUS = 0x3C
        RCCTRL0_STATUS = 0x3D

    def __init__(self, spidev):
        self._spi = spidev

    def marcstate(self):
        return self.status(self.StatusReg.MARCSTATE) & 0x1f

    def waitstate(self, state):
        while self.marcstate() != state:
            pass

    def reset(self):
        self._spi.xfer([self.CommandStrobe.SRES])
        self.waitstate(1)

        # disable GDO[0,2] pins
        self.set(self.ConfReg.IOCFG2, 0x2e)
        self.set(self.ConfReg.IOCFG0, 0x2e)

    def wait_sidle(self):
        self.waitstate(1)

    def sidle(self):
        self._spi.xfer([self.CommandStrobe.SIDLE])
        self.wait_sidle()

    def get(self, reg):
        return self._spi.xfer2([reg | 0x80, 0])[1]

    def set(self, reg, value):
        self._spi.xfer2([reg, value])

    def xfer2(self, values):
        self._spi.xfer2(values)

    def rcmd(self, cmd):
        self._spi.xfer2([cmd | 0x80])

    def wcmd(self, cmd):
        self._spi.xfer2([cmd])

    def status(self, reg):
        return self._spi.xfer2([reg | 0xc0, 0])[1]

    def release(self):
        self._spi.xfer([self.CommandStrobe.SNOP])

    def read_rxfifo(self, length):
        to = [0xff] + [0] * length
        return self._spi.xfer(to)[1:]

    def write_txfifo(self, data):
        self._spi.xfer([0x7f] + data)
