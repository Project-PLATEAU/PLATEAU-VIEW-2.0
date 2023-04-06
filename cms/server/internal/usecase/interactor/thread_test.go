package interactor

import (
	"context"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestThread_FindByID(t *testing.T) {
	id1 := id.NewThreadID()
	wid1 := id.NewWorkspaceID()
	comments := []*thread.Comment{}
	th1 := thread.New().ID(id1).Workspace(wid1).Comments(comments).MustBuild()

	op := &usecase.Operator{}

	type args struct {
		id       id.ThreadID
		operator *usecase.Operator
	}

	tests := []struct {
		name    string
		seeds   []*thread.Thread
		args    args
		want    *thread.Thread
		wantErr error
	}{
		{
			name:  "Not found in empty db",
			seeds: []*thread.Thread{},
			args: args{
				id:       id.NewThreadID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: []*thread.Thread{
				thread.New().ID(id1).Workspace(wid1).Comments(comments).MustBuild(),
			},
			args: args{
				id:       id.NewThreadID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: []*thread.Thread{
				th1,
			},
			args: args{
				id:       id1,
				operator: op,
			},
			want:    th1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: []*thread.Thread{
				th1,
				thread.New().ID(id1).Workspace(wid1).Comments(comments).MustBuild(),
				thread.New().ID(id1).Workspace(wid1).Comments(comments).MustBuild(),
			},
			args: args{
				id:       id1,
				operator: op,
			},
			want:    th1,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, a := range tc.seeds {
				err := db.Thread.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}
			threadUC := NewThread(db, nil)

			got, err := threadUC.FindByID(ctx, tc.args.id, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestThread_FindByIDs(t *testing.T) {
	id1 := id.NewThreadID()
	wid1 := id.NewWorkspaceID()
	comments1 := []*thread.Comment{}
	th1 := thread.New().ID(id1).Workspace(wid1).Comments(comments1).MustBuild()

	id2 := id.NewThreadID()
	wid2 := id.NewWorkspaceID()
	comments2 := []*thread.Comment{}
	th2 := thread.New().ID(id2).Workspace(wid2).Comments(comments2).MustBuild()

	tests := []struct {
		name    string
		seeds   thread.List
		arg     id.ThreadIDList
		want    thread.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   thread.List{},
			arg:     []id.ThreadID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with thread for another workspaces",
			seeds: thread.List{
				thread.New().NewID().Workspace(id.NewWorkspaceID()).Comments([]*thread.Comment{}).MustBuild(),
			},
			arg:     []id.ThreadID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single thread",
			seeds: thread.List{
				th1,
			},
			arg:     []id.ThreadID{id1},
			want:    thread.List{th1},
			wantErr: nil,
		},
		{
			name: "1 count with multi threads",
			seeds: thread.List{
				th1,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).Comments([]*thread.Comment{}).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).Comments([]*thread.Comment{}).MustBuild(),
			},
			arg:     []id.ThreadID{id1},
			want:    thread.List{th1},
			wantErr: nil,
		},
		{
			name: "2 count with multi threads",
			seeds: thread.List{
				th1,
				th2,
				thread.New().NewID().Workspace(id.NewWorkspaceID()).Comments([]*thread.Comment{}).MustBuild(),
				thread.New().NewID().Workspace(id.NewWorkspaceID()).Comments([]*thread.Comment{}).MustBuild(),
			},
			arg:     []id.ThreadID{id1, id2},
			want:    thread.List{th1, th2},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, a := range tc.seeds {
				err := db.Thread.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}
			threadUC := NewThread(db, nil)

			got, err := threadUC.FindByIDs(ctx, tc.arg, &usecase.Operator{})
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestThreadRepo_CreateThread(t *testing.T) {
	wid := id.NewWorkspaceID()
	wid2 := id.WorkspaceID{}
	uid := id.NewUserID()
	op := &usecase.Operator{
		User:               &uid,
		ReadableWorkspaces: nil,
		WritableWorkspaces: nil,
		OwningWorkspaces:   []id.WorkspaceID{wid},
	}

	tests := []struct {
		name     string
		arg      id.WorkspaceID
		operator *usecase.Operator
		wantErr  error
	}{
		{
			name:     "Save succeed",
			arg:      wid,
			operator: op,
			wantErr:  nil,
		},
		{
			name: "Save error: invalid workspace id",
			arg:  wid2,
			operator: &usecase.Operator{
				User:               &uid,
				ReadableWorkspaces: nil,
				WritableWorkspaces: nil,
				OwningWorkspaces:   []id.WorkspaceID{wid2},
			},
			wantErr: thread.ErrNoWorkspaceID,
		},
		{
			name:     "operator error",
			arg:      wid,
			operator: &usecase.Operator{},
			wantErr:  interfaces.ErrOperationDenied,
		},
		{
			name:     "operator succeed",
			arg:      wid,
			operator: op,
			wantErr:  nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			db := memory.New()
			threadUC := NewThread(db, nil)

			th, err := threadUC.CreateThread(ctx, tc.arg, tc.operator)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}

			res, err := threadUC.FindByID(ctx, th.ID(), tc.operator)
			assert.NoError(t, err)
			assert.Equal(t, res, th)
		})
	}
}

func TestThread_AddComment(t *testing.T) {
	c1 := thread.NewComment(thread.NewCommentID(), operator.OperatorFromUser(id.NewUserID()), "aaa")
	wid := id.NewWorkspaceID()
	th1 := thread.New().NewID().Workspace(wid).Comments([]*thread.Comment{}).MustBuild()
	uid := id.NewUserID()
	op := &usecase.Operator{
		User:               &uid,
		ReadableWorkspaces: nil,
		WritableWorkspaces: nil,
		OwningWorkspaces:   []id.WorkspaceID{wid},
	}

	type args struct {
		content  string
		operator *usecase.Operator
	}

	tests := []struct {
		name      string
		seed      *thread.Thread
		args      args
		wantErr   error
		mockError bool
	}{
		{
			name: "workspaces invalid operator",
			seed: th1,
			args: args{
				content:  c1.Content(),
				operator: &usecase.Operator{},
			},
			wantErr: interfaces.ErrInvalidOperator,
		},
		{
			name: "workspaces operation success",
			seed: th1,
			args: args{
				content:  c1.Content(),
				operator: op,
			},
			wantErr: nil,
		},
		{
			name: "add comment success",
			seed: th1,
			args: args{
				content:  c1.Content(),
				operator: op,
			},
			wantErr: nil,
		},
		{
			name: "add comment fail",
			seed: th1,
			args: args{
				content:  c1.Content(),
				operator: op,
			},
			wantErr:   rerror.ErrNotFound,
			mockError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			thread := tc.seed.Clone()
			err := db.Thread.Save(ctx, thread)
			assert.NoError(t, err)

			threadUC := NewThread(db, nil)
			if tc.mockError && tc.wantErr != nil {
				thid := id.NewThreadID()
				_, _, err := threadUC.AddComment(ctx, thid, tc.args.content, tc.args.operator)
				assert.Equal(t, tc.wantErr, err)
				return
			}

			th, c, err := threadUC.AddComment(ctx, thread.ID(), tc.args.content, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, th)
				assert.NotNil(t, c)
			}

			th, err = threadUC.FindByID(ctx, thread.ID(), tc.args.operator)
			assert.NoError(t, err)
			assert.Equal(t, 1, len(th.Comments()))
			assert.True(t, th.HasComment(c.ID()))

		})
	}
}

func TestThread_UpdateComment(t *testing.T) {
	c1 := thread.NewComment(thread.NewCommentID(), operator.OperatorFromUser(id.NewUserID()), "aaa")
	c2 := thread.NewComment(thread.NewCommentID(), operator.OperatorFromUser(id.NewUserID()), "test")
	wid := id.NewWorkspaceID()
	th1 := thread.New().NewID().Workspace(wid).Comments([]*thread.Comment{c1, c2}).MustBuild()
	uid := id.NewUserID()
	op := &usecase.Operator{
		User:               &uid,
		ReadableWorkspaces: nil,
		WritableWorkspaces: nil,
		OwningWorkspaces:   []id.WorkspaceID{wid},
	}

	type args struct {
		comment  *thread.Comment
		content  string
		operator *usecase.Operator
	}

	tests := []struct {
		name      string
		seed      *thread.Thread
		args      args
		want      *thread.Comment
		wantErr   error
		mockError bool
	}{
		{
			name: "workspaces operation denied",
			seed: th1,
			args: args{
				comment:  c1,
				content:  "updated",
				operator: &usecase.Operator{},
			},
			wantErr: interfaces.ErrInvalidOperator,
		},
		{
			name: "workspaces operation success",
			args: args{
				comment:  c1,
				content:  "updated",
				operator: op,
			},
			seed:    th1,
			wantErr: nil,
		},
		{
			name: "update comment success",
			seed: th1,
			args: args{
				comment:  c1,
				content:  "updated",
				operator: op,
			},
			wantErr: nil,
		},
		{
			name: "update comment fail",
			seed: th1,
			args: args{
				comment:  c1,
				content:  "updated",
				operator: op,
			},
			wantErr:   rerror.ErrNotFound,
			mockError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			thread := tc.seed.Clone()
			err := db.Thread.Save(ctx, thread)
			assert.NoError(t, err)

			threadUC := NewThread(db, nil)
			if tc.mockError && tc.wantErr != nil {
				thid := id.NewThreadID()
				_, _, err := threadUC.UpdateComment(ctx, thid, tc.args.comment.ID(), tc.args.content, tc.args.operator)
				assert.Equal(t, tc.wantErr, err)
				return
			}
			if _, _, err := threadUC.UpdateComment(ctx, thread.ID(), tc.args.comment.ID(), tc.args.content, tc.args.operator); tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			} else {
				assert.NoError(t, err)
			}

			thread2, _ := threadUC.FindByID(ctx, thread.ID(), tc.args.operator)
			comment := thread2.Comments()[0]
			assert.Equal(t, tc.args.content, comment.Content())
		})
	}
}

func TestThread_DeleteComment(t *testing.T) {
	c1 := thread.NewComment(thread.NewCommentID(), operator.OperatorFromUser(id.NewUserID()), "aaa")
	c2 := thread.NewComment(thread.NewCommentID(), operator.OperatorFromUser(id.NewUserID()), "test")
	wid := id.NewWorkspaceID()
	th1 := thread.New().NewID().Workspace(wid).Comments([]*thread.Comment{c1, c2}).MustBuild()
	uid := id.NewUserID()
	op := &usecase.Operator{
		User:               &uid,
		ReadableWorkspaces: nil,
		WritableWorkspaces: nil,
		OwningWorkspaces:   []id.WorkspaceID{wid},
	}

	type args struct {
		commentId id.CommentID
		operator  *usecase.Operator
	}

	tests := []struct {
		name      string
		seed      *thread.Thread
		args      args
		want      *thread.Comment
		wantErr   error
		mockError bool
	}{
		{
			name:    "workspaces operation denied",
			seed:    th1,
			args:    args{commentId: c1.ID(), operator: &usecase.Operator{}},
			wantErr: interfaces.ErrInvalidOperator,
		},
		{
			name:    "workspaces operation success",
			seed:    th1,
			args:    args{commentId: c1.ID(), operator: op},
			wantErr: nil,
		},
		{
			name:    "delete comment success",
			seed:    th1,
			args:    args{commentId: c1.ID(), operator: op},
			wantErr: nil,
		},
		{
			name:      "delete comment fail",
			seed:      th1,
			args:      args{commentId: c1.ID(), operator: op},
			wantErr:   rerror.ErrNotFound,
			mockError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			thread1 := tc.seed.Clone()
			err := db.Thread.Save(ctx, thread1)
			assert.NoError(t, err)

			threadUC := NewThread(db, nil)
			if tc.mockError && tc.wantErr != nil {
				thid := id.NewThreadID()
				_, err := threadUC.DeleteComment(ctx, thid, tc.args.commentId, tc.args.operator)
				assert.Equal(t, tc.wantErr, err)
				return
			}

			if _, err := threadUC.DeleteComment(ctx, tc.seed.ID(), tc.args.commentId, tc.args.operator); tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			} else {
				assert.NoError(t, err)
			}

			commentID := tc.seed.Comments()[0].ID()
			thread2, err := threadUC.FindByID(ctx, tc.seed.ID(), tc.args.operator)
			assert.NoError(t, err)
			assert.Equal(t, len(tc.seed.Comments())-1, len(thread2.Comments()))
			assert.False(t, lo.ContainsBy(thread2.Comments(), func(cc *thread.Comment) bool { return cc.ID() == commentID }))
		})
	}
}
