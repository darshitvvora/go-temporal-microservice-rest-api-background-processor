package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.temporal.io/sdk/client"

	"workflow-trigger-project/handlers"
)

const (
	TemporalHostPort = "localhost:7233"
	TaskQueue        = "hello-world-task-queue"
	ServerPort       = ":3000"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: TemporalHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	app := fiber.New(fiber.Config{
		AppName: "Temporal Workflow Trigger API",
	})

	app.Use(logger.New())

	workflowHandler := handlers.NewWorkflowHandler(c, TaskQueue)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Temporal Workflow Trigger API",
			"version": "1.0.0",
		})
	})

	app.Post("/api/workflow/trigger", workflowHandler.TriggerWorkflow)
	app.Get("/api/workflow/status/:workflowId", workflowHandler.GetWorkflowStatus)

	log.Printf("Server starting on port %s\n", ServerPort)
	if err := app.Listen(ServerPort); err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
