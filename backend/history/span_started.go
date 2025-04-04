package history

import "github.com/nerdynz/go-workflows/backend/payload"

type TraceStartedAttributes struct {
	SpanID payload.Payload `json:"spanID"`
}
