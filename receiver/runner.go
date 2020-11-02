package receiver

import (
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rkojedzinszky/thermo-center/aggregator"
	"github.com/rkojedzinszky/thermo-center/configurator"
	"github.com/rkojedzinszky/thermo-center/receiver/cc1101"
	"github.com/rkojedzinszky/thermo-center/receiver/gpiointerrupt"
)

// Run the receiver code
func Run(ctx context.Context, grpcClient *grpc.ClientConn, cc1101 *cc1101.CC1101, gpioInt *gpiointerrupt.Interrupt) {
	r := &runner{
		configurator: configurator.NewConfiguratorClient(grpcClient),
		aggregator:   aggregator.NewAggregatorClient(grpcClient),
		radio:        radio{cc1101: cc1101},
		interrupt:    gpioInt,
		task:         make(chan task, 1),
	}

	// Initial task
	r.task <- newReceiver(r)

	r.run(ctx)
}

type runner struct {
	configurator configurator.ConfiguratorClient
	aggregator   aggregator.AggregatorClient
	radio        radio
	interrupt    *gpiointerrupt.Interrupt

	task chan task
}

func (r *runner) run(ctx context.Context) {
	// Handle context cancellation
	go func() {
		<-ctx.Done()
		close(r.task)
	}()

	task := <-r.task

	for {
		// Run current task in loop
		tctx, tcancel := context.WithCancel(ctx)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()

			log.Printf("%s: starting\n", task.name())
			if err := task.run(ctx); err != nil {
				log.Printf("%s exited with: %+v\n", task.name(), err)
			}
			log.Printf("%s: finished\n", task.name())

			// Repeat until context closed
			select {
			case <-ctx.Done():
				return
			default:
			}

			time.Sleep(1 * time.Second)
		}(tctx)

		// Wait for new task
		var ok bool
		task, ok = <-r.task
		// Stop previous task
		tcancel()
		wg.Wait()

		if !ok {
			break
		}
	}
}

type task interface {
	name() string
	run(ctx context.Context) error
}
