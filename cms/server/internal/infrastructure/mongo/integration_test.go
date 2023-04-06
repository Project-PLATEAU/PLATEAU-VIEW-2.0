package mongo

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/rerror"
)

func testSuite() (now time.Time, uri *url.URL, uId id.UserID, iId1, iId2 id.IntegrationID, i1, i2 *integration.Integration) {
	now = time.Now().Truncate(time.Millisecond).UTC()
	uri = lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test"))
	uId = id.NewUserID()
	iId1 = id.NewIntegrationID()
	iId2 = id.NewIntegrationID()
	i1 = integration.New().ID(iId1).Name("i1").Webhook([]*integration.Webhook{}).Developer(uId).Type(integration.TypePrivate).LogoUrl(uri).UpdatedAt(now).MustBuild()
	i2 = integration.New().ID(iId2).Name("i2").Webhook([]*integration.Webhook{}).Developer(uId).Type(integration.TypePrivate).LogoUrl(uri).UpdatedAt(now).MustBuild()
	return
}

func TestIntegrationRepo_FindByID(t *testing.T) {
	now, u, uId, iId1, _, i1, _ := testSuite()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.ID
		want    *integration.Integration
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   integration.List{},
			arg:     integration.NewID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "Not found",
			seeds:   integration.List{i1},
			arg:     integration.NewID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "Found 1",
			seeds:   integration.List{i1},
			arg:     iId1,
			want:    i1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: integration.List{
				integration.New().NewID().Developer(uId).LogoUrl(u).UpdatedAt(now).MustBuild(),
				i1,
				integration.New().NewID().Developer(uId).LogoUrl(u).UpdatedAt(now).MustBuild(),
			},
			arg:     iId1,
			want:    i1,
			wantErr: nil,
		},
	}
	initDB := mongotest.Connect(t)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))
			r := NewIntegration(client)

			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntegrationRepo_FindByIDs(t *testing.T) {
	_, _, _, iId1, iId2, i1, i2 := testSuite()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     id.IntegrationIDList
		want    integration.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   integration.List{},
			arg:     id.IntegrationIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "0 count",
			seeds:   integration.List{i1, i2},
			arg:     id.IntegrationIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "1 count with single",
			seeds:   integration.List{i1, i2},
			arg:     id.IntegrationIDList{iId2},
			want:    integration.List{i2},
			wantErr: nil,
		},
		{
			name:    "2 count with multi",
			seeds:   integration.List{i1, i2},
			arg:     id.IntegrationIDList{iId1, iId2},
			want:    integration.List{i1, i2},
			wantErr: nil,
		},
		{
			name:    "2 count with multi (reverse order)",
			seeds:   integration.List{i1, i2},
			arg:     id.IntegrationIDList{iId2, iId1},
			want:    integration.List{i2, i1},
			wantErr: nil,
		},
	}
	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))
			r := NewIntegration(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}

			got, err := r.FindByIDs(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntegrationRepo_FindByUser(t *testing.T) {
	_, _, uId, _, _, i1, i2 := testSuite()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.UserID
		want    integration.List
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   integration.List{},
			arg:     uId,
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "Not found",
			seeds:   integration.List{i1, i2},
			arg:     id.NewUserID(),
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "Found",
			seeds:   integration.List{i1, i2},
			arg:     uId,
			want:    integration.List{i1, i2},
			wantErr: nil,
		},
	}
	initDB := mongotest.Connect(t)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))
			r := NewIntegration(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}

			got, err := r.FindByUser(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntegrationRepo_Save(t *testing.T) {
	_, _, _, _, _, i1, _ := testSuite()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     *integration.Integration
		want    integration.List
		wantErr error
		mockErr bool
	}{
		{
			name:    "Saved",
			seeds:   integration.List{},
			arg:     i1,
			want:    integration.List{i1},
			wantErr: nil,
		},
		{
			name:    "Saved same data",
			seeds:   integration.List{i1},
			arg:     i1,
			want:    integration.List{i1},
			wantErr: nil,
		},
	}
	initDB := mongotest.Connect(t)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))
			r := NewIntegration(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p.Clone())
				if tc.wantErr != nil {
					assert.ErrorIs(t, err, tc.wantErr)
					return
				}
			}

			err := r.Save(ctx, tc.arg.Clone())
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)

			got, err := r.FindByID(ctx, tc.arg.ID())
			assert.NoError(t, err)
			assert.Equal(t, tc.arg, got)
		})
	}
}

func TestIntegrationRepo_Remove(t *testing.T) {
	_, _, _, iId1, _, i1, _ := testSuite()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.ID
		want    integration.List
		wantErr error
	}{
		{
			name:    "Saved",
			seeds:   integration.List{},
			arg:     iId1,
			want:    integration.List{},
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "Saved same data",
			seeds:   integration.List{i1},
			arg:     iId1,
			want:    nil,
			wantErr: nil,
		},
	}
	initDB := mongotest.Connect(t)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))
			r := NewIntegration(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p.Clone())
				if tc.wantErr != nil {
					assert.ErrorIs(t, err, tc.wantErr)
					return
				}
			}

			err := r.Remove(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			_, err = r.FindByID(ctx, tc.arg)

			assert.Equal(t, rerror.ErrNotFound, err)
		})
	}
}
