package memory

import (
	"context"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

type Schema struct {
	data *util.SyncMap[id.SchemaID, *schema.Schema]
	f    repo.WorkspaceFilter
	now  *util.TimeNow
	err  error
}

func NewSchema() repo.Schema {
	return &Schema{
		data: &util.SyncMap[id.SchemaID, *schema.Schema]{},
		now:  &util.TimeNow{},
	}
}

func (r *Schema) Filtered(f repo.WorkspaceFilter) repo.Schema {
	return &Schema{
		data: r.data,
		f:    r.f.Merge(f),
		now:  &util.TimeNow{},
	}
}

func (r *Schema) FindByID(_ context.Context, sid id.SchemaID) (*schema.Schema, error) {
	if r.err != nil {
		return nil, r.err
	}

	s := r.data.Find(func(k id.SchemaID, s *schema.Schema) bool {
		return k == sid && r.f.CanRead(s.Workspace())
	})

	if s != nil {
		return s, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Schema) FindByIDs(_ context.Context, ids id.SchemaIDList) (schema.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	result := r.data.FindAll(func(k id.SchemaID, s *schema.Schema) bool {
		return ids.Has(k) && r.f.CanRead(s.Workspace())
	})

	return schema.List(result).SortByID(), nil
}

func (r *Schema) Save(_ context.Context, s *schema.Schema) error {
	if r.err != nil {
		return r.err
	}

	if !r.f.CanWrite(s.Workspace()) {
		return repo.ErrOperationDenied
	}

	r.data.Store(s.ID(), s)
	return nil
}

func (r *Schema) Remove(_ context.Context, sId id.SchemaID) error {
	if r.err != nil {
		return r.err
	}

	if s, ok := r.data.Load(sId); ok && r.f.CanWrite(s.Workspace()) {
		r.data.Delete(sId)
		return nil
	}
	return rerror.ErrNotFound
}

func MockSchemaNow(r repo.Schema, t time.Time) func() {
	return r.(*Schema).now.Mock(t)
}

func SetSchemaError(r repo.Schema, err error) {
	r.(*Schema).err = err
}
