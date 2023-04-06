package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_projectRepo_CountByWorkspace(t *testing.T) {
	tid1 := id.NewWorkspaceID()
	tests := []struct {
		name    string
		seeds   project.List
		arg     id.WorkspaceID
		filter  *repo.WorkspaceFilter
		want    int
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   project.List{},
			arg:     id.NewWorkspaceID(),
			filter:  nil,
			want:    0,
			wantErr: nil,
		},
		{
			name: "0 count with project for another workspaces",
			seeds: project.List{
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id.NewWorkspaceID(),
			filter:  nil,
			want:    0,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: project.List{
				project.New().NewID().Workspace(tid1).MustBuild(),
			},
			arg:     tid1,
			filter:  nil,
			want:    1,
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: project.List{
				project.New().NewID().Workspace(tid1).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     tid1,
			filter:  nil,
			want:    1,
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: project.List{
				project.New().NewID().Workspace(tid1).MustBuild(),
				project.New().NewID().Workspace(tid1).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     tid1,
			filter:  nil,
			want:    2,
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: project.List{
				project.New().NewID().Workspace(tid1).MustBuild(),
				project.New().NewID().Workspace(tid1).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     tid1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{}},
			want:    2,
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, err := r.CountByWorkspace(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_Filtered(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	id2 := id.NewProjectID()
	p1 := project.New().ID(id1).Workspace(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().ID(id2).Workspace(tid1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   project.List
		arg     repo.WorkspaceFilter
		wantErr error
	}{
		{
			name: "no r/w workspaces operation denied",
			seeds: project.List{
				p1,
				p2,
			},
			arg: repo.WorkspaceFilter{
				Readable: []id.WorkspaceID{},
				Writable: []id.WorkspaceID{},
			},
			wantErr: repo.ErrOperationDenied,
		},
		{
			name: "r/w workspaces operation success",
			seeds: project.List{
				p1,
				p2,
			},
			arg: repo.WorkspaceFilter{
				Readable: []id.WorkspaceID{tid1},
				Writable: []id.WorkspaceID{tid1},
			},
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client).Filtered(tc.arg)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.ErrorIs(t, err, tc.wantErr)
			}
		})
	}
}

func Test_projectRepo_FindByID(t *testing.T) {
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	now := time.Now().Truncate(time.Millisecond).UTC()
	p1 := project.New().ID(id1).Workspace(tid1).UpdatedAt(now).MustBuild()
	tests := []struct {
		name    string
		seeds   project.List
		arg     id.ProjectID
		filter  *repo.WorkspaceFilter
		want    *project.Project
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   project.List{},
			arg:     id.NewProjectID(),
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: project.List{
				project.New().NewID().MustBuild(),
			},
			arg:     id.NewProjectID(),
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: project.List{
				p1,
			},
			arg:     id1,
			filter:  nil,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  nil,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Filtered Found 0",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{}},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Filtered Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{tid1}, Writable: []id.WorkspaceID{}},
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
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

func Test_projectRepo_FindByIDs(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	id2 := id.NewProjectID()
	p1 := project.New().ID(id1).Workspace(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().ID(id2).Workspace(tid1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   project.List
		arg     id.ProjectIDList
		filter  *repo.WorkspaceFilter
		want    project.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   project.List{},
			arg:     []id.ProjectID{},
			filter:  nil,
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with project for another workspaces",
			seeds: project.List{
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     []id.ProjectID{},
			filter:  nil,
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: project.List{
				p1,
			},
			arg:     []id.ProjectID{id1},
			filter:  nil,
			want:    project.List{p1},
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1},
			filter:  nil,
			want:    project.List{p1},
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1, id2},
			filter:  nil,
			want:    project.List{p1, p2},
			wantErr: nil,
		},
		{
			name: "Filter 2 count with multi projects",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1, id2},
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{}},
			want:    project.List{ /*nil, nil*/ },
			wantErr: nil,
		},
		{
			name: "Filter 2 count with multi projects",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1, id2},
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{tid1}, Writable: []id.WorkspaceID{}},
			want:    project.List{p1, p2},
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
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

func Test_projectRepo_FindByPublicName(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	p1 := project.New().
		ID(id1).
		Workspace(tid1).
		Alias("xyz123").
		UpdatedAt(now).
		MustBuild()

	id2 := id.NewProjectID()
	p2 := project.New().
		ID(id2).
		Workspace(id.NewWorkspaceID()).
		Alias("xyz321").
		UpdatedAt(now).
		MustBuild()

	tests := []struct {
		name    string
		seeds   project.List
		arg     string
		filter  *repo.WorkspaceFilter
		want    *project.Project
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   project.List{},
			arg:     "xyz123",
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: project.List{
				project.New().NewID().Alias("abc123").MustBuild(),
			},
			arg:     "xyz123",
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "public Found",
			seeds: project.List{
				p1,
			},
			arg:     "xyz123",
			filter:  nil,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "linited Found",
			seeds: project.List{
				p2,
			},
			arg:     "xyz321",
			want:    p2,
			filter:  nil,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     "xyz123",
			filter:  nil,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Filtered should not Found",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     "xyz123",
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{}},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Filtered should Found",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     "xyz123",
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{tid1}, Writable: []id.WorkspaceID{}},
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, err := r.FindByPublicName(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_FindByWorkspace(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewWorkspaceID()
	p1 := project.New().NewID().Workspace(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().NewID().Workspace(tid1).UpdatedAt(now).MustBuild()

	type args struct {
		wids  id.WorkspaceIDList
		pInfo *usecasex.Pagination
	}
	tests := []struct {
		name    string
		seeds   project.List
		args    args
		filter  *repo.WorkspaceFilter
		want    project.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   project.List{},
			args:    args{id.WorkspaceIDList{id.NewWorkspaceID()}, nil},
			filter:  nil,
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with project for another workspaces",
			seeds: project.List{
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{id.NewWorkspaceID()}, nil},
			filter:  nil,
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: project.List{
				p1,
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  nil,
			want:    project.List{p1},
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  nil,
			want:    project.List{p1},
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{First: lo.ToPtr(int64(2))}.Wrap()},
			filter:  nil,
			want:    project.List{p1, p2},
			wantErr: nil,
		},
		{
			name: "get 1st page of 2",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  nil,
			want:    project.List{p1},
			wantErr: nil,
		},
		{
			name: "get last page of 2",
			seeds: project.List{
				p1,
				p2,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{Last: lo.ToPtr(int64(1))}.Wrap()},
			filter:  nil,
			want:    project.List{p2},
			wantErr: nil,
		},
		{
			name: "Filtered sholud not 1 count with multi projects",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			args:    args{id.WorkspaceIDList{tid1}, usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap()},
			filter:  &repo.WorkspaceFilter{Readable: id.WorkspaceIDList{id.NewWorkspaceID()}, Writable: id.WorkspaceIDList{}},
			want:    project.List{p1},
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			got, _, err := r.FindByWorkspaces(ctx, tc.args.wids, tc.args.pInfo)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_Remove(t *testing.T) {
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	p1 := project.New().ID(id1).Workspace(tid1).MustBuild()
	tests := []struct {
		name    string
		seeds   project.List
		arg     id.ProjectID
		filter  *repo.WorkspaceFilter
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   project.List{},
			arg:     id.NewProjectID(),
			filter:  nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: project.List{
				project.New().NewID().MustBuild(),
			},
			arg:     id.NewProjectID(),
			filter:  nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: project.List{
				p1,
			},
			arg:     id1,
			filter:  nil,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  nil,
			wantErr: nil,
		},
		{
			name: "Filtered should fail Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{id.NewWorkspaceID()}},
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Filtered should work Found 2",
			seeds: project.List{
				p1,
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				project.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{}, Writable: []id.WorkspaceID{tid1}},
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			err := r.Remove(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			_, err = r.FindByID(ctx, tc.arg)
			assert.ErrorIs(t, err, rerror.ErrNotFound)
		})
	}
}

func Test_projectRepo_Save(t *testing.T) {
	tid1 := id.NewWorkspaceID()
	id1 := id.NewProjectID()
	p1 := project.New().ID(id1).Workspace(tid1).UpdatedAt(time.Now().Truncate(time.Millisecond).UTC()).MustBuild()
	tests := []struct {
		name    string
		seeds   project.List
		arg     *project.Project
		filter  *repo.WorkspaceFilter
		want    *project.Project
		wantErr error
	}{
		{
			name: "Saved",
			seeds: project.List{
				p1,
			},
			arg:     p1,
			filter:  nil,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Filtered should fail - Saved",
			seeds: project.List{
				p1,
			},
			arg:     p1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{}, Writable: []id.WorkspaceID{}},
			want:    nil,
			wantErr: repo.ErrOperationDenied,
		},
		{
			name: "Filtered should work - Saved",
			seeds: project.List{
				p1,
			},
			arg:     p1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{tid1}, Writable: []id.WorkspaceID{tid1}},
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Filtered should work - Saved same data",
			seeds: project.List{
				p1,
			},
			arg:     p1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{}, Writable: []id.WorkspaceID{tid1}},
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewProject(client)
			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				if tc.wantErr != nil {
					assert.ErrorIs(t, err, tc.wantErr)
					return
				}
			}

			err := r.Save(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			got, err := r.CountByWorkspace(ctx, tc.arg.Workspace())
			assert.NoError(t, err)
			if tc.wantErr != nil {
				assert.Zero(t, got)
				return
			}
			assert.Equal(t, 1, got)
		})
	}
}
