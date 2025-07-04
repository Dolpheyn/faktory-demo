package main

import (
	"context"
	"errors"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
	"github.com/google/uuid"
)

func SendOnboardingEmail(ctx context.Context, args ...interface{}) error {
	time.Sleep(time.Millisecond * time.Duration(rand.IntN(1000)))
	help := worker.HelperFor(ctx)
	log.Printf("[SendOnboardingEmail] Working on job %s\n", help.Jid())
	if rand.IntN(100) > 90 {
		return errors.New("helllo")
	}
	return nil
}

var (
	queues = []string{"send-onboarding-email"}

	jobHandlers = map[string]worker.Perform{
		"send-onboarding-email": SendOnboardingEmail,
	}
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mgr := newJobManager(queues)
	for name, fn := range jobHandlers {
		mgr.Register(name, fn)
	}

	go func() { mgr.RunWithContext(ctx) }()
	go func() { simulateMockJobs(ctx, queues) }()

	go func() {
		stopSignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, stopSignals...)

		select {
		case <-ctx.Done():
			return
		case <-stop:
			break
		}

		cancel()
	}()

	<-ctx.Done()
	mgr.Terminate(true)
}

func newJobManager(queues []string) *worker.Manager {
	mgr := worker.NewManager()
	// use up to N goroutines to execute jobs
	mgr.Concurrency = 20
	// wait up to 25 seconds to let jobs in progress finish
	mgr.ShutdownTimeout = 25 * time.Second
	// pull jobs from these queues, in this order of precedence
	mgr.ProcessStrictPriorityQueues(queues...)

	return mgr
}

func simulateMockJobs(ctx context.Context, queues []string) {
	faktory, err := client.Open()
	if err != nil {
		log.Printf("[simulateMockJobs] failed to connect to faktory: %v\n", err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(time.Millisecond * time.Duration(rand.IntN(500)))

		jobID := uuid.NewString()
		queue := queues[rand.IntN(10)%len(queues)]

		log.Printf("[simulateMockJobs] pushing job. jobID=%v\n", jobID)
		err := faktory.Push(&client.Job{
			Retry:   new(int),
			Failure: &client.Failure{},
			Jid:     jobID,
			Queue:   queue,
			Type:    queue,
			Args: []any{
				jobID,
			},
		})

		if err != nil {
			log.Printf("[simulateMockJobs] err pushing job. jobID=%v, err=%v\n", jobID, err)
		}
	}

}
