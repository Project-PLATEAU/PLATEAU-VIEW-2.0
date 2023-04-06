package gateway

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/task"
)

type TaskRunner interface {
	Run(context.Context, task.Payload) error
}
