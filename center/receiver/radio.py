
from . import cc1101

class Radio(cc1101.CC1101):
    async def calibrate(self):
        self.wcmd(self.CommandStrobe.SCAL)
        await self.waitstate(1)

    async def setup_basic(self):
        await self.reset()

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

    async def setup_for_rx(self):
        # fix packet length
        self.set(self.ConfReg.PKTLEN, 16)

        # packet automation
        self.set(self.ConfReg.PKTCTRL0, 0x44)

        self.set(self.ConfReg.PKTCTRL1, 0x2c)

        # main radio control state machine configuration
        self.set(self.ConfReg.MCSM1, 0x3c)
        self.set(self.ConfReg.MCSM0, 0x34)

        self.set(self.ConfReg.IOCFG0, 0x07)

        await self.calibrate()

    async def setup_for_conf(self):
        # treat 0 as broadcast
        self.set(self.ConfReg.PKTLEN, 0x10)
        self.set(self.ConfReg.PKTCTRL1, 0x0a)
        self.set(self.ConfReg.PKTCTRL0, 0x45)
        self.set(self.ConfReg.IOCFG0, 0x07)
        self.set(self.ConfReg.MCSM1, 0x0b)
        self.set(self.ConfReg.ADDR, 0)

        await self.calibrate()
