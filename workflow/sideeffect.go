package workflow

import (
	"github.com/nerdynz/go-workflows/internal/command"
	"github.com/nerdynz/go-workflows/internal/contextvalue"
	"github.com/nerdynz/go-workflows/internal/sync"
	"github.com/nerdynz/go-workflows/internal/workflowstate"
)

// SideEffect executes the given function and returns a future that will be resolved with the result of
// the function.
//
// In contrast to Activities, SideEffects are executed inline with the workflow code. They should only
// be used for short and inexpensive operations. For longer operations, consider using an Activity.
func SideEffect[TResult any](ctx Context, f func(ctx Context) TResult) Future[TResult] {
	ctx, span := Tracer(ctx).Start(ctx, "SideEffect")
	defer span.End()

	future := sync.NewFuture[TResult]()

	if ctx.Err() != nil {
		future.Set(*new(TResult), ctx.Err())
		return future
	}

	wfState := workflowstate.WorkflowState(ctx)
	scheduleEventID := wfState.GetNextScheduleEventID()

	cv := contextvalue.Converter(ctx)
	wfState.TrackFuture(scheduleEventID, workflowstate.AsDecodingSettable(cv, "sideeffect", future))

	cmd := command.NewSideEffectCommand(scheduleEventID)
	wfState.AddCommand(cmd)

	if !Replaying(ctx) {
		// Execute side effect
		r := f(ctx)

		payload, err := cv.To(r)
		if err != nil {
			future.Set(*new(TResult), err)
		}

		cmd.SetResult(payload)
		future.Set(r, nil)
		wfState.RemoveFuture(scheduleEventID)
	}

	return future
}
