package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/stretchr/testify/assert"
)

func TestNewWorkspace(t *testing.T) {
	expected := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	got := NewWorkspace()
	assert.Equal(t, expected, got)
}

func TestWorkspace_FindByID(t *testing.T) {
	ctx := context.Background()
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	r.data.Store(ws.ID(), ws)
	out, err := r.FindByID(ctx, ws.ID())
	assert.NoError(t, err)
	assert.Equal(t, ws, out)

	out2, err := r.FindByID(ctx, id.WorkspaceID{})
	assert.Nil(t, out2)
	assert.Same(t, rerror.ErrNotFound, err)

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, ws))
}

func TestWorkspace_FindByIDs(t *testing.T) {
	ctx := context.Background()
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()
	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	r.data.Store(ws.ID(), ws)
	r.data.Store(ws2.ID(), ws2)

	ids := user.WorkspaceIDList{ws.ID()}
	wsl := user.WorkspaceList{ws}
	out, err := r.FindByIDs(ctx, ids)
	assert.NoError(t, err)
	assert.Equal(t, wsl, out)

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, ws))
}

func TestWorkspace_FindByUser(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("aaa").Email("aaa@bbb.com").MustBuild()
	ws := user.NewWorkspace().NewID().Name("hoge").Members(map[user.ID]user.Member{u.ID(): {Role: user.RoleOwner}}).MustBuild()
	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	r.data.Store(ws.ID(), ws)
	wsl := user.WorkspaceList{ws}
	out, err := r.FindByUser(ctx, u.ID())
	assert.NoError(t, err)
	assert.Equal(t, wsl, out)

	out2, err := r.FindByUser(ctx, id.UserID{})
	assert.Same(t, rerror.ErrNotFound, err)
	assert.Nil(t, out2)

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, ws))
}

func TestWorkspace_Save(t *testing.T) {
	ctx := context.Background()
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()

	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	_ = r.Save(ctx, ws)
	assert.Equal(t, 1, r.data.Len())

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, ws))
}

func TestWorkspace_SaveAll(t *testing.T) {
	ctx := context.Background()
	ws1 := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()

	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	_ = r.SaveAll(ctx, []*user.Workspace{ws1, ws2})
	assert.Equal(t, 2, r.data.Len())

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Remove(ctx, ws1.ID()))
}

func TestWorkspace_Remove(t *testing.T) {
	ctx := context.Background()
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()
	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	r.data.Store(ws.ID(), ws)
	r.data.Store(ws2.ID(), ws2)

	_ = r.Remove(ctx, ws2.ID())
	assert.Equal(t, 1, r.data.Len())

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.Remove(ctx, ws.ID()))
}

func TestWorkspace_RemoveAll(t *testing.T) {
	ctx := context.Background()
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()
	r := &Workspace{
		data: &util.SyncMap[id.WorkspaceID, *user.Workspace]{},
	}
	r.data.Store(ws.ID(), ws)
	r.data.Store(ws2.ID(), ws2)

	ids := user.WorkspaceIDList{ws.ID(), ws2.ID()}
	_ = r.RemoveAll(ctx, ids)
	assert.Equal(t, 0, r.data.Len())

	wantErr := errors.New("test")
	SetWorkspaceError(r, wantErr)
	assert.Same(t, wantErr, r.RemoveAll(ctx, ids))
}
