package history

import "github.com/nerdynz/go-workflows/internal/workflowerrors"

type ActivityFailedAttributes struct {
	Error *workflowerrors.Error `json:"error,omitempty"`
}
