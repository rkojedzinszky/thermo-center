package receiver

import (
	"github.com/rkojedzinszky/thermo-center/v5/receiver/cc1101"
)

type radio struct {
	cc1101 *cc1101.CC1101
}

func (r radio) calibrate() (err error) {
	if err = r.cc1101.Cmd(cc1101.SCAL); err != nil {
		return
	}

	return r.cc1101.Waitstate(1)
}

type configSetting struct {
	confreg cc1101.ConfReg
	value   uint8
}

// Basic radio configuration
var basicSettings = []configSetting{
	// Frequency
	{cc1101.FREQ2, 0x10},
	{cc1101.FREQ1, 0xa7},
	{cc1101.FREQ0, 0xe1},

	// Modem
	{cc1101.MDMCFG4, 0xaa},
	{cc1101.MDMCFG3, 0xbc},
	{cc1101.MDMCFG2, 0x71},
	{cc1101.DEVIATN, 0x40},

	// PATABLE
	{cc1101.PATABLE, 0xc0},
}

var rxSettings = []configSetting{
	// fix packet length
	{cc1101.PKTLEN, 16},

	// packet automation
	{cc1101.PKTCTRL1, 0x2c},
	{cc1101.PKTCTRL0, 0x44},

	// main radio control state machine configuration
	{cc1101.MCSM1, 0x3c},
	{cc1101.MCSM0, 0x34},

	// Interrupt on receive
	{cc1101.IOCFG0, 0x07},
}

var confSettings = []configSetting{
	// fix packet length
	{cc1101.PKTLEN, 16},

	// packet automation, treat 0 as broadcast
	{cc1101.PKTCTRL1, 0x0a},
	{cc1101.PKTCTRL0, 0x45},
	{cc1101.ADDR, 0},

	{cc1101.MCSM1, 0x08},

	// Interrupt on receive
	{cc1101.IOCFG0, 0x07},
}

func (r radio) applySettings(settings []configSetting) error {
	for _, s := range settings {
		if err := r.cc1101.SetConf(s.confreg, s.value); err != nil {
			return err
		}
	}

	return nil
}

func (r radio) setupBasic() error {
	return r.applySettings(basicSettings)
}

func (r radio) setupForRX() error {
	if err := r.applySettings(rxSettings); err != nil {
		return err
	}

	if err := r.calibrate(); err != nil {
		return err
	}

	return r.cc1101.Cmd(cc1101.SRX)
}

func (r radio) setupForConf() error {
	if err := r.applySettings(confSettings); err != nil {
		return err
	}

	return r.calibrate()
}
