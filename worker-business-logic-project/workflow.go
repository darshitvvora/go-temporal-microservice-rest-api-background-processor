package main

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type HelloWorkflowInput struct {
	Name string
}

type HelloWorkflowOutput struct {
	Message string
}

func HelloWorkflow(ctx workflow.Context, input HelloWorkflowInput) (HelloWorkflowOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, SayHelloActivity, input.Name).Get(ctx, &result)
	if err != nil {
		return HelloWorkflowOutput{}, err
	}

	return HelloWorkflowOutput{Message: result}, nil
}
