package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"

	"workflow-trigger-project/models"
)

type WorkflowHandler struct {
	temporalClient client.Client
	taskQueue      string
}

func NewWorkflowHandler(temporalClient client.Client, taskQueue string) *WorkflowHandler {
	return &WorkflowHandler{
		temporalClient: temporalClient,
		taskQueue:      taskQueue,
	}
}

func (h *WorkflowHandler) TriggerWorkflow(c *fiber.Ctx) error {
	var req models.TriggerWorkflowRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	workflowID := fmt.Sprintf("hello-workflow-%d", time.Now().Unix())

	input := models.HelloWorkflowInput{
		Name: req.Name,
	}

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: h.taskQueue,
	}

	we, err := h.temporalClient.ExecuteWorkflow(
		context.Background(),
		options,
		"HelloWorkflow",
		input,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to start workflow: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.TriggerWorkflowResponse{
		WorkflowID: we.GetID(),
		RunID:      we.GetRunID(),
		Message:    "Workflow triggered successfully",
	})
}

func (h *WorkflowHandler) GetWorkflowStatus(c *fiber.Ctx) error {
	workflowID := c.Params("workflowId")
	runID := c.Query("runId", "")

	if workflowID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Workflow ID is required",
		})
	}

	we := h.temporalClient.GetWorkflow(context.Background(), workflowID, runID)

	var result models.HelloWorkflowOutput
	err := we.Get(context.Background(), &result)

	response := models.WorkflowStatusResponse{
		WorkflowID: workflowID,
		RunID:      runID,
	}

	if err != nil {
		response.Status = "failed or running"
		response.Error = err.Error()
	} else {
		response.Status = "completed"
		response.Result = &result
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
