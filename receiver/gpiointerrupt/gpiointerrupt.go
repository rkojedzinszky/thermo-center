package gpiointerrupt

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"
	"syscall"
)

// Interrupt provides an interrupt source received from an exportd
// GPIO pin
type Interrupt struct {
	epollFd int
	valueFh *os.File
}

// New allocates a new GPI
func New(dir string) (i *Interrupt, err error) {
	i = &Interrupt{
		epollFd: -1,
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

// Wait waits for one interrupt to arrive
func (i *Interrupt) Wait(ctx context.Context) error {
	wchan := make(chan struct{})
	defer close(wchan)

	fds := make([]int, 2)
	if err := syscall.Pipe(fds); err != nil {
		return err
	}

	defer syscall.Close(fds[0])

	go func() {
		defer syscall.Close(fds[1])

		select {
		case <-wchan:
		case <-ctx.Done():
		}
	}()

	event := syscall.EpollEvent{
		Fd:     int32(fds[0]),
		Events: syscall.EPOLLIN,
	}

	if err := syscall.EpollCtl(i.epollFd, syscall.EPOLL_CTL_ADD, int(event.Fd), &event); err != nil {
		return err
	}
	defer syscall.EpollCtl(i.epollFd, syscall.EPOLL_CTL_DEL, int(event.Fd), &event)

	for i.value() != 1 {
		events := make([]syscall.EpollEvent, 2)

		nevents, _ := syscall.EpollWait(i.epollFd, events, -1)
		for e := 0; e < nevents; e++ {
			if events[e].Fd == int32(fds[0]) {
				return fmt.Errorf("Interrupt.Wait() cancelled")
			}
		}
	}

	return nil
}

// Close closes all file descriptors
func (i *Interrupt) Close() {
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
