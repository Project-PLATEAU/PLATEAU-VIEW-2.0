package integration

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

func (s Server) AssetFilter(ctx context.Context, request AssetFilterRequestObject) (AssetFilterResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	var sort *usecasex.Sort
	if request.Params.Sort != nil {
		sort = &usecasex.Sort{
			Key:      string(*request.Params.Sort),
			Reverted: request.Params.Dir != nil && *request.Params.Dir == integrationapi.AssetFilterParamsDirDesc,
		}
	}

	f := interfaces.AssetFilter{
		Keyword:    nil,
		Sort:       sort,
		Pagination: fromPagination(request.Params.Page, request.Params.PerPage),
	}

	assets, pi, err := uc.Asset.FindByProject(ctx, request.ProjectId, f, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetFilter404Response{}, err
		}
		return AssetFilter400Response{}, err
	}

	itemList, err := util.TryMap(assets, func(a *asset.Asset) (integrationapi.Asset, error) {
		aurl := uc.Asset.GetURL(a)
		aa := integrationapi.NewAsset(a, nil, aurl, true)
		return *aa, nil
	})
	if err != nil {
		return AssetFilter400Response{}, err
	}

	return AssetFilter200JSONResponse{
		Items:      &itemList,
		Page:       request.Params.Page,
		PerPage:    request.Params.PerPage,
		TotalCount: lo.ToPtr(int(pi.TotalCount)),
	}, nil
}

func (s Server) AssetCreate(ctx context.Context, request AssetCreateRequestObject) (AssetCreateResponseObject, error) {
	uc := adapter.Usecases(ctx)
	op := adapter.Operator(ctx)

	var f *file.File
	var err error
	if request.MultipartBody != nil {
		f, err = file.FromMultipart(request.MultipartBody, "file")
		if err != nil {
			return AssetCreate400Response{}, err
		}
	}

	var url string
	skipDecompression := false
	if request.JSONBody != nil {
		url = *request.JSONBody.Url
		if request.JSONBody.SkipDecompression != nil {
			skipDecompression = *request.JSONBody.SkipDecompression
		}
	}

	cp := interfaces.CreateAssetParam{
		ProjectID:         request.ProjectId,
		File:              f,
		URL:               url,
		SkipDecompression: skipDecompression,
	}

	a, af, err := uc.Asset.Create(ctx, cp, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetCreate404Response{}, err
		}
		return AssetCreate400Response{}, err
	}

	aurl := uc.Asset.GetURL(a)
	aa := integrationapi.NewAsset(a, af, aurl, true)
	return AssetCreate200JSONResponse(*aa), nil
}

func (s Server) AssetDelete(ctx context.Context, request AssetDeleteRequestObject) (AssetDeleteResponseObject, error) {
	uc := adapter.Usecases(ctx)
	op := adapter.Operator(ctx)
	aId, err := uc.Asset.Delete(ctx, request.AssetId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetDelete404Response{}, err
		}
		return AssetDelete400Response{}, err
	}

	return AssetDelete200JSONResponse{
		Id: &aId,
	}, nil
}

func (s Server) AssetGet(ctx context.Context, request AssetGetRequestObject) (AssetGetResponseObject, error) {
	uc := adapter.Usecases(ctx)
	op := adapter.Operator(ctx)

	a, err := uc.Asset.FindByID(ctx, request.AssetId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return AssetGet404Response{}, err
		}
		return AssetGet400Response{}, err
	}

	f, err := uc.Asset.FindFileByID(ctx, request.AssetId, op)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return AssetGet400Response{}, err
	}

	aurl := uc.Asset.GetURL(a)
	aa := integrationapi.NewAsset(a, f, aurl, true)
	return AssetGet200JSONResponse(*aa), nil
}
