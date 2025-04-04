package history

import (
	"github.com/nerdynz/go-workflows/backend/metadata"
	"github.com/nerdynz/go-workflows/backend/payload"
	"github.com/nerdynz/go-workflows/core"
)

type ActivityScheduledAttributes struct {
	Name string `json:"name,omitempty"`

	Attempt int `json:"attempt,omitempty"`

	Inputs []payload.Payload `json:"inputs,omitempty"`

	Metadata *metadata.WorkflowMetadata `json:"metadata,omitempty"`

	Queue core.Queue `json:"queue,omitempty"`
}
