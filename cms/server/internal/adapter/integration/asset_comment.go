package integration

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
)

func (s Server) AssetCommentList(ctx context.Context, request AssetCommentListRequestObject) (AssetCommentListResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	aID := id.AssetID(request.AssetId)

	asset, err := uc.Asset.FindByID(ctx, aID, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetCommentList404Response{}, err
		}
		return AssetCommentList400Response{}, err
	}

	threadID := asset.Thread()
	th, err := uc.Thread.FindByID(ctx, threadID, op)
	if err != nil {
		return nil, err
	}

	comments := lo.Map(th.Comments(), func(c *thread.Comment, _ int) integrationapi.Comment {
		return *integrationapi.NewComment(c)
	})

	return AssetCommentList200JSONResponse{Comments: &comments}, nil
}

func (s Server) AssetCommentCreate(ctx context.Context, request AssetCommentCreateRequestObject) (AssetCommentCreateResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	asset, err := uc.Asset.FindByID(ctx, request.AssetId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetCommentCreate404Response{}, err
		}
		return AssetCommentCreate400Response{}, err
	}

	threadID := asset.Thread()
	_, comment, err := uc.Thread.AddComment(ctx, threadID, *request.Body.Content, op)
	if err != nil {
		return nil, err
	}

	return AssetCommentCreate200JSONResponse(*integrationapi.NewComment(comment)), nil
}

func (s Server) AssetCommentUpdate(ctx context.Context, request AssetCommentUpdateRequestObject) (AssetCommentUpdateResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	asset, err := uc.Asset.FindByID(ctx, request.AssetId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetCommentUpdate404Response{}, err
		}
		return AssetCommentUpdate400Response{}, err
	}

	_, comment, err := uc.Thread.UpdateComment(ctx, asset.Thread(), request.CommentId, *request.Body.Content, op)
	if err != nil {
		return nil, err
	}

	return AssetCommentUpdate200JSONResponse(*integrationapi.NewComment(comment)), nil
}

func (s Server) AssetCommentDelete(ctx context.Context, request AssetCommentDeleteRequestObject) (AssetCommentDeleteResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	asset, err := uc.Asset.FindByID(ctx, request.AssetId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetCommentDelete404Response{}, err
		}
		return AssetCommentDelete400Response{}, err
	}

	threadID := asset.Thread()
	_, err = uc.Thread.DeleteComment(ctx, threadID, request.CommentId, op)
	if err != nil {
		return nil, err
	}

	return AssetCommentDelete200JSONResponse{Id: request.CommentId.Ref()}, nil
}
