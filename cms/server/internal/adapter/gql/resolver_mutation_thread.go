package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateThread(ctx context.Context, input gqlmodel.CreateThreadInput) (*gqlmodel.ThreadPayload, error) {
	wid, err := gqlmodel.ToID[id.Workspace](input.WorkspaceID)
	if err != nil {
		return nil, err
	}

	uc := usecases(ctx).Thread
	th, err := uc.CreateThread(ctx, wid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.ThreadPayload{Thread: gqlmodel.ToThread(th)}, nil
}

func (r *mutationResolver) AddComment(ctx context.Context, input gqlmodel.AddCommentInput) (*gqlmodel.CommentPayload, error) {
	thid := lo.Must(gqlmodel.ToID[id.Thread](input.ThreadID))

	uc := usecases(ctx).Thread
	th, c, err := uc.AddComment(ctx, thid, input.Content, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.CommentPayload{
		Thread:  gqlmodel.ToThread(th),
		Comment: gqlmodel.ToComment(c, th),
	}, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, input gqlmodel.UpdateCommentInput) (*gqlmodel.CommentPayload, error) {
	thid, err := gqlmodel.ToID[id.Thread](input.ThreadID)
	if err != nil {
		return nil, err
	}

	cid, err := gqlmodel.ToID[id.Comment](input.CommentID)
	if err != nil {
		return nil, err
	}

	uc := usecases(ctx).Thread
	th, c, err := uc.UpdateComment(ctx, thid, cid, input.Content, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.CommentPayload{
		Thread:  gqlmodel.ToThread(th),
		Comment: gqlmodel.ToComment(c, th),
	}, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, input gqlmodel.DeleteCommentInput) (*gqlmodel.DeleteCommentPayload, error) {
	thid, err := gqlmodel.ToID[id.Thread](input.ThreadID)
	if err != nil {
		return nil, err
	}

	cid, err := gqlmodel.ToID[id.Comment](input.CommentID)
	if err != nil {
		return nil, err
	}

	uc := usecases(ctx).Thread
	th, err := uc.DeleteComment(ctx, thid, cid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteCommentPayload{
		Thread:    gqlmodel.ToThread(th),
		CommentID: gqlmodel.IDFrom(cid),
	}, nil
}
