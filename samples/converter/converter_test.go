package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/nerdynz/go-workflows/tester"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Workflow(t *testing.T) {
	tester := tester.NewWorkflowTester[int](Workflow1)

	tester.OnActivity(Activity1, mock.Anything, 35, 12).Return(47, nil)
	tester.OnActivity(Activity2, mock.Anything, mock.Anything, mock.Anything).Return(12, nil)

	tester.Execute(context.Background(), "Hello world"+uuid.NewString(), 42, Inputs{
		Msg:   "",
		Times: 0,
	})

	require.True(t, tester.WorkflowFinished())

	wr, werr := tester.WorkflowResult()
	require.Equal(t, 59, wr)
	require.Empty(t, werr)
	tester.AssertExpectations(t)
}
