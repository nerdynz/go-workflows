package postgresbackend

import (
	"context"
	"database/sql"
	"time"

	"github.com/nerdynz/go-workflows/core"
	"github.com/nerdynz/go-workflows/diag"
)

var _ diag.Backend = (*postgresBackend)(nil)

func (mb *postgresBackend) GetWorkflowInstances(ctx context.Context, afterInstanceID, afterExecutionID string, count int) ([]*diag.WorkflowInstanceRef, error) {
	var err error
	tx, err := mb.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var rows *sql.Rows
	if afterInstanceID != "" {
		rows, err = tx.QueryContext(
			ctx,
			SQLReplacer(`SELECT i.instance_id, i.execution_id, i.parent_instance_id, i.parent_execution_id, i.parent_schedule_event_id, i.created_at, i.completed_at, i.queue
			FROM instances i
			INNER JOIN (SELECT instance_id, created_at FROM instances WHERE instance_id = ? AND execution_id = ?) ii
				ON i.created_at < ii.created_at OR (i.created_at = ii.created_at AND i.instance_id < ii.instance_id)
			ORDER BY i.created_at DESC, i.instance_id DESC
			LIMIT ?`),
			afterInstanceID,
			afterExecutionID,
			count,
		)
	} else {
		rows, err = tx.QueryContext(
			ctx,
			SQLReplacer(`SELECT i.instance_id, i.execution_id, i.parent_instance_id, i.parent_execution_id, i.parent_schedule_event_id, i.created_at, i.completed_at, i.queue
			FROM instances i
			ORDER BY i.created_at DESC, i.instance_id DESC
		LIMIT ?`),
			count,
		)
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var instances []*diag.WorkflowInstanceRef

	for rows.Next() {
		var id, executionID, queue string
		var parentID, parentExecutionID *string
		var parentScheduleEventID *int64
		var createdAt time.Time
		var completedAt *time.Time
		err = rows.Scan(&id, &executionID, &parentID, &parentExecutionID, &parentScheduleEventID, &createdAt, &completedAt, &queue)
		if err != nil {
			return nil, err
		}

		var state core.WorkflowInstanceState
		if completedAt != nil {
			state = core.WorkflowInstanceStateFinished
		}

		var instance *core.WorkflowInstance
		if parentID != nil {
			parentInstance := core.NewWorkflowInstance(*parentID, *parentExecutionID)
			instance = core.NewSubWorkflowInstance(id, executionID, parentInstance, *parentScheduleEventID)
		} else {
			instance = core.NewWorkflowInstance(id, executionID)
		}

		instances = append(instances, &diag.WorkflowInstanceRef{
			Instance:    instance,
			CreatedAt:   createdAt,
			CompletedAt: completedAt,
			State:       state,
			Queue:       queue,
		})
	}

	return instances, nil
}

func (mb *postgresBackend) GetWorkflowInstance(ctx context.Context, instance *core.WorkflowInstance) (*diag.WorkflowInstanceRef, error) {
	tx, err := mb.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res := tx.QueryRowContext(
		ctx,
		SQLReplacer(`SELECT instance_id, execution_id, parent_instance_id, parent_execution_id, parent_schedule_event_id, created_at, completed_at, queue
			FROM instances
			WHERE instance_id = ? AND execution_id = ?`), instance.InstanceID, instance.ExecutionID)

	var id, executionID, queue string
	var parentID, parentExecutionID *string
	var parentScheduleEventID *int64
	var createdAt time.Time
	var completedAt *time.Time

	err = res.Scan(&id, &executionID, &parentID, &parentExecutionID, &parentScheduleEventID, &createdAt, &completedAt, &queue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var state core.WorkflowInstanceState
	if completedAt != nil {
		state = core.WorkflowInstanceStateFinished
	}

	if parentID != nil {
		parentInstance := core.NewWorkflowInstance(*parentID, *parentExecutionID)
		instance = core.NewSubWorkflowInstance(id, executionID, parentInstance, *parentScheduleEventID)
	} else {
		instance = core.NewWorkflowInstance(id, executionID)
	}

	return &diag.WorkflowInstanceRef{
		Instance:    core.NewWorkflowInstance(id, executionID),
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		State:       state,
		Queue:       queue,
	}, nil
}

func (mb *postgresBackend) GetWorkflowTree(ctx context.Context, instance *core.WorkflowInstance) (*diag.WorkflowInstanceTree, error) {
	itb := diag.NewInstanceTreeBuilder(mb)
	return itb.BuildWorkflowInstanceTree(ctx, instance)
}
