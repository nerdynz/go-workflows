package test

import (
	"context"

	"github.com/nerdynz/go-workflows/backend"
	"github.com/nerdynz/go-workflows/backend/history"
)

type TestBackend interface {
	backend.Backend

	GetFutureEvents(ctx context.Context) ([]*history.Event, error)
}
