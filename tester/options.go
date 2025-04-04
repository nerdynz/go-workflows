package tester

import (
	"log/slog"
	"time"

	"github.com/nerdynz/go-workflows/backend/converter"
	"github.com/nerdynz/go-workflows/workflow"
)

type options struct {
	TestTimeout    time.Duration
	Logger         *slog.Logger
	Converter      converter.Converter
	Propagators    []workflow.ContextPropagator
	InitialTime    time.Time
	MaxHistorySize int64
}

type WorkflowTesterOption func(*options)

func WithLogger(logger *slog.Logger) WorkflowTesterOption {
	return func(o *options) {
		o.Logger = logger
	}
}

func WithConverter(converter converter.Converter) WorkflowTesterOption {
	return func(o *options) {
		o.Converter = converter
	}
}

func WithContextPropagator(prop workflow.ContextPropagator) WorkflowTesterOption {
	return func(o *options) {
		o.Propagators = append(o.Propagators, prop)
	}
}

func WithTestTimeout(timeout time.Duration) WorkflowTesterOption {
	return func(o *options) {
		o.TestTimeout = timeout
	}
}

func WithInitialTime(t time.Time) WorkflowTesterOption {
	return func(o *options) {
		o.InitialTime = t
	}
}

func WithMaxHistorySize(size int64) WorkflowTesterOption {
	return func(o *options) {
		o.MaxHistorySize = size
	}
}
