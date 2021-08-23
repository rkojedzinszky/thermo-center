package receiver

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"log"
	"time"

	"github.com/rkojedzinszky/thermo-center/aggregator"
	"github.com/rkojedzinszky/thermo-center/configurator"
	"github.com/rkojedzinszky/thermo-center/receiver/cc1101"
)

const (
	watchdogTimeout = 300 * time.Second
)

func (r *Runner) receiverTask() *receiver {
	return &receiver{
		runner: r,
	}
}

type receiver struct {
	runner *Runner
	cfg    *configurator.RadioCfgResponse
	aes    cipher.Block
}

func (r *receiver) name() string {
	return "receiver"
}

func (r *receiver) run(ctx context.Context) (err error) {
	r.cfg, err = r.runner.configurator.GetRadioCfg(ctx, &configurator.RadioCfgRequest{Cluster: 1})
	if err != nil {
		return
	}

	if r.aes, err = aes.NewCipher(r.cfg.AesKey); err != nil {
		return
	}

	if err = r.runner.radio.cc1101.Reset(); err != nil {
		return
	}

	if err = r.runner.radio.setupBasic(); err != nil {
		return
	}

	_, err = r.runner.radio.cc1101.Xfer(r.cfg.RadioConfig)
	if err != nil {
		return
	}

	if err = r.runner.radio.setupForRX(); err != nil {
		return
	}

	return r.loop(ctx)
}

func (r *receiver) loop(ctx context.Context) (err error) {
	watchdogctx, wcancel := context.WithCancel(ctx)
	defer wcancel()

	watchdogPing := make(chan struct{})

	// watchdog routine
	go func() {
		defer wcancel()

		timer := time.NewTimer(watchdogTimeout)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				log.Printf("Watchdog timeout (no packet received for %+v)", watchdogTimeout)
				return
			case <-watchdogctx.Done():
				return
			case <-watchdogPing:
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(watchdogTimeout)
			}
		}
	}()

	r.runner.interrupt.SetContext(watchdogctx)

	for {
		if err = r.runner.interrupt.Wait(); err != nil {
			return
		}

		var dataLen uint8
		dataLen, err = r.runner.radio.cc1101.GetStatus(cc1101.RXBYTES)
		if err != nil {
			return
		}

		if dataLen < 18 { // should not get here
			continue
		}

		if dataLen&0x80 == 0x80 {
			r.runner.radio.cc1101.Cmd(cc1101.SFRX)
			if err = r.runner.radio.cc1101.Waitstate(1); err != nil {
				return
			}
			r.runner.radio.cc1101.Cmd(cc1101.SRX)

			continue
		}

		if dataLen > 54 {
			r.runner.radio.cc1101.Cmd(cc1101.SIDLE)
			if err = r.runner.radio.cc1101.Waitstate(1); err != nil {
				return
			}
			r.runner.radio.cc1101.Cmd(cc1101.SFRX)
			r.runner.radio.cc1101.Cmd(cc1101.SRX)

			continue
		}

		var p []byte
		p, err = r.runner.radio.cc1101.ReadRXFifo(int(dataLen/18) * 18)
		if err != nil {
			return
		}

		for len(p) >= 18 {
			go r.receive(ctx, p[:18])

			p = p[18:]
		}

		watchdogPing <- struct{}{}
	}
}

func (r *receiver) receive(ctx context.Context, p []byte) {
	rssi := float32(int8(p[16]))/2 - 74
	lqi := uint32(p[17] & 0x7f)

	p = p[:16]
	r.aes.Decrypt(p, p)

	network := uint32(binary.LittleEndian.Uint16(p[:2]))
	length := p[6]

	if network != r.cfg.Network {
		log.Printf("Received packet for invalid network: %d", network)
		return
	}

	sensorPacket := &aggregator.SensorPacket{
		Id:   uint32(p[7]),
		Seq:  binary.LittleEndian.Uint32(p[2:6]),
		Rssi: rssi,
		Lqi:  lqi,
		Raw:  p[8:length],
	}

	resp, err := r.runner.aggregator.FeedSensorPacket(ctx, sensorPacket)
	if err != nil {
		return
	}

	log.Printf("Packet from %02x: processed=%v", sensorPacket.Id, resp.Processed)
}
