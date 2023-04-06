package memory

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/stretchr/testify/assert"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/rerror"
)

func TestIntegrationRepo_FindByID(t *testing.T) {
	now := time.Now()
	iId1 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.ID
		want    *integration.Integration
		wantErr error
		mockErr bool
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
				integration.New().NewID().UpdatedAt(now).MustBuild(),
				i1,
				integration.New().NewID().UpdatedAt(now).MustBuild(),
			},
			arg:     iId1,
			want:    i1,
			wantErr: nil,
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

			r := NewIntegration()
			if tc.mockErr {
				SetIntegrationError(r, tc.wantErr)
			}
			defer MockIntegrationNow(r, now)()
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
	now := time.Now()
	iId1 := id.NewIntegrationID()
	iId2 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).UpdatedAt(now).MustBuild()
	i2 := integration.New().ID(iId2).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     id.IntegrationIDList
		want    integration.List
		wantErr error
		mockErr bool
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
			want:    integration.List{i1, i2},
			wantErr: nil,
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

			r := NewIntegration()
			if tc.mockErr {
				SetIntegrationError(r, tc.wantErr)
			}
			defer MockIntegrationNow(r, now)()
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
	now := time.Now()
	uId := id.NewUserID()
	iId1 := id.NewIntegrationID()
	iId2 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).Developer(uId).UpdatedAt(now).MustBuild()
	i2 := integration.New().ID(iId2).Developer(uId).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.UserID
		want    integration.List
		wantErr error
		mockErr bool
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

			r := NewIntegration()
			if tc.mockErr {
				SetIntegrationError(r, tc.wantErr)
			}
			defer MockIntegrationNow(r, now)()
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

func TestIntegrationRepo_Remove(t *testing.T) {
	now := time.Now()
	iId1 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).UpdatedAt(now).MustBuild()

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

			r := NewIntegration()
			if tc.mockErr {
				SetIntegrationError(r, tc.wantErr)
			}
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

			assert.Equal(t, []*integration.Integration(tc.want), r.(*Integration).data.Values())
		})
	}
}

func TestIntegrationRepo_Save(t *testing.T) {
	now := time.Now()
	iId1 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   integration.List
		arg     integration.ID
		want    integration.List
		wantErr error
		mockErr bool
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

			r := NewIntegration()
			if tc.mockErr {
				SetIntegrationError(r, tc.wantErr)
			}
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

			assert.Equal(t, []*integration.Integration(tc.want), r.(*Integration).data.Values())
		})
	}
}
