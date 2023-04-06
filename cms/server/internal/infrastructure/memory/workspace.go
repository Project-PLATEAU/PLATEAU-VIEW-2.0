package memory

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"golang.org/x/exp/slices"
)

type Workspace struct {
	data *util.SyncMap[id.WorkspaceID, *user.Workspace]
	err  error
}

func NewWorkspace() repo.Workspace {
	return &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
}

func (r *Workspace) FindByUser(ctx context.Context, i id.UserID) (user.WorkspaceList, error) {
	if r.err != nil {
		return nil, r.err
	}

	return rerror.ErrIfNil(r.data.FindAll(func(key id.WorkspaceID, value *user.Workspace) bool {
		return value.Members().HasUser(i)
	}), rerror.ErrNotFound)
}

func (r *Workspace) FindByIntegration(_ context.Context, i id.IntegrationID) (user.WorkspaceList, error) {
	if r.err != nil {
		return nil, r.err
	}

	return rerror.ErrIfNil(r.data.FindAll(func(key id.WorkspaceID, value *user.Workspace) bool {
		return value.Members().HasIntegration(i)
	}), rerror.ErrNotFound)
}

func (r *Workspace) FindByIDs(ctx context.Context, ids id.WorkspaceIDList) (user.WorkspaceList, error) {
	if r.err != nil {
		return nil, r.err
	}

	res := r.data.FindAll(func(key id.WorkspaceID, value *user.Workspace) bool {
		return ids.Has(key)
	})
	slices.SortFunc(res, func(a, b *user.Workspace) bool { return a.ID().Compare(b.ID()) < 0 })
	return res, nil
}

func (r *Workspace) FindByID(ctx context.Context, v id.WorkspaceID) (*user.Workspace, error) {
	if r.err != nil {
		return nil, r.err
	}

	return rerror.ErrIfNil(r.data.Find(func(key id.WorkspaceID, value *user.Workspace) bool {
		return key == v
	}), rerror.ErrNotFound)
}

func (r *Workspace) Save(ctx context.Context, t *user.Workspace) error {
	if r.err != nil {
		return r.err
	}

	r.data.Store(t.ID(), t)
	return nil
}

func (r *Workspace) SaveAll(ctx context.Context, workspaces []*user.Workspace) error {
	if r.err != nil {
		return r.err
	}

	for _, t := range workspaces {
		r.data.Store(t.ID(), t)
	}
	return nil
}

func (r *Workspace) Remove(ctx context.Context, wid id.WorkspaceID) error {
	if r.err != nil {
		return r.err
	}

	r.data.Delete(wid)
	return nil
}

func (r *Workspace) RemoveAll(ctx context.Context, ids id.WorkspaceIDList) error {
	if r.err != nil {
		return r.err
	}

	for _, wid := range ids {
		r.data.Delete(wid)
	}
	return nil
}

func SetWorkspaceError(r repo.Workspace, err error) {
	r.(*Workspace).err = err
}
