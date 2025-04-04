package history

import (
	"github.com/nerdynz/go-workflows/backend/payload"
	"github.com/nerdynz/go-workflows/internal/workflowerrors"
)

type ExecutionCompletedAttributes struct {
	Result payload.Payload       `json:"result,omitempty"`
	Error  *workflowerrors.Error `json:"error,omitempty"`
}
