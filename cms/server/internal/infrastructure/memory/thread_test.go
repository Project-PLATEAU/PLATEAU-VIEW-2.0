package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestThreadRepo_Save(t *testing.T) {
	wid1 := id.NewWorkspaceID()
	id1 := id.NewThreadID()
	th1 := thread.New().ID(id1).Workspace(wid1).MustBuild()

	tests := []struct {
		name    string
		seeds   thread.List
		arg     *thread.Thread
		filter  *repo.WorkspaceFilter
		want    *thread.Thread
		wantErr error
	}{
		{
			name: "Save succeed",
			seeds: thread.List{
				th1,
			},
			arg:     th1,
			want:    th1,
			wantErr: nil,
		},
		{
			name: "Filtered operation error",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     th1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{}, Writable: []id.WorkspaceID{}},
			want:    nil,
			wantErr: repo.ErrOperationDenied,
		},
		{
			name: "Filtered succeed",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     th1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{wid1}, Writable: []id.WorkspaceID{wid1}},
			want:    th1,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := NewThread()
			ctx := context.Background()

			if tc.filter != nil {
				r = r.Filtered(*tc.filter)
			}

			for _, th := range tc.seeds {
				err := r.Save(ctx, th)
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
		})
	}
}

func TestThread_Filtered(t *testing.T) {
	r := &Thread{}
	wid := id.NewWorkspaceID()

	assert.Equal(t, &Thread{
		f: repo.WorkspaceFilter{
			Readable: id.WorkspaceIDList{wid},
			Writable: nil,
		},
	}, r.Filtered(repo.WorkspaceFilter{
		Readable: id.WorkspaceIDList{wid},
		Writable: nil,
	}))
}

func TestThreadRepo_FindByID(t *testing.T) {
	tid1 := id.NewWorkspaceID()
	id1 := id.NewThreadID()
	th1 := thread.New().ID(id1).Workspace(tid1).MustBuild()
	tests := []struct {
		name    string
		seeds   thread.List
		arg     id.ThreadID
		filter  *repo.WorkspaceFilter
		want    *thread.Thread
		wantErr error
		mockErr bool
	}{
		{
			name:    "Not found in empty db",
			seeds:   thread.List{},
			arg:     id.NewThreadID(),
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: thread.List{
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id.NewThreadID(),
			filter:  nil,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: thread.List{
				th1,
			},
			arg:     id1,
			filter:  nil,
			want:    th1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  nil,
			want:    th1,
			wantErr: nil,
		},
		{
			name: "Filtered Found 0",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{id.NewWorkspaceID()}, Writable: []id.WorkspaceID{}},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Filtered Found 2",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id1,
			filter:  &repo.WorkspaceFilter{Readable: []id.WorkspaceID{tid1}, Writable: []id.WorkspaceID{}},
			want:    th1,
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

			r := NewThread()
			if tc.mockErr {
				SetThreadError(r, tc.wantErr)
			}

			ctx := context.Background()
			for _, th := range tc.seeds {
				err := r.Save(ctx, th.Clone())
				assert.Nil(t, err)
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

func TestThreadRepo_FindByIDs(t *testing.T) {
	wid1 := id.NewWorkspaceID()
	id1 := id.NewThreadID()
	id2 := id.NewThreadID()
	th1 := thread.New().ID(id1).Workspace(wid1).MustBuild()
	th2 := thread.New().ID(id2).Workspace(wid1).MustBuild()

	tests := []struct {
		name    string
		seeds   []*thread.Thread
		arg     id.ThreadIDList
		want    []*thread.Thread
		wantErr error
		mockErr bool
	}{
		{
			name:    "0 count in empty db",
			seeds:   []*thread.Thread{},
			arg:     id.ThreadIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with thread for another workspaces",
			seeds: []*thread.Thread{
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id.ThreadIDList{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single thread",
			seeds: []*thread.Thread{
				th1,
			},
			arg:     id.ThreadIDList{id1},
			want:    []*thread.Thread{th1},
			wantErr: nil,
		},
		{
			name: "1 count with multi threads",
			seeds: []*thread.Thread{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id.ThreadIDList{id1},
			want:    []*thread.Thread{th1},
			wantErr: nil,
		},
		{
			name: "2 count with multi threads",
			seeds: []*thread.Thread{
				th1,
				th2,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).MustBuild(),
			},
			arg:     id.ThreadIDList{id1, id2},
			want:    []*thread.Thread{th1, th2},
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

			r := NewThread()
			if tc.mockErr {
				SetThreadError(r, tc.wantErr)
			}

			ctx := context.Background()
			for _, a := range tc.seeds {
				err := r.Save(ctx, a.Clone())
				assert.Nil(t, err)
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
