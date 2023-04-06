package memory

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestModelRepo_Filtered(t *testing.T) {
	r := &Model{}
	pid := id.NewProjectID()

	assert.Equal(t, &Model{
		f: repo.ProjectFilter{
			Readable: id.ProjectIDList{pid},
			Writable: nil,
		},
		now: &util.TimeNow{},
	}, r.Filtered(repo.ProjectFilter{
		Readable: id.ProjectIDList{pid},
		Writable: nil,
	}))
}

func TestModelRepo_FindByID(t *testing.T) {
	mocknow := time.Now().Truncate(time.Millisecond).UTC()
	pid1 := id.NewProjectID()
	id1 := id.NewModelID()
	sid1 := id.NewSchemaID()
	k := key.New("T123456")
	m1 := model.New().ID(id1).Project(pid1).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild()

	tests := []struct {
		name    string
		seeds   model.List
		arg     id.ModelID
		filter  *repo.ProjectFilter
		want    *model.Model
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   model.List{},
			arg:     id.NewModelID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: model.List{
				model.New().NewID().Project(pid1).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.NewModelID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: model.List{
				m1,
			},
			arg:     id1,
			want:    m1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id1,
			want:    m1,
			wantErr: nil,
		},
		{
			name: "project filter operation success",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{pid1}, Writable: []id.ProjectID{pid1}},
			want:    m1,
			wantErr: nil,
		},
		{
			name: "project filter operation denied",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{}, Writable: []id.ProjectID{}},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewModel()
			defer MockModelNow(r, mocknow)()
			ctx := context.Background()

			for _, a := range tc.seeds {
				err := r.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestModelRepo_FindByIDs(t *testing.T) {
	mocknow := time.Now().Truncate(time.Millisecond).UTC()
	pid1 := id.NewProjectID()
	id1 := id.NewModelID()
	id2 := id.NewModelID()
	sid1 := id.NewSchemaID()
	sid2 := id.NewSchemaID()
	k := key.New("T123456")
	m1 := model.New().ID(id1).Project(pid1).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild()
	m2 := model.New().ID(id2).Project(pid1).Schema(sid2).Key(k).UpdatedAt(mocknow).MustBuild()

	tests := []struct {
		name    string
		seeds   model.List
		arg     id.ModelIDList
		filter  *repo.ProjectFilter
		want    model.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   model.List{},
			arg:     id.ModelIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with model for another workspaces",
			seeds: model.List{
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.ModelIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single model",
			seeds: model.List{
				m1,
			},
			arg:     id.ModelIDList{id1},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "1 count with multi models",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.ModelIDList{id1},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "2 count with multi models",
			seeds: model.List{
				m1,
				m2,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.ModelIDList{id1, id2},
			want:    model.List{m1, m2},
			wantErr: nil,
		},
		{
			name: "project filter operation success",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.ModelIDList{id1},
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{pid1}, Writable: []id.ProjectID{pid1}},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "project filter operation denied",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			arg:     id.ModelIDList{id1},
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{}, Writable: []id.ProjectID{}},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewModel()
			defer MockModelNow(r, mocknow)()
			ctx := context.Background()
			for _, a := range tc.seeds {
				err := r.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, err := r.FindByIDs(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestModelRepo_FindByProject(t *testing.T) {
	mocknow := time.Now().Truncate(time.Millisecond).UTC()
	pid1 := id.NewProjectID()
	id1 := id.NewModelID()
	id2 := id.NewModelID()
	sid1 := id.NewSchemaID()
	sid2 := id.NewSchemaID()
	k := key.New("T123456")
	m1 := model.New().ID(id1).Project(pid1).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild()
	m2 := model.New().ID(id2).Project(pid1).Schema(sid2).Key(k).UpdatedAt(mocknow).MustBuild()

	type args struct {
		tid   id.ProjectID
		pInfo *usecasex.Pagination
	}
	tests := []struct {
		name    string
		seeds   model.List
		args    args
		filter  *repo.ProjectFilter
		want    model.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   model.List{},
			args:    args{id.NewProjectID(), nil},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with model for another projects",
			seeds: model.List{
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			args:    args{id.NewProjectID(), nil},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single model",
			seeds: model.List{
				m1,
			},
			args:    args{pid1, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "1 count with multi models",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			args:    args{pid1, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "2 count with multi models",
			seeds: model.List{
				m1,
				m2,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			args:    args{pid1, usecasex.CursorPagination{First: lo.ToPtr(int64(2))}.Wrap()},
			want:    model.List{m1, m2},
			wantErr: nil,
		},
		{
			name: "project filter operation success",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			args:    args{pid1, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{pid1}, Writable: []id.ProjectID{pid1}},
			want:    model.List{m1},
			wantErr: nil,
		},
		{
			name: "project filter operation denied",
			seeds: model.List{
				m1,
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
				model.New().NewID().Project(id.NewProjectID()).Schema(sid1).Key(k).UpdatedAt(mocknow).MustBuild(),
			},
			args:    args{pid1, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  &repo.ProjectFilter{Readable: []id.ProjectID{}, Writable: []id.ProjectID{}},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := NewModel()
			defer MockModelNow(r, mocknow)()
			ctx := context.Background()
			for _, a := range tc.seeds {
				err := r.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, _, err := r.FindByProject(ctx, tc.args.tid, tc.args.pInfo)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
