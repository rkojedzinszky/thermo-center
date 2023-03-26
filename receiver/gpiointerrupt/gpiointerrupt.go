package gpiointerrupt

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"
	"syscall"

	"github.com/rkojedzinszky/thermo-center/v5/receiver/tbf"
)

const (
	interruptMaxRate = 5
	interruptBurst   = 60
)

// Interrupt provides an interrupt source received from an exportd
// GPIO pin
type Interrupt struct {
	epollFd int
	valueFh *os.File

	// Context handling
	ctxRfd   int
	ctxClose chan struct{}

	// Storm handling
	tbf *tbf.TokenBucket
}

// New allocates a new GPI
func New(dir string) (i *Interrupt, err error) {
	i = &Interrupt{
		epollFd: -1,
		ctxRfd:  -1,
		tbf:     tbf.New(interruptMaxRate, interruptBurst),
	}

	defer func() {
		if err != nil {
			if i.valueFh != nil {
				i.valueFh.Close()
			}

			if i.epollFd != -1 {
				syscall.Close(i.epollFd)
			}

			i = nil
		}
	}()

	if err = direction(dir, "in"); err != nil {
		return
	}

	if err = edge(dir, "none"); err != nil {
		return
	}

	if err = edge(dir, "rising"); err != nil {
		return
	}

	if i.epollFd, err = syscall.EpollCreate1(0); err != nil {
		return
	}

	valuePath := path.Join(dir, "value")

	if i.valueFh, err = os.Open(valuePath); err != nil {
		return
	}

	event := syscall.EpollEvent{
		Fd:     int32(i.valueFh.Fd()),
		Events: syscall.EPOLLPRI,
	}

	if err = syscall.EpollCtl(i.epollFd, syscall.EPOLL_CTL_ADD, int(event.Fd), &event); err != nil {
		return
	}

	return
}

// SetContext sets a context used during Wait().
// If the context gets cancelled, Wait() will return immediately return.
// To clear existing context, pass nil.
func (i *Interrupt) SetContext(ctx context.Context) error {
	var event syscall.EpollEvent

	if i.ctxRfd >= 0 {
		// Clear old fd
		close(i.ctxClose)

		event.Fd = int32(i.ctxRfd)
		event.Events = syscall.EPOLLIN

		syscall.EpollCtl(i.epollFd, syscall.EPOLL_CTL_DEL, int(event.Fd), &event)

		syscall.Close(i.ctxRfd)
	}

	if ctx == nil {
		i.ctxRfd = -1
		return nil
	}

	pipeFD := make([]int, 2)
	if err := syscall.Pipe(pipeFD); err != nil {
		return err
	}

	event.Fd = int32(pipeFD[0])
	event.Events = syscall.EPOLLIN

	if err := syscall.EpollCtl(i.epollFd, syscall.EPOLL_CTL_ADD, int(event.Fd), &event); err != nil {
		syscall.Close(pipeFD[0])
		syscall.Close(pipeFD[1])

		return err
	}

	i.ctxRfd = pipeFD[0]

	// create new close channel
	i.ctxClose = make(chan struct{})

	// Start watching context
	go func() {
		defer syscall.Close(pipeFD[1])

		select {
		case <-ctx.Done():
		case <-i.ctxClose:
		}
	}()

	return nil
}

// Wait waits for one interrupt to arrive
func (i *Interrupt) Wait() error {
	for i.value() != 1 {
		events := make([]syscall.EpollEvent, 2)

		nevents, _ := syscall.EpollWait(i.epollFd, events, -1)
		for e := 0; e < nevents; e++ {
			if events[e].Fd == int32(i.ctxRfd) {
				return fmt.Errorf("Interrupt.Wait() cancelled")
			}
		}
	}

	if i.tbf.Get(1) == false {
		return fmt.Errorf("Interrupt storm detected")
	}

	return nil
}

// Close closes all file descriptors
func (i *Interrupt) Close() {
	i.SetContext(nil)

	syscall.Close(i.epollFd)
	i.valueFh.Close()
}

func (i *Interrupt) value() int {
	i.valueFh.Seek(0, 0)

	buf := make([]byte, 1)
	_, err := i.valueFh.Read(buf)
	if err != nil {
		return -1
	}

	value, _ := strconv.Atoi(string(buf))

	return value
}

func direction(gpio string, dir string) error {
	directionPath := path.Join(gpio, "direction")

	fh, err := os.OpenFile(directionPath, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer fh.Close()

	if _, err := fh.WriteString(dir); err != nil {
		return err
	}

	return nil
}

func edge(gpio string, v string) error {
	edgePath := path.Join(gpio, "edge")

	fh, err := os.OpenFile(edgePath, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer fh.Close()

	if _, err := fh.WriteString(v); err != nil {
		return err
	}

	return nil
}
