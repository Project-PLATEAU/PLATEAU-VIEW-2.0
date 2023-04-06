package mongo

import (
	"context"
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestWorkspace_FindByID(t *testing.T) {
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	tests := []struct {
		Name               string
		Input              id.WorkspaceID
		RepoData, Expected *user.Workspace
		WantErr            bool
	}{
		{
			Name:     "must find a workspace",
			Input:    ws.ID(),
			RepoData: ws,
			Expected: ws,
		},
		{
			Name:     "must not find any workspace",
			Input:    user.NewWorkspaceID(),
			RepoData: ws,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewWorkspace(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByID(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
			}
		})
	}
}

func TestWorkspace_FindByIDs(t *testing.T) {
	ws1 := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()
	ws3 := user.NewWorkspace().NewID().Name("xxx").MustBuild()

	tests := []struct {
		Name               string
		Input              id.WorkspaceIDList
		RepoData, Expected user.WorkspaceList
	}{
		{
			Name:     "must find users",
			RepoData: user.WorkspaceList{ws1, ws2},
			Input:    id.WorkspaceIDList{ws1.ID(), ws2.ID()},
			Expected: user.WorkspaceList{ws1, ws2},
		},
		{
			Name:     "must not find any user",
			Input:    id.WorkspaceIDList{ws3.ID()},
			RepoData: user.WorkspaceList{ws2, ws1},
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewWorkspace(client)
			ctx := context.Background()
			err := repo.SaveAll(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByIDs(ctx, tc.Input)
			assert.NoError(tt, err)
			for k, ws := range got {
				if ws != nil {
					assert.Equal(tt, tc.Expected[k].ID(), ws.ID())
					assert.Equal(tt, tc.Expected[k].Name(), ws.Name())
				}
			}
		})
	}
}

func TestWorkspace_FindByUser(t *testing.T) {
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").MustBuild()
	ws := user.NewWorkspace().NewID().Name("hoge").Members(map[user.ID]user.Member{u.ID(): {Role: user.RoleOwner, InvitedBy: u.ID()}}).MustBuild()
	tests := []struct {
		Name     string
		Input    id.UserID
		RepoData *user.Workspace
		Expected user.WorkspaceList
	}{
		{
			Name:     "must find a workspace",
			Input:    u.ID(),
			RepoData: ws,
			Expected: user.WorkspaceList{ws},
		},
		{
			Name:     "must not find any workspace",
			Input:    user.NewID(),
			RepoData: ws,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewWorkspace(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByUser(ctx, tc.Input)
			assert.NoError(tt, err)
			for k, ws := range got {
				if ws != nil {
					assert.Equal(tt, tc.Expected[k].ID(), ws.ID())
					assert.Equal(tt, tc.Expected[k].Name(), ws.Name())
				}
			}
		})
	}
}

func TestWorkspace_Remove(t *testing.T) {
	ws := user.NewWorkspace().NewID().Name("hoge").MustBuild()

	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))

	repo := NewWorkspace(client)
	ctx := context.Background()
	err := repo.Save(ctx, ws)
	assert.NoError(t, err)

	err = repo.Remove(ctx, ws.ID())
	assert.NoError(t, err)
}

func TestWorkspace_RemoveAll(t *testing.T) {
	ws1 := user.NewWorkspace().NewID().Name("hoge").MustBuild()
	ws2 := user.NewWorkspace().NewID().Name("foo").MustBuild()

	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))

	repo := NewWorkspace(client)
	ctx := context.Background()
	err := repo.SaveAll(ctx, user.WorkspaceList{ws1, ws2})
	assert.NoError(t, err)

	err = repo.RemoveAll(ctx, user.WorkspaceIDList{ws1.ID(), ws2.ID()})
	assert.NoError(t, err)
}
