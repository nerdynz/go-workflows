package workflowstate

import (
	"log/slog"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/google/uuid"
	"github.com/nerdynz/go-workflows/backend/converter"
	"github.com/nerdynz/go-workflows/core"
	"github.com/nerdynz/go-workflows/internal/sync"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace/noop"
)

func Test_PendingFutures(t *testing.T) {
	i := core.NewWorkflowInstance(uuid.NewString(), "")

	wfState := NewWorkflowState(i, slog.Default(), noop.NewTracerProvider().Tracer("test"), clock.New())

	require.False(t, wfState.HasPendingFutures())

	f := sync.NewFuture[int]()
	wfState.TrackFuture(1, AsDecodingSettable[int](converter.DefaultConverter, "f", f))

	require.True(t, wfState.HasPendingFutures())

	wfState.RemoveFuture(1)

	require.False(t, wfState.HasPendingFutures())
}
