package memory

import (
	"context"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type Model struct {
	data *util.SyncMap[id.ModelID, *model.Model]
	f    repo.ProjectFilter
	now  *util.TimeNow
	err  error
}

func NewModel() repo.Model {
	return &Model{
		data: &util.SyncMap[id.ModelID, *model.Model]{},
		now:  &util.TimeNow{},
	}
}

func (r *Model) Filtered(f repo.ProjectFilter) repo.Model {
	return &Model{
		data: r.data,
		f:    r.f.Merge(f),
		now:  &util.TimeNow{},
	}
}

func (r *Model) FindByProject(_ context.Context, pid id.ProjectID, _ *usecasex.Pagination) (model.List, *usecasex.PageInfo, error) {
	if r.err != nil {
		return nil, nil, r.err
	}

	// TODO: implement pagination

	if !r.f.CanRead(pid) {
		return nil, nil, nil
	}

	result := model.List(r.data.FindAll(func(_ id.ModelID, m *model.Model) bool {
		return m.Project() == pid
	})).SortByID()

	var startCursor, endCursor *usecasex.Cursor
	if len(result) > 0 {
		startCursor = lo.ToPtr(usecasex.Cursor(result[0].ID().String()))
		endCursor = lo.ToPtr(usecasex.Cursor(result[len(result)-1].ID().String()))
	}

	return result, usecasex.NewPageInfo(
		int64(len(result)),
		startCursor,
		endCursor,
		true,
		true,
	), nil
}

func (r *Model) CountByProject(_ context.Context, pid id.ProjectID) (int, error) {
	if r.err != nil {
		return 0, r.err
	}

	if !r.f.CanRead(pid) {
		return 0, nil
	}

	return r.data.CountAll(func(_ id.ModelID, m *model.Model) bool {
		return m.Project() == pid
	}), nil
}

func (r *Model) FindByKey(_ context.Context, pid id.ProjectID, key string) (*model.Model, error) {
	if r.err != nil {
		return nil, r.err
	}

	if !r.f.CanRead(pid) {
		return nil, nil
	}

	m := r.data.Find(func(_ id.ModelID, m *model.Model) bool {
		return m.Key().String() == key && m.Project() == pid
	})
	if m == nil {
		return nil, rerror.ErrNotFound
	}

	return m, nil
}

func (r *Model) FindByIDOrKey(ctx context.Context, projectID id.ProjectID, q model.IDOrKey) (*model.Model, error) {
	if r.err != nil {
		return nil, r.err
	}

	modelID := q.ID()
	key := q.Key()
	if modelID == nil && (key == nil || *key == "") {
		return nil, rerror.ErrNotFound
	}

	m := r.data.Find(func(_ id.ModelID, m *model.Model) bool {
		return r.f.CanRead(m.Project()) && (modelID != nil && m.ID() == *modelID || key != nil && m.Key().String() == *key)
	})
	if m == nil {
		return nil, rerror.ErrNotFound
	}

	return m, nil
}

func (r *Model) FindByID(_ context.Context, mid id.ModelID) (*model.Model, error) {
	if r.err != nil {
		return nil, r.err
	}

	m := r.data.Find(func(k id.ModelID, m *model.Model) bool {
		return k == mid && r.f.CanRead(m.Project())
	})

	if m != nil {
		return m, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Model) FindByIDs(_ context.Context, ids id.ModelIDList) (model.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	result := r.data.FindAll(func(k id.ModelID, m *model.Model) bool {
		return ids.Has(k) && r.f.CanRead(m.Project())
	})

	return model.List(result).SortByID(), nil
}

func (r *Model) Save(_ context.Context, m *model.Model) error {
	if r.err != nil {
		return r.err
	}

	if !r.f.CanWrite(m.Project()) {
		return repo.ErrOperationDenied
	}

	m.SetUpdatedAt(r.now.Now())
	r.data.Store(m.ID(), m)
	return nil
}

func (r *Model) Remove(_ context.Context, mId id.ModelID) error {
	if r.err != nil {
		return r.err
	}

	if m, ok := r.data.Load(mId); ok && r.f.CanWrite(m.Project()) {
		r.data.Delete(mId)
		return nil
	}
	return rerror.ErrNotFound
}

func MockModelNow(r repo.Model, t time.Time) func() {
	return r.(*Model).now.Mock(t)
}

func SetModelError(r repo.Model, err error) {
	r.(*Model).err = err
}
