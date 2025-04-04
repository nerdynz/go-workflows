package tester

import (
	"context"

	"github.com/nerdynz/go-workflows/core"
	"github.com/nerdynz/go-workflows/internal/signals"
)

type signaler[T any] struct {
	wt *workflowTester[T]
}

func (s *signaler[T]) SignalWorkflow(ctx context.Context, instanceID string, name string, arg interface{}) error {
	return s.wt.SignalWorkflowInstance(core.NewWorkflowInstance(instanceID, ""), name, arg)
}

var _ signals.Signaler = (*signaler[any])(nil)
