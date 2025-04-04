package activitytester

import (
	"context"
	"testing"

	"github.com/nerdynz/go-workflows/activity"
	"github.com/stretchr/testify/require"
)

func Activity(ctx context.Context, a int, b int) (int, error) {
	activity.Logger(ctx).Debug("Activity is called", "a", a)

	return a + b, nil
}

func TestActivityTester(t *testing.T) {
	ctx := WithActivityTestState(context.Background(), "activityID", "instanceID", nil)

	r, err := Activity(ctx, 35, 12)
	require.Equal(t, 47, r)
	require.NoError(t, err)
}
