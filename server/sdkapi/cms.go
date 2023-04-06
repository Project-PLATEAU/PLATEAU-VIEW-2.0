package sdkapi

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type CMS struct {
	Project              string
	PublicAPI            bool
	IntegrationAPIClient cms.Interface
	PublicAPIClient      *cms.PublicAPIClient[Item]
}

func NewCMS(c cms.Interface, pac *cms.PublicAPIClient[Item], project string, usePublic bool) *CMS {
	return &CMS{
		Project:              project,
		PublicAPI:            usePublic,
		IntegrationAPIClient: c,
		PublicAPIClient:      pac,
	}
}

func (c *CMS) Datasets(ctx context.Context, model string) (*DatasetResponse, error) {
	if c.PublicAPI {
		return c.DatasetsWithPublicAPI(ctx, model)
	}
	return c.DatasetsWithIntegrationAPI(ctx, model)
}

func (c *CMS) Files(ctx context.Context, model, id string) (any, error) {
	if c.PublicAPI {
		return c.FilesWithPublicAPI(ctx, model, id)
	}
	return c.FilesWithIntegrationAPI(ctx, model, id)
}

func (c *CMS) DatasetsWithPublicAPI(ctx context.Context, model string) (*DatasetResponse, error) {
	items, err := c.PublicAPIClient.GetAllItems(ctx, model)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	return Items(items).DatasetResponse(), nil
}

func (c *CMS) FilesWithPublicAPI(ctx context.Context, model, id string) (any, error) {
	item, err := c.PublicAPIClient.GetItem(ctx, model, id)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	if item.CityGML == nil || item.MaxLOD == nil {
		return nil, rerror.ErrNotFound
	}

	asset, err := c.PublicAPIClient.GetAsset(ctx, item.CityGML.ID)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	maxlod, err := getMaxLOD(ctx, item.MaxLOD.URL)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	files := lo.FilterMap(asset.Files, func(u string, _ int) (*url.URL, bool) {
		res, err := url.Parse(u)
		return res, err == nil && path.Ext(res.Path) == ".gml"
	})

	return maxlod.Map().Files(files), nil
}

func (c *CMS) DatasetsWithIntegrationAPI(ctx context.Context, model string) (*DatasetResponse, error) {
	items, err := c.IntegrationAPIClient.GetItemsByKey(ctx, c.Project, model, true)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	return ItemsFromIntegration(items.Items).DatasetResponse(), nil
}

func (c *CMS) FilesWithIntegrationAPI(ctx context.Context, model, id string) (any, error) {
	item, err := c.IntegrationAPIClient.GetItem(ctx, id, true)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	iitem := ItemFromIntegration(item)
	if iitem.CityGML == nil || iitem.MaxLOD == nil || !iitem.IsPublic() {
		return nil, rerror.ErrNotFound
	}

	asset, err := c.IntegrationAPIClient.Asset(ctx, iitem.CityGML.ID)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	if asset.File == nil {
		return nil, rerror.ErrNotFound
	}

	assetURL, err := url.Parse(asset.URL)
	if asset.File == nil {
		return nil, rerror.ErrInternalBy(fmt.Errorf("failed to parse asset url %s: %w", asset.URL, err))
	}

	assetBase := util.CloneRef(assetURL)
	assetBase.Path = path.Dir(assetBase.Path)

	maxlod, err := getMaxLOD(ctx, iitem.MaxLOD.URL)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	files := lo.FilterMap(asset.File.Paths(), func(u string, _ int) (*url.URL, bool) {
		if path.Ext(u) != ".gml" {
			return nil, false
		}
		fu := util.CloneRef(assetBase)
		fu.Path = path.Join(fu.Path, u)
		return fu, true
	})

	return maxlod.Map().Files(files), nil
}
