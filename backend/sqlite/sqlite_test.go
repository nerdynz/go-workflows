package sqlite

import (
	"testing"

	"github.com/nerdynz/go-workflows/backend"
	"github.com/nerdynz/go-workflows/backend/test"
	"github.com/stretchr/testify/require"
)

func Test_SqliteBackend(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	test.BackendTest(t, func(options ...backend.BackendOption) test.TestBackend {
		// Disable sticky workflow behavior for the test execution
		return NewInMemoryBackend(WithBackendOptions(append(options, backend.WithStickyTimeout(0))...))
		// return NewSqliteBackend("test.sqlite", WithBackendOptions(append(options, backend.WithStickyTimeout(0))...))
	}, func(b test.TestBackend) {
		// Ensure we close the database so the next test will get a clean in-memory db
		require.NoError(t, b.(*sqliteBackend).Close())
	})
}

func Test_EndToEndSqliteBackend(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	test.EndToEndBackendTest(t, func(options ...backend.BackendOption) test.TestBackend {
		// Disable sticky workflow behavior for the test execution
		return NewInMemoryBackend(WithBackendOptions(append(options, backend.WithStickyTimeout(0))...))
	}, func(b test.TestBackend) {
		// Ensure we close the database so the next test will get a clean in-memory db
		require.NoError(t, b.Close())
	})
}
