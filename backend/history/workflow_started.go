package history

import (
	"github.com/nerdynz/go-workflows/backend/metadata"
	"github.com/nerdynz/go-workflows/backend/payload"
	"github.com/nerdynz/go-workflows/core"
)

type ExecutionStartedAttributes struct {
	Queue core.Queue `json:"queue,omitempty"`

	Name string `json:"name,omitempty"`

	Metadata *metadata.WorkflowMetadata `json:"metadata,omitempty"`

	Inputs []payload.Payload `json:"inputs,omitempty"`

	WorkflowSpanID [8]byte `json:"workflowSpanID,omitempty"`
}
