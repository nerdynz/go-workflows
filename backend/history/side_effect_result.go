package history

import "github.com/nerdynz/go-workflows/backend/payload"

type SideEffectResultAttributes struct {
	Result payload.Payload `json:"result,omitempty"`
}
