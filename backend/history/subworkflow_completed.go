package history

import "github.com/nerdynz/go-workflows/backend/payload"

type SubWorkflowCompletedAttributes struct {
	Result payload.Payload `json:"result,omitempty"`
}
