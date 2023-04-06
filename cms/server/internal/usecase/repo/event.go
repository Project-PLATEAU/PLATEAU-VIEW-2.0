package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

type Event interface {
	FindByID(context.Context, id.EventID) (*event.Event[any], error)
	Save(context.Context, *event.Event[any]) error
}
