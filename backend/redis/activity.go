package redis

import (
	"context"

	"github.com/nerdynz/go-workflows/backend"
	"github.com/nerdynz/go-workflows/backend/history"
	"github.com/nerdynz/go-workflows/workflow"
)

func (rb *redisBackend) PrepareActivityQueues(ctx context.Context, queues []workflow.Queue) error {
	return rb.activityQueue.Prepare(ctx, rb.rdb, queues)
}

func (rb *redisBackend) GetActivityTask(ctx context.Context, queues []workflow.Queue) (*backend.ActivityTask, error) {
	activityTask, err := rb.activityQueue.Dequeue(ctx, rb.rdb, queues, rb.options.ActivityLockTimeout, rb.options.BlockTimeout)
	if err != nil {
		return nil, err
	}

	if activityTask == nil {
		return nil, nil
	}

	return &backend.ActivityTask{
		WorkflowInstance: activityTask.Data.Instance,
		Queue:            workflow.Queue(activityTask.Data.Queue),
		ID:               activityTask.TaskID, // Use the queue generated ID here
		ActivityID:       activityTask.Data.ID,
		Event:            activityTask.Data.Event,
	}, nil
}

func (rb *redisBackend) ExtendActivityTask(ctx context.Context, task *backend.ActivityTask) error {
	p := rb.rdb.Pipeline()

	if err := rb.activityQueue.Extend(ctx, p, task.Queue, task.ID); err != nil {
		return err
	}

	_, err := p.Exec(ctx)
	return err
}

func (rb *redisBackend) CompleteActivityTask(ctx context.Context, task *backend.ActivityTask, result *history.Event) error {
	instanceState, err := readInstance(ctx, rb.rdb, rb.keys.instanceKey(task.WorkflowInstance))
	if err != nil {
		return err
	}

	p := rb.rdb.TxPipeline()

	if err := rb.addWorkflowInstanceEventP(ctx, p, workflow.Queue(instanceState.Queue), task.WorkflowInstance, result); err != nil {
		return err
	}

	// Unlock activity
	if _, err := rb.activityQueue.Complete(ctx, p, task.Queue, task.ID); err != nil {
		return err
	}

	_, err = p.Exec(ctx)
	return err
}
