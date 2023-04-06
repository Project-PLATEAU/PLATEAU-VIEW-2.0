package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_FindByID(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              id.UserID
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    user1.ID(),
			RepoData: user1,
			Expected: user1,
		},
		{
			Name:     "must not find any user",
			Input:    user.NewID(),
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByID(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindByIDs(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()
	user2 := user.New().
		NewID().
		Email("aa2@bb.cc").
		Workspace(wsid).
		Name("hoge").
		MustBuild()
	user3 := user.New().
		NewID().
		Email("aa3@bb.cc").
		Workspace(wsid).
		Name("xxx").
		MustBuild()

	tests := []struct {
		Name               string
		Input              id.UserIDList
		RepoData, Expected []*user.User
	}{
		{
			Name:     "must find users",
			RepoData: []*user.User{user1, user2},
			Input: id.UserIDList{
				user1.ID(),
				user2.ID(),
			},
			Expected: []*user.User{user1, user2},
		},
		{
			Name:     "must not find any user",
			Input:    id.UserIDList{user3.ID()},
			RepoData: []*user.User{user1, user2},
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			for _, u := range tc.RepoData {
				err := repo.Save(ctx, u)
				assert.NoError(tt, err)
			}

			got, err := repo.FindByIDs(ctx, tc.Input)
			assert.NoError(tt, err)
			for k, u := range got {
				if u != nil {
					assert.Equal(tt, tc.Expected[k].ID(), u.ID())
					assert.Equal(tt, tc.Expected[k].Email(), u.Email())
					assert.Equal(tt, tc.Expected[k].Name(), u.Name())
					assert.Equal(tt, tc.Expected[k].Workspace(), u.Workspace())
				}
			}
		})
	}
}

func TestUserRepo_FindByName(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    user1.Name(),
			RepoData: user1,
			Expected: user1,
		},
		{
			Name:     "must not find any user",
			Input:    "xxx",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByName(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    user1.Email(),
			RepoData: user1,
			Expected: user1,
		},
		{
			Name:     "must not find any user",
			Input:    "xx@yy.zz",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByEmail(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindByNameOrEmail(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user by email",
			Input:    user1.Email(),
			RepoData: user1,
			Expected: user1,
		},
		{
			Name:     "must find a user by name",
			Input:    user1.Name(),
			RepoData: user1,
			Expected: user1,
		},
		{
			Name:     "must not find any user",
			Input:    "xx@yy.zz",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByNameOrEmail(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindByPasswordResetRequest(t *testing.T) {
	pr := user.PasswordReset{
		Token: "123abc",
	}
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		PasswordReset(pr.Clone()).
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    pr.Token,
			RepoData: user1,
			Expected: user1,
		},

		{
			Name:     "must not find any user",
			Input:    "x@yxz",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByPasswordResetRequest(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindByVerification(t *testing.T) {
	vr := user.VerificationFrom("123abc", time.Now(), false)

	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Verification(vr).
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    vr.Code(),
			RepoData: user1,
			Expected: user1,
		},

		{
			Name:     "must not find any user",
			Input:    "x@yxz",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByVerification(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_FindBySub(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Auths([]user.Auth{{
			Sub: "xxx",
		}}).
		Workspace(wsid).
		Name("foo").
		MustBuild()
	tests := []struct {
		Name               string
		Input              string
		RepoData, Expected *user.User
		WantErr            bool
	}{
		{
			Name:     "must find a user",
			Input:    "xxx",
			RepoData: user1,
			Expected: user1,
		},

		{
			Name:     "must not find any user",
			Input:    "x@yxz",
			RepoData: user1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewUser(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindBySub(ctx, tc.Input)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.ID())
				assert.Equal(tt, tc.Expected.Email(), got.Email())
				assert.Equal(tt, tc.Expected.Name(), got.Name())
				assert.Equal(tt, tc.Expected.Workspace(), got.Workspace())
			}
		})
	}
}

func TestUserRepo_Remove(t *testing.T) {
	wsid := user.NewWorkspaceID()
	user1 := user.New().
		NewID().
		Email("aa@bb.cc").
		Workspace(wsid).
		Name("foo").
		MustBuild()

	init := mongotest.Connect(t)

	client := mongox.NewClientWithDatabase(init(t))

	repo := NewUser(client)
	ctx := context.Background()
	err := repo.Save(ctx, user1)
	assert.NoError(t, err)

	err = repo.Remove(ctx, user1.ID())
	assert.NoError(t, err)
}
