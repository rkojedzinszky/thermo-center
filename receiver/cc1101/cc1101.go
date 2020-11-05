package cc1101

import (
	"fmt"
	"time"

	"periph.io/x/periph/conn/spi"
)

// ConfReg enumerates Configuration registers
type ConfReg uint8

// ConfReg values
const (
	IOCFG2   ConfReg = 0x00
	IOCFG1   ConfReg = 0x01
	IOCFG0   ConfReg = 0x02
	FIFOTHR  ConfReg = 0x03
	SYNC1    ConfReg = 0x04
	SYNC0    ConfReg = 0x05
	PKTLEN   ConfReg = 0x06
	PKTCTRL1 ConfReg = 0x07
	PKTCTRL0 ConfReg = 0x08
	ADDR     ConfReg = 0x09
	CHANNR   ConfReg = 0x0A
	FSCTRL1  ConfReg = 0x0B
	FSCTRL0  ConfReg = 0x0C
	FREQ2    ConfReg = 0x0D
	FREQ1    ConfReg = 0x0E
	FREQ0    ConfReg = 0x0F
	MDMCFG4  ConfReg = 0x10
	MDMCFG3  ConfReg = 0x11
	MDMCFG2  ConfReg = 0x12
	MDMCFG1  ConfReg = 0x13
	MDMCFG0  ConfReg = 0x14
	DEVIATN  ConfReg = 0x15
	MCSM2    ConfReg = 0x16
	MCSM1    ConfReg = 0x17
	MCSM0    ConfReg = 0x18
	FOCCFG   ConfReg = 0x19
	BSCFG    ConfReg = 0x1A
	AGCCTRL2 ConfReg = 0x1B
	AGCCTRL1 ConfReg = 0x1C
	AGCCTRL0 ConfReg = 0x1D
	WOREVT1  ConfReg = 0x1E
	WOREVT0  ConfReg = 0x1F
	WORCTRL  ConfReg = 0x20
	FREND1   ConfReg = 0x21
	FREND0   ConfReg = 0x22
	FSCAL3   ConfReg = 0x23
	FSCAL2   ConfReg = 0x24
	FSCAL1   ConfReg = 0x25
	FSCAL0   ConfReg = 0x26
	RCCTRL1  ConfReg = 0x27
	RCCTRL0  ConfReg = 0x28
	FSTEST   ConfReg = 0x29
	PTEST    ConfReg = 0x2A
	AGCTEST  ConfReg = 0x2B
	TEST2    ConfReg = 0x2C
	TEST1    ConfReg = 0x2D
	TEST0    ConfReg = 0x2E
	PATABLE  ConfReg = 0x3E
)

// CommandStrobe enumerates commands
type CommandStrobe uint8

// CommandStrobe values
const (
	SRES    CommandStrobe = 0x30
	SFSTXON CommandStrobe = 0x31
	SXOFF   CommandStrobe = 0x32
	SCAL    CommandStrobe = 0x33
	SRX     CommandStrobe = 0x34
	STX     CommandStrobe = 0x35
	SIDLE   CommandStrobe = 0x36
	SWOR    CommandStrobe = 0x38
	SPWD    CommandStrobe = 0x39
	SFRX    CommandStrobe = 0x3A
	SFTX    CommandStrobe = 0x3B
	SWORRST CommandStrobe = 0x3C
	SNOP    CommandStrobe = 0x3D
)

// StatusReg enumerares Status registers
type StatusReg uint8

// StatusReg values
const (
	PARTNUM        StatusReg = 0x30
	VERSION        StatusReg = 0x31
	FREQEST        StatusReg = 0x32
	LQI            StatusReg = 0x33
	RSSI           StatusReg = 0x34
	MARCSTATE      StatusReg = 0x35
	WORTIME1       StatusReg = 0x36
	WORTIME0       StatusReg = 0x37
	PKTSTATUS      StatusReg = 0x38
	VCO_VC_DAC     StatusReg = 0x39
	TXBYTES        StatusReg = 0x3A
	RXBYTES        StatusReg = 0x3B
	RCCTRL1_STATUS StatusReg = 0x3C
	RCCTRL0_STATUS StatusReg = 0x3D
)

// CC1101 handles a CC1101 device
type CC1101 struct {
	spi spi.Conn
}

// New creates a new CC1101 instance
func New(spi spi.Conn) *CC1101 {
	return &CC1101{
		spi: spi,
	}
}

// Xfer transfers the packets and keeps chipselect low
func (cc *CC1101) Xfer(tx []uint8) (rx []uint8, err error) {
	rx = make([]byte, len(tx))
	packet := spi.Packet{
		R:      rx,
		W:      tx,
		KeepCS: true,
	}

	err = cc.spi.TxPackets([]spi.Packet{packet})

	return
}

// SetConf sets a configurarion register
func (cc *CC1101) SetConf(reg ConfReg, value uint8) error {
	_, err := cc.Xfer([]byte{byte(reg), value})

	return err
}

// GetConf reads a configurarion register
func (cc *CC1101) GetConf(reg ConfReg) (uint8, error) {
	rx, err := cc.Xfer([]byte{byte(reg) | 0x80, 0})

	if err != nil {
		return 0, err
	}

	return rx[1], nil
}

// Cmd issues a command strobe
func (cc *CC1101) Cmd(cmd CommandStrobe) error {
	_, err := cc.Xfer([]byte{byte(cmd)})

	return err
}

// GetStatus reads a status register
func (cc *CC1101) GetStatus(sreg StatusReg) (uint8, error) {
	rx, err := cc.Xfer([]byte{byte(sreg) | 0xc0, 0})

	if err != nil {
		return 0, err
	}

	return rx[1], nil
}

// Release releases the CS line
func (cc *CC1101) Release() error {
	return cc.spi.Tx([]byte{byte(SNOP)}, []byte{0})
}

// ReadRXFifo reads len bytes from rxfifo
func (cc *CC1101) ReadRXFifo(len int) ([]byte, error) {
	tx := make([]byte, len+1)
	tx[0] = 0xff

	rx := make([]byte, len+1)

	if err := cc.spi.Tx(tx, rx); err != nil {
		return nil, err
	}

	return rx[1:], nil
}

// WriteTXFifo reads len bytes from rxfifo
func (cc *CC1101) WriteTXFifo(tx []byte) error {
	rtx := make([]byte, len(tx)+1)
	copy(rtx[1:], tx)
	rtx[0] = 0x7f

	rx := make([]byte, len(rtx))

	return cc.spi.Tx(rtx, rx)
}

// Reset resets the CC1101
func (cc *CC1101) Reset() error {
	cmd := make([]byte, 1)

	cmd[0] = byte(SRES)
	if err := cc.spi.Tx(cmd, cmd); err != nil {
		return err
	}

	cmd[0] = byte(SRES)
	if err := cc.spi.Tx(cmd, cmd); err != nil {
		return err
	}

	if err := cc.Waitstate(1); err != nil {
		return err
	}

	cc.SetConf(IOCFG2, 0x2e)
	cc.SetConf(IOCFG0, 0x2e)

	return nil
}

// Waitstate waits for the desired state. It will timeout in 1 second
func (cc *CC1101) Waitstate(desired uint8) error {
	deadline := time.Now().Add(1 * time.Second)

	for {
		state, err := cc.GetStatus(MARCSTATE)
		if err != nil {
			return err
		}
		if state == desired {
			return nil
		}
		if deadline.Before(time.Now()) {
			return fmt.Errorf("cc1101.Waitstate: timeout reached")
		}

		time.Sleep(10 * time.Millisecond)
	}
}
