package integration

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
)

func (s Server) ItemCommentList(ctx context.Context, request ItemCommentListRequestObject) (ItemCommentListResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	i, err := uc.Item.FindByID(ctx, request.ItemId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return ItemCommentList404Response{}, err
		}
		return ItemCommentList400Response{}, err
	}

	thId := i.Value().Thread()
	th, err := uc.Thread.FindByID(ctx, thId, op)
	if err != nil {
		return ItemCommentList400Response{}, err
	}

	comments := lo.Map(th.Comments(), func(c *thread.Comment, _ int) integrationapi.Comment {
		return *integrationapi.NewComment(c)
	})

	return ItemCommentList200JSONResponse{Comments: &comments}, nil
}

func (s Server) ItemCommentCreate(ctx context.Context, request ItemCommentCreateRequestObject) (ItemCommentCreateResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	i, err := uc.Item.FindByID(ctx, request.ItemId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return ItemCommentCreate404Response{}, err
		}
		return ItemCommentCreate400Response{}, err
	}

	thId := i.Value().Thread()
	_, comment, err := uc.Thread.AddComment(ctx, thId, *request.Body.Content, op)
	if err != nil {
		return ItemCommentCreate400Response{}, err
	}

	return ItemCommentCreate200JSONResponse(*integrationapi.NewComment(comment)), nil
}

func (s Server) ItemCommentUpdate(ctx context.Context, request ItemCommentUpdateRequestObject) (ItemCommentUpdateResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	i, err := uc.Item.FindByID(ctx, request.ItemId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return ItemCommentUpdate404Response{}, err
		}
		return ItemCommentUpdate400Response{}, err
	}

	_, comment, err := uc.Thread.UpdateComment(ctx, i.Value().Thread(), request.CommentId, *request.Body.Content, op)
	if err != nil {
		return ItemCommentUpdate400Response{}, err
	}

	return ItemCommentUpdate200JSONResponse(*integrationapi.NewComment(comment)), nil
}

func (s Server) ItemCommentDelete(ctx context.Context, request ItemCommentDeleteRequestObject) (ItemCommentDeleteResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	i, err := uc.Item.FindByID(ctx, request.ItemId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return ItemCommentDelete404Response{}, err
		}
		return ItemCommentDelete400Response{}, err
	}

	thId := i.Value().Thread()
	_, err = uc.Thread.DeleteComment(ctx, thId, request.CommentId, op)
	if err != nil {
		return ItemCommentDelete400Response{}, err
	}

	return ItemCommentDelete200JSONResponse{Id: &request.CommentId}, nil
}
