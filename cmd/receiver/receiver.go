package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
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
	receiverPort := flag.Int("receiver-port", 8079, "Receiver port")

	flag.Parse()

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	port, err := spireg.Open(fmt.Sprintf("SPI%d.%d", *spiBusNum, *spiCSNum))
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	spiConn, err := port.Connect(physic.Frequency(*spiFreq)*physic.Hertz, spi.Mode(*spiMode), 8)
	if err != nil {
		log.Fatal(err)
	}

	ih, err := gpiointerrupt.New(*gpioDir)
	if err != nil {
		log.Fatal(err)
	}
	defer ih.Close()

	grpcClient, err := grpc.Dial(fmt.Sprintf("%s:%d", *grpcHost, *grpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *receiverPort))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	runner := receiver.NewRunner(grpcClient, cc1101.New(spiConn), ih)

	receiver.RegisterReceiverServer(grpcServer, runner)

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		go grpcServer.Serve(listen)
		<-ctx.Done()
		grpcServer.Stop()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		runner.Run(ctx)
	}()

	// Watch for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	<-sigChan

	cancel()

	wg.Wait()
}
