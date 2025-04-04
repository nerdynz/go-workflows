package history

import "github.com/nerdynz/go-workflows/internal/workflowerrors"

type SubWorkflowFailedAttributes struct {
	Error *workflowerrors.Error `json:"error,omitempty"`
}
