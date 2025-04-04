package workflow

import (
	"fmt"
	"time"

	"github.com/nerdynz/go-workflows/internal/command"
	"github.com/nerdynz/go-workflows/internal/contextvalue"
	"github.com/nerdynz/go-workflows/internal/sync"
	"github.com/nerdynz/go-workflows/internal/tracing"
	"github.com/nerdynz/go-workflows/internal/workflowstate"
)

type timerConfig struct {
	Name string
}

type withNameOption struct {
	Name string
}

// applySpanEnd implements TimerOption.
func (w withNameOption) apply(tc timerConfig) timerConfig {
	tc.Name = w.Name
	return tc
}

var _ timerOption = withNameOption{}

type timerOption interface {
	apply(timerConfig) timerConfig
}

func WithTimerName(name string) timerOption {
	return withNameOption{
		Name: name,
	}
}

// ScheduleTimer schedules a timer to fire after the given delay.
func ScheduleTimer(ctx Context, delay time.Duration, opts ...timerOption) Future[any] {
	var timerConfig timerConfig
	for _, opt := range opts {
		timerConfig = opt.apply(timerConfig)
	}

	f := sync.NewFuture[any]()

	// If the context is already canceled, return immediately.
	if ctx.Err() != nil {
		f.Set(struct{}{}, ctx.Err())
		return f
	}

	wfState := workflowstate.WorkflowState(ctx)
	scheduleEventID := wfState.GetNextScheduleEventID()

	at := Now(ctx).Add(delay)

	traceContext := tracing.ContextFromWfCtx(ctx)

	timerCmd := command.NewScheduleTimerCommand(scheduleEventID, at, timerConfig.Name, traceContext)
	wfState.AddCommand(timerCmd)

	timerSuffix := ""
	if timerConfig.Name != "" {
		timerSuffix = "-" + timerConfig.Name
	}

	wfState.TrackFuture(
		scheduleEventID,
		workflowstate.AsDecodingSettable(contextvalue.Converter(ctx), fmt.Sprintf("timer%s:%v", timerSuffix, delay), f))

	cancelReceiver := &sync.Receiver[struct{}]{
		Receive: func(v struct{}, ok bool) {
			timerCmd.Cancel()

			// Remove the timer future from the workflow state and mark it as canceled if it hasn't already fired. This is different
			// from subworkflow behavior, where we want to wait for the subworkflow to complete before proceeding. Here we can
			// continue right away.
			if fi, ok := f.(sync.FutureInternal[any]); ok {
				if !fi.Ready() {
					wfState.RemoveFuture(scheduleEventID)
					f.Set(v, Canceled)
				}
			}
		},
	}

	// Check if the context is cancelable
	if c, cancelable := ctx.Done().(sync.CancelChannel); cancelable {
		// Register a callback for when it's canceled. The only operation on the `Done` channel
		// is that it's closed when the context is canceled.
		c.AddReceiveCallback(cancelReceiver)

		timerCmd.WhenDone(func() {
			c.RemoveReceiveCallback(cancelReceiver)
		})
	}

	return f
}
