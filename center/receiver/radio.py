
import cc1101

class Radio(cc1101.CC1101):
    def calibrate(self):
        self.wcmd(self.CommandStrobe.SCAL)
        self.waitstate(1)

    def setup_basic(self):
        self.reset()

        # frequency configuration
        self.set(self.ConfReg.FREQ2, 0x10)
        self.set(self.ConfReg.FREQ1, 0xa7)
        self.set(self.ConfReg.FREQ0, 0xe1)

        # modem configuration
        self.set(self.ConfReg.MDMCFG4, 0xaa)
        self.set(self.ConfReg.MDMCFG3, 0xbc)
        self.set(self.ConfReg.MDMCFG2, 0x71)
        self.set(self.ConfReg.DEVIATN, 0x40)

        # PATABLE
        self.set(self.ConfReg.PATABLE, 0xc0)

        self.calibrate()

    def setup_common(self):
        # fix packet length
        self.set(self.ConfReg.PKTLEN, 16)

        # packet automation
        self.set(self.ConfReg.PKTCTRL0, 0x44)

    def setup_for_rx(self):
        self.setup_common()

        self.set(self.ConfReg.PKTCTRL1, 0x2c)

        # main radio control state machine configuration
        self.set(self.ConfReg.MCSM1, 0x3c)
        self.set(self.ConfReg.MCSM0, 0x34)

        self.set(self.ConfReg.IOCFG0, 0x07)

        self.calibrate()

    def setup_for_tx(self):
        self.setup_common()

        # main radio control state machine configuration
        self.set(self.ConfReg.MCSM1, 0x30)
        self.set(self.ConfReg.MCSM0, 0x38)

        self.calibrate()
