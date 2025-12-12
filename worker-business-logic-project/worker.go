package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	TemporalHostPort = "localhost:7233"
	TaskQueue        = "hello-world-task-queue"
)

func startWorker() error {
	// Create Temporal client
	c, err := client.Dial(client.Options{
		HostPort: TemporalHostPort,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	// Create worker
	w := worker.New(c, TaskQueue, worker.Options{})

	// Register workflows and activities
	w.RegisterWorkflow(HelloWorkflow)
	w.RegisterActivity(SayHelloActivity)

	log.Println("Temporal Worker starting...")
	log.Printf("- Temporal Host: %s\n", TemporalHostPort)
	log.Printf("- Task Queue: %s\n", TaskQueue)
	log.Println("Waiting for workflow tasks...")

	// Start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		return err
	}

	return nil
}
