package memory

import (
	"context"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

type Integration struct {
	data *util.SyncMap[id.IntegrationID, *integration.Integration]
	now  *util.TimeNow
	err  error
}

func NewIntegration() repo.Integration {
	return &Integration{
		data: &util.SyncMap[id.IntegrationID, *integration.Integration]{},
		now:  &util.TimeNow{},
	}
}

func (r *Integration) FindByID(_ context.Context, iId id.IntegrationID) (*integration.Integration, error) {
	if r.err != nil {
		return nil, r.err
	}

	i := r.data.Find(func(k id.IntegrationID, i *integration.Integration) bool {
		return k == iId
	})

	if i != nil {
		return i, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Integration) FindByToken(_ context.Context, token string) (*integration.Integration, error) {
	if r.err != nil {
		return nil, r.err
	}

	i := r.data.Find(func(_ id.IntegrationID, i *integration.Integration) bool {
		return i.Token() == token
	})

	if i != nil {
		return i, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Integration) FindByIDs(_ context.Context, iIds id.IntegrationIDList) (integration.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	result := r.data.FindAll(func(k id.IntegrationID, i *integration.Integration) bool {
		return iIds.Has(k)
	})

	return integration.List(result).SortByID(), nil
}

func (r *Integration) FindByUser(_ context.Context, uID id.UserID) (integration.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	result := r.data.FindAll(func(k id.IntegrationID, i *integration.Integration) bool {
		return i.Developer() == uID
	})

	return integration.List(result).SortByID(), nil
}

func (r *Integration) Save(_ context.Context, i *integration.Integration) error {
	if r.err != nil {
		return r.err
	}

	r.data.Store(i.ID(), i)
	return nil
}

func (r *Integration) Remove(_ context.Context, iId id.IntegrationID) error {
	if r.err != nil {
		return r.err
	}

	if _, ok := r.data.Load(iId); ok {
		r.data.Delete(iId)
		return nil
	}
	return rerror.ErrNotFound
}

func MockIntegrationNow(r repo.Integration, t time.Time) func() {
	return r.(*Integration).now.Mock(t)
}

func SetIntegrationError(r repo.Integration, err error) {
	r.(*Integration).err = err
}
