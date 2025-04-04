package workflow

import (
	"github.com/nerdynz/go-workflows/backend/metadata"
	"github.com/nerdynz/go-workflows/core"
)

type (
	// Queue represents a queue for workflow instances and activities.
	Queue = core.Queue

	// Instance represents a workflow instance.
	Instance = core.WorkflowInstance

	// Metadata represents the metadata of a workflow instance.
	Metadata = metadata.WorkflowMetadata

	Workflow = interface{}

	Activity = interface{}
)

const (
	QueueDefault = core.QueueDefault
)
