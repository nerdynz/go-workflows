package workflow

import (
	"github.com/nerdynz/go-workflows/core"
	"github.com/nerdynz/go-workflows/internal/log"
	"github.com/nerdynz/go-workflows/internal/signals"
	"github.com/nerdynz/go-workflows/internal/workflowstate"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// NewSignalChannel returns a new signal channel.
func NewSignalChannel[T any](ctx Context, name string) Channel[T] {
	wfState := workflowstate.WorkflowState(ctx)
	return workflowstate.GetSignalChannel[T](ctx, wfState, name)
}

// SignalWorkflow sends a signal to another running workflow instance.
func SignalWorkflow[T any](ctx Context, instanceID string, name string, arg T) Future[any] {
	ctx, span := Tracer(ctx).Start(ctx, "SignalWorkflow",
		trace.WithAttributes(
			attribute.String(log.SignalNameKey, name),
		),
	)
	defer span.End()

	var a *signals.Activities
	return ExecuteActivity[any](ctx, ActivityOptions{
		RetryOptions: RetryOptions{
			MaxAttempts: 1,
		},
		Queue: core.QueueSystem,
	}, a.DeliverWorkflowSignal, instanceID, name, arg)
}
