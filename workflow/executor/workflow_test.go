package executor

import (
	"reflect"
	"testing"

	"github.com/nerdynz/go-workflows/backend/converter"
	"github.com/nerdynz/go-workflows/internal/contextvalue"
	"github.com/nerdynz/go-workflows/internal/sync"
	"github.com/nerdynz/go-workflows/internal/workflowerrors"
	"github.com/stretchr/testify/require"
)

func Test_Workflow_WrapsPanic(t *testing.T) {
	f := func() {
		panic("wf panic")
	}

	w := func(ctx sync.Context) error {
		f()

		return nil
	}

	ctx := sync.Background()
	ctx = contextvalue.WithConverter(ctx, converter.DefaultConverter)

	wf := newWorkflow(reflect.ValueOf(w))
	err := wf.Execute(ctx, nil)
	require.NoError(t, err)

	for !wf.Completed() {
		require.NoError(t, wf.Continue())
	}

	wfErr := wf.Error()
	require.Error(t, wfErr)
	var panicErr *workflowerrors.PanicError
	require.ErrorAs(t, wfErr, &panicErr)
}
