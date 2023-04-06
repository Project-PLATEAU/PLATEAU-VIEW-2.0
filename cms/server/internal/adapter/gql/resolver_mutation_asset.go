package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

func (r *mutationResolver) CreateAsset(ctx context.Context, input gqlmodel.CreateAssetInput) (*gqlmodel.CreateAssetPayload, error) {
	pid, err := gqlmodel.ToID[id.Project](input.ProjectID)
	if err != nil {
		return nil, err
	}
	uc := usecases(ctx).Asset
	params := interfaces.CreateAssetParam{
		ProjectID: pid,
		File:      gqlmodel.FromFile(input.File),
	}
	if input.URL != nil {
		params.URL = *input.URL
	}
	if input.SkipDecompression != nil {
		params.SkipDecompression = *input.SkipDecompression
	}

	res, _, err := uc.Create(ctx, params, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.CreateAssetPayload{
		Asset: gqlmodel.ToAsset(res, uc.GetURL),
	}, nil
}

func (r *mutationResolver) UpdateAsset(ctx context.Context, input gqlmodel.UpdateAssetInput) (*gqlmodel.UpdateAssetPayload, error) {
	aid, err := gqlmodel.ToID[id.Asset](input.ID)
	if err != nil {
		return nil, err
	}

	uc := usecases(ctx).Asset
	res, err2 := uc.Update(ctx, interfaces.UpdateAssetParam{
		AssetID:     aid,
		PreviewType: gqlmodel.FromPreviewType(input.PreviewType),
	}, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &gqlmodel.UpdateAssetPayload{
		Asset: gqlmodel.ToAsset(res, uc.GetURL),
	}, nil
}

func (r *mutationResolver) DeleteAsset(ctx context.Context, input gqlmodel.DeleteAssetInput) (*gqlmodel.DeleteAssetPayload, error) {
	aid, err := gqlmodel.ToID[id.Asset](input.AssetID)
	if err != nil {
		return nil, err
	}

	res, err2 := usecases(ctx).Asset.Delete(ctx, aid, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &gqlmodel.DeleteAssetPayload{AssetID: gqlmodel.IDFrom(res)}, nil
}

func (r *mutationResolver) DecompressAsset(ctx context.Context, input gqlmodel.DecompressAssetInput) (*gqlmodel.DecompressAssetPayload, error) {
	aid, err := gqlmodel.ToID[id.Asset](input.AssetID)
	if err != nil {
		return nil, err
	}

	uc := usecases(ctx).Asset
	res, err2 := uc.DecompressByID(ctx, aid, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &gqlmodel.DecompressAssetPayload{Asset: gqlmodel.ToAsset(res, uc.GetURL)}, nil
}
