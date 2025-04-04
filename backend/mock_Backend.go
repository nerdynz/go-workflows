// Code generated by mockery v2.20.0. DO NOT EDIT.

package backend

import (
	context "context"

	history "github.com/nerdynz/go-workflows/backend/history"
	core "github.com/nerdynz/go-workflows/core"

	metrics "github.com/nerdynz/go-workflows/backend/metrics"

	mock "github.com/stretchr/testify/mock"

	trace "go.opentelemetry.io/otel/trace"
)

// MockBackend is an autogenerated mock type for the Backend type
type MockBackend struct {
	mock.Mock
}

// CancelWorkflowInstance provides a mock function with given fields: ctx, instance, cancelEvent
func (_m *MockBackend) CancelWorkflowInstance(ctx context.Context, instance *core.WorkflowInstance, cancelEvent *history.Event) error {
	ret := _m.Called(ctx, instance, cancelEvent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *history.Event) error); ok {
		r0 = rf(ctx, instance, cancelEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockBackend) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteActivityTask provides a mock function with given fields: ctx, task, result
func (_m *MockBackend) CompleteActivityTask(ctx context.Context, task *ActivityTask, result *history.Event) error {
	ret := _m.Called(ctx, task, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ActivityTask, *history.Event) error); ok {
		r0 = rf(ctx, task, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteWorkflowTask provides a mock function with given fields: ctx, task, state, executedEvents, activityEvents, timerEvents, workflowEvents
func (_m *MockBackend) CompleteWorkflowTask(ctx context.Context, task *WorkflowTask, state core.WorkflowInstanceState, executedEvents []*history.Event, activityEvents []*history.Event, timerEvents []*history.Event, workflowEvents []*history.WorkflowEvent) error {
	ret := _m.Called(ctx, task, state, executedEvents, activityEvents, timerEvents, workflowEvents)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *WorkflowTask, core.WorkflowInstanceState, []*history.Event, []*history.Event, []*history.Event, []*history.WorkflowEvent) error); ok {
		r0 = rf(ctx, task, state, executedEvents, activityEvents, timerEvents, workflowEvents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWorkflowInstance provides a mock function with given fields: ctx, instance, event
func (_m *MockBackend) CreateWorkflowInstance(ctx context.Context, instance *core.WorkflowInstance, event *history.Event) error {
	ret := _m.Called(ctx, instance, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *history.Event) error); ok {
		r0 = rf(ctx, instance, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendActivityTask provides a mock function with given fields: ctx, task
func (_m *MockBackend) ExtendActivityTask(ctx context.Context, task *ActivityTask) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ActivityTask) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendWorkflowTask provides a mock function with given fields: ctx, task
func (_m *MockBackend) ExtendWorkflowTask(ctx context.Context, task *WorkflowTask) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *WorkflowTask) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FeatureSupported provides a mock function with given fields: feature
func (_m *MockBackend) FeatureSupported(feature Feature) bool {
	ret := _m.Called(feature)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Feature) bool); ok {
		r0 = rf(feature)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetActivityTask provides a mock function with given fields: ctx, queues
func (_m *MockBackend) GetActivityTask(ctx context.Context, queues []core.Queue) (*ActivityTask, error) {
	ret := _m.Called(ctx, queues)

	var r0 *ActivityTask
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) (*ActivityTask, error)); ok {
		return rf(ctx, queues)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) *ActivityTask); ok {
		r0 = rf(ctx, queues)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ActivityTask)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []core.Queue) error); ok {
		r1 = rf(ctx, queues)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: ctx
func (_m *MockBackend) GetStats(ctx context.Context) (*Stats, error) {
	ret := _m.Called(ctx)

	var r0 *Stats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*Stats, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *Stats); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Stats)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowInstanceHistory provides a mock function with given fields: ctx, instance, lastSequenceID
func (_m *MockBackend) GetWorkflowInstanceHistory(ctx context.Context, instance *core.WorkflowInstance, lastSequenceID *int64) ([]*history.Event, error) {
	ret := _m.Called(ctx, instance, lastSequenceID)

	var r0 []*history.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *int64) ([]*history.Event, error)); ok {
		return rf(ctx, instance, lastSequenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *int64) []*history.Event); ok {
		r0 = rf(ctx, instance, lastSequenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*history.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowInstance, *int64) error); ok {
		r1 = rf(ctx, instance, lastSequenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowInstanceState provides a mock function with given fields: ctx, instance
func (_m *MockBackend) GetWorkflowInstanceState(ctx context.Context, instance *core.WorkflowInstance) (core.WorkflowInstanceState, error) {
	ret := _m.Called(ctx, instance)

	var r0 core.WorkflowInstanceState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance) (core.WorkflowInstanceState, error)); ok {
		return rf(ctx, instance)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance) core.WorkflowInstanceState); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Get(0).(core.WorkflowInstanceState)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowInstance) error); ok {
		r1 = rf(ctx, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowTask provides a mock function with given fields: ctx, queues
func (_m *MockBackend) GetWorkflowTask(ctx context.Context, queues []core.Queue) (*WorkflowTask, error) {
	ret := _m.Called(ctx, queues)

	var r0 *WorkflowTask
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) (*WorkflowTask, error)); ok {
		return rf(ctx, queues)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) *WorkflowTask); ok {
		r0 = rf(ctx, queues)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*WorkflowTask)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []core.Queue) error); ok {
		r1 = rf(ctx, queues)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Metrics provides a mock function with given fields:
func (_m *MockBackend) Metrics() metrics.Client {
	ret := _m.Called()

	var r0 metrics.Client
	if rf, ok := ret.Get(0).(func() metrics.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Client)
		}
	}

	return r0
}

// Options provides a mock function with given fields:
func (_m *MockBackend) Options() *Options {
	ret := _m.Called()

	var r0 *Options
	if rf, ok := ret.Get(0).(func() *Options); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Options)
		}
	}

	return r0
}

// PrepareActivityQueues provides a mock function with given fields: ctx, queues
func (_m *MockBackend) PrepareActivityQueues(ctx context.Context, queues []core.Queue) error {
	ret := _m.Called(ctx, queues)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) error); ok {
		r0 = rf(ctx, queues)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrepareWorkflowQueues provides a mock function with given fields: ctx, queues
func (_m *MockBackend) PrepareWorkflowQueues(ctx context.Context, queues []core.Queue) error {
	ret := _m.Called(ctx, queues)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []core.Queue) error); ok {
		r0 = rf(ctx, queues)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveWorkflowInstance provides a mock function with given fields: ctx, instance
func (_m *MockBackend) RemoveWorkflowInstance(ctx context.Context, instance *core.WorkflowInstance) error {
	ret := _m.Called(ctx, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance) error); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveWorkflowInstances provides a mock function with given fields: ctx, options
func (_m *MockBackend) RemoveWorkflowInstances(ctx context.Context, options ...RemovalOption) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...RemovalOption) error); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignalWorkflow provides a mock function with given fields: ctx, instanceID, event
func (_m *MockBackend) SignalWorkflow(ctx context.Context, instanceID string, event *history.Event) error {
	ret := _m.Called(ctx, instanceID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *history.Event) error); ok {
		r0 = rf(ctx, instanceID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tracer provides a mock function with given fields:
func (_m *MockBackend) Tracer() trace.Tracer {
	ret := _m.Called()

	var r0 trace.Tracer
	if rf, ok := ret.Get(0).(func() trace.Tracer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(trace.Tracer)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBackend creates a new instance of MockBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBackend(t mockConstructorTestingTNewMockBackend) *MockBackend {
	mock := &MockBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
