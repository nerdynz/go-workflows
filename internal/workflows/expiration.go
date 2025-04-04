package workflows

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/nerdynz/go-workflows/backend"
	"github.com/nerdynz/go-workflows/core"
	"github.com/nerdynz/go-workflows/internal/sync"
	"github.com/nerdynz/go-workflows/workflow"
)

const (
	maxIterations = 10

	UpdateExpirationSignal = "update-expiration"
)

func ExpireWorkflowInstances(ctx workflow.Context, delay time.Duration) error {
	logger := workflow.Logger(ctx)

	updates := workflow.NewSignalChannel[time.Duration](ctx, UpdateExpirationSignal)

	for i := 0; i < maxIterations; i++ {
		tctx, cancelTimer := workflow.WithCancel(ctx)
		t := workflow.ScheduleTimer(tctx, delay)

		timerFired := false
		for !timerFired {
			workflow.Select(ctx,
				workflow.Receive(updates, func(ctx workflow.Context, s time.Duration, _ bool) {
					delay = s

					cancelTimer()
					tctx, cancelTimer = workflow.WithCancel(ctx)
					t = workflow.ScheduleTimer(tctx, delay)
				}),
				workflow.Await(t, func(ctx sync.Context, _ workflow.Future[any]) {
					timerFired = true
				}),
			)
		}

		before := workflow.Now(ctx).Add(-delay)

		logger.Info("removing workflow instances", slog.Time("before", before))

		var a *Activities
		_, err := workflow.ExecuteActivity[any](
			ctx, workflow.ActivityOptions{
				Queue: core.QueueSystem,
				RetryOptions: workflow.RetryOptions{
					MaxAttempts: 2,
				},
			}, a.RemoveWorkflowInstances, before).Get(ctx)
		if err != nil {
			if errors.As(err, &backend.ErrNotSupported{}) {
				logger.Warn("removing workflow instances not supported")

				// Stop execution
				return nil
			}

			logger.Error("removing workflow instances", slog.Any("error", err))
		}
	}

	return workflow.ContinueAsNew(ctx, delay)
}

type Activities struct {
	Backend backend.Backend
}

func (a *Activities) RemoveWorkflowInstances(ctx context.Context, before time.Time) error {
	return a.Backend.RemoveWorkflowInstances(ctx, backend.RemoveFinishedBefore(before))
}
