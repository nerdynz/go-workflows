package cache

import (
	"fmt"

	"github.com/nerdynz/go-workflows/core"
)

func getKey(instance *core.WorkflowInstance) string {
	return fmt.Sprintf("%s-%s", instance.InstanceID, instance.ExecutionID)
}
