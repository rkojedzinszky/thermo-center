package receiver

import (
	"context"
	"fmt"
	"log"
	"time"

	configurator "github.com/rkojedzinszky/thermo-center/configurator"
	"github.com/rkojedzinszky/thermo-center/receiver/cc1101"
)

const (
	discoveryPacketTimeout = 22 * time.Second
)

func (r *Runner) configTask(task *configurator.Task) task {
	return &configTask{
		r:    r,
		task: task,
	}
}

type configTask struct {
	r    *Runner
	task *configurator.Task
}

func (c *configTask) name() string {
	return "configurator"
}

func (c *configTask) run(ctx context.Context) error {
	// switch to receiver mode after exit
	defer func() {
		select {
		case <-ctx.Done():
		case c.r.task <- c.r.receiverTask():
		}

		// Wait to be cancelled
		<-ctx.Done()
	}()

	task, err := c.r.configurator.TaskAcquire(ctx, c.task)
	if err != nil {
		return err
	}

	if task == nil {
		return fmt.Errorf("Received nil task")
	}

	finishRequest := configurator.TaskFinishedRequest{
		TaskId: task.TaskId,
	}
	defer c.r.configurator.TaskFinished(ctx, &finishRequest)

	replPacket := prepareReplyPacket(task)

	if err = c.r.radio.cc1101.Reset(); err != nil {
		return err
	}

	if err = c.r.radio.setupBasic(); err != nil {
		return err
	}

	if err = c.r.radio.setupForConf(); err != nil {
		return err
	}

	deadline := time.Now().Add(discoveryPacketTimeout)
	seen := false

LOOP:
	for {
		intctx, cancel := context.WithDeadline(ctx, deadline)
		sensorID, err := c.waitDiscoveryPacket(intctx)
		cancel()

		select {
		case <-ctx.Done():
			return nil
		default:
		}

		// Assume timeout has occured
		if err != nil {
			break LOOP
		}

		if (sensorID & 0x80) == 0x80 {
			log.Printf("Received discovery packet from %d\n", sensorID)
		} else {
			log.Printf("Received reconfiguration request from %d\n", sensorID)
		}

		if (sensorID&0x80) == 0 && sensorID != uint8(task.SensorId) {
			log.Println("Received unexpected reconfiguration discovery packet")
			continue
		}

		if !seen {
			replPacket[1] = sensorID
			seen = true
		} else if replPacket[1] != sensorID {
			log.Printf("Ignoring unexpected discovery packet from %d, expecting %d", sensorID, replPacket[1])
			continue
		}

		c.sendReplyPacket(replPacket)

		log.Printf("Replied to %d\n", replPacket[1])

		if _, err := c.r.configurator.TaskDiscoveryReceived(ctx, c.task); err != nil {
			return err
		}

		deadline = time.Now().Add(discoveryPacketTimeout)
	}

	if !seen {
		finishRequest.Error = "No discovery received"
	}

	return nil
}

func prepareReplyPacket(t *configurator.TaskDetails) []byte {
	packet := make([]byte, 0, 64)

	packet = append(packet,
		54, // total length
		0,  // reply address

		// configuration starts here
		0,                           // crc
		byte(t.SensorId),            // id
		byte(t.Config.Network&0xff), // network id lsb
		byte(t.Config.Network>>8),   // network id msb
	)
	packet = append(packet, t.Config.AesKey...)
	packet = append(packet, t.Config.RadioConfig...)

	for len(packet) < int(packet[0]) {
		packet = append(packet, 0xff)
	}

	packet[0]--

	return packet
}

func (c *configTask) waitDiscoveryPacket(ctx context.Context) (uint8, error) {
	c.r.interrupt.SetContext(ctx)

	for {
		if err := c.r.radio.cc1101.Cmd(cc1101.SRX); err != nil {
			return 0, err
		}

		if err := c.r.interrupt.Wait(); err != nil {
			return 0, err
		}

		dataLen, err := c.r.radio.cc1101.GetStatus(cc1101.RXBYTES)
		if err != nil {
			return 0, err
		}

		data, err := c.r.radio.cc1101.ReadRXFifo(int(dataLen))
		if err != nil {
			return 0, err
		}

		if dataLen == 3 && data[0] == 2 && data[1] == 0 {
			return data[2], nil
		}

		if err = c.r.radio.cc1101.Cmd(cc1101.SFRX); err != nil {
			return 0, err
		}
	}
}

func (c *configTask) sendReplyPacket(packet []byte) error {
	if err := c.r.radio.cc1101.WriteTXFifo(packet); err != nil {
		return err
	}

	return c.r.radio.cc1101.Waitstate(1)
}
