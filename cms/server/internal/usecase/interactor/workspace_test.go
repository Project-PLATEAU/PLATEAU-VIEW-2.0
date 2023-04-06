package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestWorkspace_Create(t *testing.T) {
	ctx := context.Background()

	db := memory.New()

	u := user.New().NewID().Name("aaa").Email("aaa@bbb.com").Workspace(id.NewWorkspaceID()).MustBuild()
	workspaceUC := NewWorkspace(db, nil)
	op := &usecase.Operator{User: lo.ToPtr(u.ID())}
	workspace, err := workspaceUC.Create(ctx, "workspace name", u.ID(), op)

	assert.NoError(t, err)
	assert.NotNil(t, workspace)

	resultWorkspaces, _ := workspaceUC.Fetch(ctx, []id.WorkspaceID{workspace.ID()}, &usecase.Operator{
		ReadableWorkspaces: []id.WorkspaceID{workspace.ID()},
	})

	assert.NotNil(t, resultWorkspaces)
	assert.NotEmpty(t, resultWorkspaces)
	assert.Equal(t, resultWorkspaces[0].ID(), workspace.ID())
	assert.Equal(t, resultWorkspaces[0].Name(), "workspace name")
	assert.Equal(t, user.WorkspaceIDList{resultWorkspaces[0].ID()}, op.OwningWorkspaces)

	// mock workspace error
	wantErr := errors.New("test")
	memory.SetWorkspaceError(db.Workspace, wantErr)
	workspace2, err := workspaceUC.Create(ctx, "workspace name 2", u.ID(), op)
	assert.Nil(t, workspace2)
	assert.Equal(t, wantErr, err)
}

func TestWorkspace_Fetch(t *testing.T) {
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).MustBuild()

	u := user.New().NewID().Name("aaa").Email("aaa@bbb.com").Workspace(id1).MustBuild()
	op := &usecase.Operator{
		User:               lo.ToPtr(u.ID()),
		ReadableWorkspaces: []id.WorkspaceID{id1, id2},
	}

	tests := []struct {
		name  string
		seeds []*user.Workspace
		args  struct {
			ids      []id.WorkspaceID
			operator *usecase.Operator
		}
		want             []*user.Workspace
		mockWorkspaceErr bool
		wantErr          error
	}{
		{
			name:  "Fetch 1 of 2",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				ids      []id.WorkspaceID
				operator *usecase.Operator
			}{
				ids:      []id.WorkspaceID{id1},
				operator: op,
			},
			want:    []*user.Workspace{w1},
			wantErr: nil,
		},
		{
			name:  "Fetch 2 of 2",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				ids      []id.WorkspaceID
				operator *usecase.Operator
			}{
				ids:      []id.WorkspaceID{id1, id2},
				operator: op,
			},
			want:    []*user.Workspace{w1, w2},
			wantErr: nil,
		},
		{
			name:  "Fetch 1 of 0",
			seeds: []*user.Workspace{},
			args: struct {
				ids      []id.WorkspaceID
				operator *usecase.Operator
			}{
				ids:      []id.WorkspaceID{id1},
				operator: op,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:  "Fetch 2 of 0",
			seeds: []*user.Workspace{},
			args: struct {
				ids      []id.WorkspaceID
				operator *usecase.Operator
			}{
				ids:      []id.WorkspaceID{id1, id2},
				operator: op,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:             "mock error",
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.Fetch(ctx, tc.args.ids, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWorkspace_FindByUser(t *testing.T) {
	userID := id.NewUserID()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Members(map[user.ID]user.Member{userID: {Role: user.RoleReader}}).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).MustBuild()

	u := user.New().NewID().Name("aaa").Email("aaa@bbb.com").Workspace(id1).MustBuild()
	op := &usecase.Operator{
		User:               lo.ToPtr(u.ID()),
		ReadableWorkspaces: []id.WorkspaceID{id1, id2},
	}

	tests := []struct {
		name  string
		seeds []*user.Workspace
		args  struct {
			userID   id.UserID
			operator *usecase.Operator
		}
		want             []*user.Workspace
		mockWorkspaceErr bool
		wantErr          error
	}{
		{
			name:  "Fetch 1 of 2",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				userID   id.UserID
				operator *usecase.Operator
			}{
				userID:   userID,
				operator: op,
			},
			want:    []*user.Workspace{w1},
			wantErr: nil,
		},
		{
			name:  "Fetch 1 of 0",
			seeds: []*user.Workspace{},
			args: struct {
				userID   id.UserID
				operator *usecase.Operator
			}{
				userID:   userID,
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "Fetch 0 of 1",
			seeds: []*user.Workspace{w2},
			args: struct {
				userID   id.UserID
				operator *usecase.Operator
			}{
				userID:   userID,
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:             "mock error",
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.FindByUser(ctx, tc.args.userID, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWorkspace_Update(t *testing.T) {
	userID := id.NewUserID()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Name("W1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	w1Updated := user.NewWorkspace().ID(id1).Name("WW1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).Name("W2").MustBuild()
	id3 := id.NewWorkspaceID()
	w3 := user.NewWorkspace().ID(id3).Name("W3").Members(map[user.ID]user.Member{userID: {Role: user.RoleReader}}).MustBuild()

	op := &usecase.Operator{
		User:               &userID,
		ReadableWorkspaces: []id.WorkspaceID{id1, id2, id3},
		OwningWorkspaces:   []id.WorkspaceID{id1},
	}

	tests := []struct {
		name  string
		seeds []*user.Workspace
		args  struct {
			wId      id.WorkspaceID
			newName  string
			operator *usecase.Operator
		}
		want             *user.Workspace
		wantErr          error
		mockWorkspaceErr bool
	}{
		{
			name:  "Update 1",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				wId      id.WorkspaceID
				newName  string
				operator *usecase.Operator
			}{
				wId:      id1,
				newName:  "WW1",
				operator: op,
			},
			want:    w1Updated,
			wantErr: nil,
		},
		{
			name:  "Update 2",
			seeds: []*user.Workspace{},
			args: struct {
				wId      id.WorkspaceID
				newName  string
				operator *usecase.Operator
			}{
				wId:      id2,
				newName:  "WW2",
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "Update 3",
			seeds: []*user.Workspace{w3},
			args: struct {
				wId      id.WorkspaceID
				newName  string
				operator *usecase.Operator
			}{
				wId:      id3,
				newName:  "WW3",
				operator: op,
			},
			want:    nil,
			wantErr: interfaces.ErrOperationDenied,
		},
		{
			name: "mock error",
			args: struct {
				wId      id.WorkspaceID
				newName  string
				operator *usecase.Operator
			}{
				operator: op,
			},
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.Update(ctx, tc.args.wId, tc.args.newName, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
			got2, err := db.Workspace.FindByID(ctx, tc.args.wId)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got2)
		})
	}
}

func TestWorkspace_Remove(t *testing.T) {
	userID := id.NewUserID()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Name("W1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).Name("W2").MustBuild()
	id3 := id.NewWorkspaceID()
	w3 := user.NewWorkspace().ID(id3).Name("W3").Members(map[user.ID]user.Member{userID: {Role: user.RoleReader}}).MustBuild()
	id4 := id.NewWorkspaceID()
	w4 := user.NewWorkspace().ID(id4).Name("W4").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(true).MustBuild()
	id5 := id.NewWorkspaceID()
	w5 := user.NewWorkspace().ID(id5).Name("W5").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	p := project.New().NewID().Workspace(id5).MustBuild()
	id6 := id.NewWorkspaceID()
	w6 := user.NewWorkspace().ID(id6).Name("W6").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	p2 := project.New().NewID().Workspace(id6).MustBuild()

	op := &usecase.Operator{
		User:               &userID,
		ReadableWorkspaces: []id.WorkspaceID{id1, id2, id3},
		OwningWorkspaces:   []id.WorkspaceID{id1, id4, id5, id6},
	}

	tests := []struct {
		name  string
		seeds []*user.Workspace
		args  struct {
			wId      id.WorkspaceID
			operator *usecase.Operator
			project  *project.Project
		}
		wantErr          error
		mockWorkspaceErr bool
		mockProjectErr   bool
		want             *user.Workspace
	}{
		{
			name:  "Remove 1",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id1,
				operator: op,
			},
			wantErr: nil,
			want:    nil,
		},
		{
			name:  "Update 2",
			seeds: []*user.Workspace{w1, w2},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id2,
				operator: op,
			},
			wantErr: interfaces.ErrOperationDenied,
			want:    w2,
		},
		{
			name:  "Update 3",
			seeds: []*user.Workspace{w3},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id3,
				operator: op,
			},
			wantErr: interfaces.ErrOperationDenied,
			want:    w3,
		},
		{
			name:  "Remove 4",
			seeds: []*user.Workspace{w4},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id4,
				operator: op,
			},
			wantErr: user.ErrCannotModifyPersonalWorkspace,
			want:    w4,
		},
		{
			name:  "Remove 5: workspace that has a project",
			seeds: []*user.Workspace{w5},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id5,
				operator: op,
				project:  p,
			},
			wantErr: interfaces.ErrWorkspaceWithProjects,
			want:    nil,
		},
		{
			name: "mock workspace error",
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id5,
				operator: op,
			},
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
		{
			name:  "mock project count error",
			seeds: []*user.Workspace{w6},
			args: struct {
				wId      id.WorkspaceID
				operator *usecase.Operator
				project  *project.Project
			}{
				wId:      id6,
				operator: op,
				project:  p2,
			},
			wantErr:        errors.New("test2"),
			mockProjectErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			if tc.args.project != nil {
				err := db.Project.Save(ctx, tc.args.project)
				assert.NoError(t, err)
			}
			if tc.mockProjectErr {
				memory.SetProjectError(db.Project, tc.wantErr)
				projectCount, err := db.Project.CountByWorkspace(ctx, id.WorkspaceID{})
				assert.Equal(t, projectCount, 0)
				assert.NotNil(t, err)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)
			err := workspaceUC.Remove(ctx, tc.args.wId, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.NoError(t, err)
			got, err := db.Workspace.FindByID(ctx, tc.args.wId)
			if tc.want == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWorkspace_AddMember(t *testing.T) {
	userID := id.NewUserID()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Name("W1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).Name("W2").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	id3 := id.NewWorkspaceID()
	w3 := user.NewWorkspace().ID(id3).Name("W3").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(true).MustBuild()
	id4 := id.NewWorkspaceID()
	w4 := user.NewWorkspace().ID(id3).Name("W4").Members(map[user.ID]user.Member{id.NewUserID(): {Role: user.RoleOwner}}).Personal(true).MustBuild()

	u := user.New().NewID().Name("aaa").Email("a@b.c").MustBuild()

	op := &usecase.Operator{
		User:               &userID,
		ReadableWorkspaces: []id.WorkspaceID{id1, id2},
		OwningWorkspaces:   []id.WorkspaceID{id1, id2, id3},
	}

	tests := []struct {
		name       string
		seeds      []*user.Workspace
		usersSeeds []*user.User
		args       struct {
			wId      id.WorkspaceID
			users    map[id.UserID]user.Role
			operator *usecase.Operator
		}
		wantErr          error
		mockWorkspaceErr bool
		want             *user.Members
	}{
		{
			name:       "Add non existing",
			seeds:      []*user.Workspace{w1},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				users    map[id.UserID]user.Role
				operator *usecase.Operator
			}{
				wId:      id1,
				users:    map[id.UserID]user.Role{id.NewUserID(): user.RoleReader},
				operator: op,
			},
			want: user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:       "Add",
			seeds:      []*user.Workspace{w2},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				users    map[id.UserID]user.Role
				operator *usecase.Operator
			}{
				wId:      id2,
				users:    map[id.UserID]user.Role{u.ID(): user.RoleReader},
				operator: op,
			},
			wantErr: nil,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}, u.ID(): {Role: user.RoleReader, InvitedBy: userID}}),
		},
		{
			name:       "Add to personal workspace",
			seeds:      []*user.Workspace{w3},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				users    map[id.UserID]user.Role
				operator *usecase.Operator
			}{
				wId:      id3,
				users:    map[id.UserID]user.Role{u.ID(): user.RoleReader},
				operator: op,
			},
			wantErr: user.ErrCannotModifyPersonalWorkspace,
			want:    user.NewFixedMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:  "op denied",
			seeds: []*user.Workspace{w4},
			args: struct {
				wId      id.WorkspaceID
				users    map[id.UserID]user.Role
				operator *usecase.Operator
			}{
				wId:      id4,
				users:    map[id.UserID]user.Role{id.NewUserID(): user.RoleReader},
				operator: op,
			},
			wantErr:          interfaces.ErrOperationDenied,
			mockWorkspaceErr: false,
		},
		{
			name: "mock error",
			args: struct {
				wId      id.WorkspaceID
				users    map[id.UserID]user.Role
				operator *usecase.Operator
			}{
				wId:      id3,
				users:    map[id.UserID]user.Role{u.ID(): user.RoleReader},
				operator: op,
			},
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			for _, p := range tc.usersSeeds {
				err := db.User.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.AddUserMember(ctx, tc.args.wId, tc.args.users, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got.Members())

			got, err = db.Workspace.FindByID(ctx, tc.args.wId)
			if tc.want == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got.Members())
		})
	}
}

func TestWorkspace_RemoveMember(t *testing.T) {
	userID := id.NewUserID()
	u := user.New().NewID().Name("aaa").Email("a@b.c").MustBuild()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Name("W1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).Name("W2").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}, u.ID(): {Role: user.RoleReader}}).Personal(false).MustBuild()
	id3 := id.NewWorkspaceID()
	w3 := user.NewWorkspace().ID(id3).Name("W3").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(true).MustBuild()
	id4 := id.NewWorkspaceID()
	w4 := user.NewWorkspace().ID(id4).Name("W4").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()

	op := &usecase.Operator{
		User:               &userID,
		ReadableWorkspaces: []id.WorkspaceID{id1, id2},
		OwningWorkspaces:   []id.WorkspaceID{id1},
	}

	tests := []struct {
		name       string
		seeds      []*user.Workspace
		usersSeeds []*user.User
		args       struct {
			wId      id.WorkspaceID
			uId      id.UserID
			operator *usecase.Operator
		}
		wantErr          error
		mockWorkspaceErr bool
		want             *user.Members
	}{
		{
			name:       "Remove non existing",
			seeds:      []*user.Workspace{w1},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				operator *usecase.Operator
			}{
				wId:      id1,
				uId:      id.NewUserID(),
				operator: op,
			},
			wantErr: user.ErrTargetUserNotInTheWorkspace,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:       "Remove",
			seeds:      []*user.Workspace{w2},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				operator *usecase.Operator
			}{
				wId:      id2,
				uId:      u.ID(),
				operator: op,
			},
			wantErr: nil,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:       "Remove personal workspace",
			seeds:      []*user.Workspace{w3},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				operator *usecase.Operator
			}{
				wId:      id3,
				uId:      userID,
				operator: op,
			},
			wantErr: user.ErrCannotModifyPersonalWorkspace,
			want:    user.NewFixedMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:       "Remove single member",
			seeds:      []*user.Workspace{w4},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				operator *usecase.Operator
			}{
				wId:      id4,
				uId:      userID,
				operator: op,
			},
			wantErr: interfaces.ErrOwnerCannotLeaveTheWorkspace,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name: "mock error",
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				operator *usecase.Operator
			}{operator: op},
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			for _, p := range tc.usersSeeds {
				err := db.User.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.RemoveUser(ctx, tc.args.wId, tc.args.uId, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got.Members())

			got, err = db.Workspace.FindByID(ctx, tc.args.wId)
			if tc.want == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got.Members())
		})
	}
}

func TestWorkspace_UpdateMember(t *testing.T) {
	userID := id.NewUserID()
	u := user.New().NewID().Name("aaa").Email("a@b.c").MustBuild()
	id1 := id.NewWorkspaceID()
	w1 := user.NewWorkspace().ID(id1).Name("W1").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(false).MustBuild()
	id2 := id.NewWorkspaceID()
	w2 := user.NewWorkspace().ID(id2).Name("W2").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}, u.ID(): {Role: user.RoleReader}}).Personal(false).MustBuild()
	id3 := id.NewWorkspaceID()
	w3 := user.NewWorkspace().ID(id3).Name("W3").Members(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}).Personal(true).MustBuild()

	op := &usecase.Operator{
		User:               &userID,
		ReadableWorkspaces: []id.WorkspaceID{id1, id2},
		OwningWorkspaces:   []id.WorkspaceID{id1, id2, id3},
	}

	tests := []struct {
		name       string
		seeds      []*user.Workspace
		usersSeeds []*user.User
		args       struct {
			wId      id.WorkspaceID
			uId      id.UserID
			role     user.Role
			operator *usecase.Operator
		}
		wantErr          error
		mockWorkspaceErr bool
		want             *user.Members
	}{
		{
			name:       "Update non existing",
			seeds:      []*user.Workspace{w1},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				role     user.Role
				operator *usecase.Operator
			}{
				wId:      id1,
				uId:      id.NewUserID(),
				role:     user.RoleWriter,
				operator: op,
			},
			wantErr: user.ErrTargetUserNotInTheWorkspace,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name:       "Update",
			seeds:      []*user.Workspace{w2},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				role     user.Role
				operator *usecase.Operator
			}{
				wId:      id2,
				uId:      u.ID(),
				role:     user.RoleWriter,
				operator: op,
			},
			wantErr: nil,
			want:    user.NewMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}, u.ID(): {Role: user.RoleWriter}}),
		},
		{
			name:       "Update personal workspace",
			seeds:      []*user.Workspace{w3},
			usersSeeds: []*user.User{u},
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				role     user.Role
				operator *usecase.Operator
			}{
				wId:      id3,
				uId:      userID,
				role:     user.RoleReader,
				operator: op,
			},
			wantErr: user.ErrCannotModifyPersonalWorkspace,
			want:    user.NewFixedMembersWith(map[user.ID]user.Member{userID: {Role: user.RoleOwner}}),
		},
		{
			name: "mock error",
			args: struct {
				wId      id.WorkspaceID
				uId      id.UserID
				role     user.Role
				operator *usecase.Operator
			}{
				wId:      id3,
				operator: op,
			},
			wantErr:          errors.New("test"),
			mockWorkspaceErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockWorkspaceErr {
				memory.SetWorkspaceError(db.Workspace, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Workspace.Save(ctx, p)
				assert.NoError(t, err)
			}
			for _, p := range tc.usersSeeds {
				err := db.User.Save(ctx, p)
				assert.NoError(t, err)
			}
			workspaceUC := NewWorkspace(db, nil)

			got, err := workspaceUC.UpdateUser(ctx, tc.args.wId, tc.args.uId, tc.args.role, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got.Members())

			got, err = db.Workspace.FindByID(ctx, tc.args.wId)
			if tc.want == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got.Members())
		})
	}
}
