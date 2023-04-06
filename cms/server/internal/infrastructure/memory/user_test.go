package memory

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	expected := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}

	got := NewUser()
	assert.Equal(t, expected, got)
}

func TestUser_FindBySub(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").Auths([]user.Auth{{
		Sub: "xxx",
	}}).MustBuild()

	tests := []struct {
		name    string
		sub     string
		want    *user.User
		wantErr error
		mockErr bool
	}{
		{
			name: "must find user by auth",
			sub:  "xxx",
			want: u,
		},
		{
			name:    "must return ErrInvalidParams",
			sub:     "",
			wantErr: rerror.ErrInvalidParams,
		},
		{
			name:    "must mock error",
			wantErr: errors.New("test"),
			mockErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewUser()
			_ = r.Save(ctx, u.Clone())
			if tc.mockErr {
				SetUserError(r, tc.wantErr)
			}

			got, err := r.FindBySub(ctx, tc.sub)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestUser_FindByEmail(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").MustBuild()
	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	r.data.Store(u.ID(), u)
	out, err := r.FindByEmail(ctx, "aa@bb.cc")
	assert.NoError(t, err)
	assert.Equal(t, u, out)

	out, err = r.FindByEmail(ctx, "abc@bb.cc")
	assert.Same(t, rerror.ErrNotFound, err)
	assert.Nil(t, out)

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	_, err = r.FindByEmail(ctx, "")
	assert.Same(t, wantErr, err)
}

func TestUser_FindByIDs(t *testing.T) {
	ctx := context.Background()
	u1 := user.New().NewID().Name("hoge").Email("abc@bb.cc").MustBuild()
	u2 := user.New().NewID().Name("foo").Email("cba@bb.cc").MustBuild()
	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	r.data.Store(u1.ID(), u1)
	r.data.Store(u2.ID(), u2)

	ids := id.UserIDList{
		u1.ID(),
		u2.ID(),
	}
	out, err := r.FindByIDs(ctx, ids)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(out))

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	_, err = r.FindByIDs(ctx, ids)
	assert.Same(t, wantErr, err)
}

func TestUser_FindByName(t *testing.T) {
	ctx := context.Background()
	pr := user.PasswordReset{
		Token: "123abc",
	}
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").PasswordReset(pr.Clone()).MustBuild()

	tests := []struct {
		name    string
		seeds   []*user.User
		uName   string
		want    *user.User
		wantErr error
		mockErr bool
	}{
		{
			name:  "must find user by name",
			seeds: []*user.User{u},
			uName: "hoge",
			want:  u,
		},
		{
			name:    "must return ErrInvalidParams",
			wantErr: rerror.ErrInvalidParams,
		},
		{
			name:    "must return ErrNotFound",
			uName:   "xxx",
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "must mock error",
			wantErr: errors.New("test"),
			mockErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewUser()
			if tc.mockErr {
				SetUserError(r, tc.wantErr)
			}
			for _, u := range tc.seeds {
				_ = r.Save(ctx, u.Clone())
			}

			got, err := r.FindByName(ctx, tc.uName)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestUser_FindByNameOrEmail(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").MustBuild()
	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	r.data.Store(u.ID(), u)

	out, err := r.FindByNameOrEmail(ctx, "hoge")
	assert.NoError(t, err)
	assert.Equal(t, u, out)

	out2, err := r.FindByNameOrEmail(ctx, "aa@bb.cc")
	assert.NoError(t, err)
	assert.Equal(t, u, out2)

	out3, err := r.FindByNameOrEmail(ctx, "xxx")
	assert.Nil(t, out3)
	assert.Same(t, rerror.ErrNotFound, err)

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	_, err = r.FindByID(ctx, u.ID())
	assert.Same(t, wantErr, err)
}

func TestUser_FindByPasswordResetRequest(t *testing.T) {
	ctx := context.Background()
	pr := user.PasswordReset{
		Token: "123abc",
	}
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").PasswordReset(pr.Clone()).MustBuild()

	tests := []struct {
		name    string
		seeds   []*user.User
		token   string
		want    *user.User
		wantErr error
		mockErr bool
	}{
		{
			name:    "must find user by password reset",
			seeds:   []*user.User{u},
			token:   pr.Token,
			want:    u,
			wantErr: nil,
		},
		{
			name:    "must return ErrInvalidParams",
			seeds:   []*user.User{u},
			wantErr: rerror.ErrInvalidParams,
		},
		{
			name:    "must return ErrNotFound",
			seeds:   []*user.User{u},
			token:   "xxx",
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "must mock error",
			seeds:   []*user.User{u},
			wantErr: errors.New("test"),
			mockErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewUser()
			if tc.mockErr {
				SetUserError(r, tc.wantErr)
			}
			for _, u := range tc.seeds {
				_ = r.Save(ctx, u.Clone())
			}

			got, err := r.FindByPasswordResetRequest(ctx, tc.token)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestUser_FindByVerification(t *testing.T) {
	ctx := context.Background()
	vr := user.VerificationFrom("123abc", time.Now(), false)
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").Verification(vr).MustBuild()

	tests := []struct {
		name    string
		seeds   []*user.User
		code    string
		want    *user.User
		wantErr error
		mockErr bool
	}{
		{
			name:    "must find user by verification",
			seeds:   []*user.User{u},
			code:    "123abc",
			want:    u,
			wantErr: nil,
		},
		{
			name:    "must return ErrInvalidParams",
			seeds:   []*user.User{u},
			wantErr: rerror.ErrInvalidParams,
		},
		{
			name:    "must return ErrNotFound",
			seeds:   []*user.User{u},
			code:    "xxx",
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "must mock error",
			seeds:   []*user.User{u},
			wantErr: errors.New("test"),
			mockErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewUser()
			if tc.mockErr {
				SetUserError(r, tc.wantErr)
			}
			for _, u := range tc.seeds {
				_ = r.Save(ctx, u.Clone())
			}

			got, err := r.FindByVerification(ctx, tc.code)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestUser_FindByID(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").MustBuild()
	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	r.data.Store(u.ID(), u)

	out, err := r.FindByID(ctx, u.ID())
	assert.NoError(t, err)
	assert.Equal(t, u, out)

	out2, err := r.FindByID(ctx, id.UserID{})
	assert.Nil(t, out2)
	assert.Same(t, rerror.ErrNotFound, err)

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	_, err = r.FindByID(ctx, u.ID())
	assert.Same(t, wantErr, err)
}

func TestUser_FindBySubOrCreate(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").Auths([]user.Auth{{Sub: "auth0|aaa", Provider: "auth0"}}).MustBuild()

	r := &User{data: &util.SyncMap[id.UserID, *user.User]{}}

	_, err := r.FindBySubOrCreate(ctx, u, "auth0|aaa")
	assert.NoError(t, err)
	assert.Equal(t, 1, r.data.Len())

	// if same sub, it returns existing data in stead of inserting new data
	_, err = r.FindBySubOrCreate(ctx, u, "auth0|aaa")
	assert.NoError(t, err)
	assert.Equal(t, 1, r.data.Len())
}

func TestUser_Create(t *testing.T) {
	uid := id.NewUserID()
	ctx := context.Background()
	u := user.New().ID(uid).Name("hoge").Email("aa@bb.cc").Auths([]user.Auth{{Sub: "auth0|aaa", Provider: "auth0"}}).MustBuild()

	r := &User{data: &util.SyncMap[id.UserID, *user.User]{}}

	err := r.Create(ctx, u)
	assert.NoError(t, err)
	assert.Equal(t, 1, r.data.Len())

	err = r.Create(ctx, u)
	assert.Equal(t, repo.ErrDuplicatedUser, err)
}

func TestUser_Save(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").MustBuild()

	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	_ = r.Save(ctx, u)

	assert.Equal(t, 1, r.data.Len())

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, u))
}

func TestUser_Remove(t *testing.T) {
	ctx := context.Background()
	u := user.New().NewID().Name("hoge").Email("aa@bb.cc").MustBuild()
	u2 := user.New().NewID().Name("xxx").Email("abc@bb.cc").MustBuild()
	r := &User{
		data: &util.SyncMap[id.UserID, *user.User]{},
	}
	r.data.Store(u.ID(), u)
	r.data.Store(u2.ID(), u2)

	_ = r.Remove(ctx, u2.ID())
	assert.Equal(t, 1, r.data.Len())

	wantErr := errors.New("test")
	SetUserError(r, wantErr)
	assert.Same(t, wantErr, r.Remove(ctx, u.ID()))
}
