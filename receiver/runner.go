package receiver

import (
	"fmt"
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

// Runner handles configuration and receiver tasks
type Runner struct {
	configurator configurator.ConfiguratorClient
	aggregator   aggregator.AggregatorClient
	radio        radio
	interrupt    *gpiointerrupt.Interrupt

	task chan task

	UnimplementedReceiverServer
}

// NewRunner creates a Runner instance
func NewRunner(grpcClient *grpc.ClientConn, cc1101 *cc1101.CC1101, gpioInt *gpiointerrupt.Interrupt) *Runner {
	r := &Runner{
		configurator: configurator.NewConfiguratorClient(grpcClient),
		aggregator:   aggregator.NewAggregatorClient(grpcClient),
		radio:        radio{cc1101: cc1101},
		interrupt:    gpioInt,
		task:         make(chan task, 1),
	}

	// Initial task
	r.task <- r.receiverTask()

	return r
}

// Run runs the Runner main loop
func (r *Runner) Run(ctx context.Context) {
	// Handle context cancellation
	go func() {
		<-ctx.Done()
		close(r.task)
	}()

	t := <-r.task

	for {
		// Run current task in loop
		tctx, tcancel := context.WithCancel(ctx)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(ctx context.Context, t task) {
			defer wg.Done()

			for {
				log.Printf("%s: starting", t.name())
				if err := t.run(ctx); err != nil {
					log.Printf("%s exited with: %+v", t.name(), err)
				}
				log.Printf("%s: finished", t.name())

				// Repeat until context closed
				select {
				case <-ctx.Done():
					return
				default:
				}

				time.Sleep(1 * time.Second)
			}
		}(tctx, t)

		// Wait for new task
		var ok bool
		t, ok = <-r.task
		// Stop previous task
		tcancel()
		wg.Wait()

		if !ok {
			break
		}
	}
}

// HandleTask starts handling a configuration task
func (r *Runner) HandleTask(ctx context.Context, task *configurator.Task) (*HandleResponse, error) {
	select {
	case r.task <- r.configTask(task):
	default:
		return &HandleResponse{Success: false}, fmt.Errorf("failed to write to channel")
	}

	return &HandleResponse{Success: true}, nil
}

type task interface {
	name() string
	run(ctx context.Context) error
}
