package workflow

import (
	"log/slog"

	"github.com/nerdynz/go-workflows/internal/workflowstate"
)

// Logger returns the logger for the current workflow.
func Logger(ctx Context) *slog.Logger {
	wfState := workflowstate.WorkflowState(ctx)
	return wfState.Logger()
}
