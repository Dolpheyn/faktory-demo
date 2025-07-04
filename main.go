package main

import (
	"context"
	"errors"
	"fmt"
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

var (
	queues = []string{"send-onboarding-email", "generate-invoice"}

	jobHandlers = map[string]worker.Perform{
		"send-onboarding-email": makeHandler("SendOnboardingEmail"),
		"generate-invoice":      makeHandler("GenerateInvoice"),
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

func makeHandler(jobname string) worker.Perform {
	return func(ctx context.Context, args ...interface{}) error {
		help := worker.HelperFor(ctx)

		// work work work
		log.Printf("[%s] Working on job %s\n", jobname, help.Jid())
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(1000)))

		// error with random rate
		if rand.IntN(100) > 90 {
			return errors.New(fmt.Sprintf("mock error: failed working on %s", help.Jid()))
		}

		return nil
	}
}

func newJobManager(queues []string) *worker.Manager {
	mgr := worker.NewManager()
	// use up to N goroutines to execute jobs
	mgr.Concurrency = 100
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

	retryCount := 3

	for {
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(500)))

		select {
		case <-ctx.Done():
			return
		default:
		}

		go func() {
			jobID := uuid.NewString()
			queue := queues[rand.IntN(len(queues)-1)]

			log.Printf("[simulateMockJobs] pushing job. queue=%s jobID=%v\n", queue, jobID)
			err := faktory.Push(&client.Job{
				Retry:   &retryCount,
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
		}()
	}

}
