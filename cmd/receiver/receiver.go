package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"

	"github.com/rkojedzinszky/thermo-center/receiver"
	"github.com/rkojedzinszky/thermo-center/receiver/cc1101"
	"github.com/rkojedzinszky/thermo-center/receiver/gpiointerrupt"
)

func main() {
	spiBusNum := flag.Int("spi-bus-num", 0, "SPI Bus number")
	spiCSNum := flag.Int("spi-cs-num", 0, "SPI Chip-select number")
	spiMode := flag.Int("spi-mode", 0, "SPI mode")
	spiFreq := flag.Int("spi-freq", 100000, "SPI frequency")
	gpioDir := flag.String("gpio-dir", "/gpio", "GPIO dir for interrupt")
	grpcHost := flag.String("grpcserver-host", "grpcserver", "Grpcserver hostname/address")
	grpcPort := flag.Int("grpcserver-port", 8079, "Grpcserver port")

	flag.Parse()

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	port, err := spireg.Open(fmt.Sprintf("SPI%d.%d", *spiBusNum, *spiCSNum))
	if err != nil {
		log.Fatal(err)
	}

	spiConn, err := port.Connect(physic.Frequency(*spiFreq)*physic.Hertz, spi.Mode(*spiMode), 8)
	if err != nil {
		log.Fatal(err)
	}

	ih, err := gpiointerrupt.New(*gpioDir)
	if err != nil {
		log.Fatal(err)
	}

	grpcClient, err := grpc.Dial(fmt.Sprintf("%s:%d", *grpcHost, *grpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

		<-sigChan

		cancel()
	}()

	receiver.Run(ctx, grpcClient, cc1101.New(spiConn), ih)
}
