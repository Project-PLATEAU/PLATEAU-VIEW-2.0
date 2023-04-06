package memory

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

type Event struct {
	data *util.SyncMap[id.EventID, *event.Event[any]]
	err  error
}

func NewEvent() repo.Event {
	return &Event{
		data: &util.SyncMap[id.EventID, *event.Event[any]]{},
	}
}

func (r *Event) FindByID(_ context.Context, iId id.EventID) (*event.Event[any], error) {
	if r.err != nil {
		return nil, r.err
	}

	i := r.data.Find(func(k id.EventID, i *event.Event[any]) bool {
		return k == iId
	})

	if i != nil {
		return i, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Event) Save(ctx context.Context, ev *event.Event[any]) error {
	if r.err != nil {
		return r.err
	}

	r.data.Store(ev.ID(), ev)
	return nil
}
