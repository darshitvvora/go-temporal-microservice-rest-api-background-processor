package models

type HelloWorkflowInput struct {
	Name string `json:"name"`
}

type HelloWorkflowOutput struct {
	Message string `json:"message"`
}

type TriggerWorkflowRequest struct {
	Name string `json:"name" validate:"required"`
}

type TriggerWorkflowResponse struct {
	WorkflowID string `json:"workflowId"`
	RunID      string `json:"runId"`
	Message    string `json:"message"`
}

type WorkflowStatusResponse struct {
	WorkflowID string              `json:"workflowId"`
	RunID      string              `json:"runId"`
	Status     string              `json:"status"`
	Result     *HelloWorkflowOutput `json:"result,omitempty"`
	Error      string              `json:"error,omitempty"`
}
