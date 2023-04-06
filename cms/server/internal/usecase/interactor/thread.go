package interactor

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
)

type Thread struct {
	repos    *repo.Container
	gateways *gateway.Container
}

func NewThread(r *repo.Container, g *gateway.Container) interfaces.Thread {
	return &Thread{
		repos:    r,
		gateways: g,
	}
}

func (i *Thread) FindByID(ctx context.Context, aid id.ThreadID, op *usecase.Operator) (*thread.Thread, error) {
	return Run1(
		ctx, op, i.repos,
		Usecase().Transaction(),
		func(ctx context.Context) (*thread.Thread, error) {
			return i.repos.Thread.FindByID(ctx, aid)
		},
	)
}

func (i *Thread) FindByIDs(ctx context.Context, threads []id.ThreadID, operator *usecase.Operator) (thread.List, error) {
	return i.repos.Thread.FindByIDs(ctx, threads)
}

func (i *Thread) CreateThread(ctx context.Context, wid id.WorkspaceID, op *usecase.Operator) (*thread.Thread, error) {
	return Run1(
		ctx, op, i.repos,
		Usecase().WithWritableWorkspaces(wid).Transaction(),
		func(ctx context.Context) (*thread.Thread, error) {
			thread, err := thread.New().NewID().Workspace(wid).Build()
			if err != nil {
				return nil, err
			}

			if err := i.repos.Thread.Save(ctx, thread); err != nil {
				return nil, err
			}

			return thread, nil
		},
	)
}

func (i *Thread) AddComment(ctx context.Context, thid id.ThreadID, content string, op *usecase.Operator) (*thread.Thread, *thread.Comment, error) {
	if op.User == nil && op.Integration == nil {
		return nil, nil, interfaces.ErrInvalidOperator
	}
	return Run2(
		ctx, op, i.repos,
		Usecase().Transaction(),
		func(ctx context.Context) (*thread.Thread, *thread.Comment, error) {
			th, err := i.repos.Thread.FindByID(ctx, thid)
			if err != nil {
				return nil, nil, err
			}

			if !op.IsWritableWorkspace(th.Workspace()) {
				return nil, nil, interfaces.ErrOperationDenied
			}

			comment := thread.NewComment(thread.NewCommentID(), op.Operator(), content)
			if err := th.AddComment(comment); err != nil {
				return nil, nil, err
			}

			if err := i.repos.Thread.Save(ctx, th); err != nil {
				return nil, nil, err
			}

			return th, comment, nil
		},
	)
}

func (i *Thread) UpdateComment(ctx context.Context, thid id.ThreadID, cid id.CommentID, content string, op *usecase.Operator) (*thread.Thread, *thread.Comment, error) {
	if op.User == nil && op.Integration == nil {
		return nil, nil, interfaces.ErrInvalidOperator
	}
	return Run2(
		ctx, op, i.repos,
		Usecase().Transaction(),
		func(ctx context.Context) (*thread.Thread, *thread.Comment, error) {
			th, err := i.repos.Thread.FindByID(ctx, thid)
			if err != nil {
				return nil, nil, err
			}

			if !op.IsWritableWorkspace(th.Workspace()) {
				return nil, nil, interfaces.ErrOperationDenied
			}

			if err := th.UpdateComment(cid, content); err != nil {
				return nil, nil, err
			}

			if err := i.repos.Thread.Save(ctx, th); err != nil {
				return nil, nil, err
			}

			return th, th.Comment(cid), nil
		},
	)
}

func (i *Thread) DeleteComment(ctx context.Context, thid id.ThreadID, cid id.CommentID, op *usecase.Operator) (*thread.Thread, error) {
	if op.User == nil && op.Integration == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(
		ctx, op, i.repos,
		Usecase().Transaction(),
		func(ctx context.Context) (*thread.Thread, error) {
			th, err := i.repos.Thread.FindByID(ctx, thid)
			if err != nil {
				return nil, err
			}

			if !op.IsWritableWorkspace(th.Workspace()) {
				return nil, interfaces.ErrOperationDenied
			}

			if err := th.DeleteComment(cid); err != nil {
				return nil, err
			}

			if err := i.repos.Thread.Save(ctx, th); err != nil {
				return nil, err
			}

			return th, nil
		},
	)
}
