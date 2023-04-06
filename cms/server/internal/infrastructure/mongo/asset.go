package mongo

import (
	"context"
	"fmt"
	"regexp"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	assetIndexes       = []string{"project", "!createdat,!id"}
	assetUniqueIndexes = []string{"id"}
)

type Asset struct {
	client *mongox.Collection
	f      repo.ProjectFilter
}

func NewAsset(client *mongox.Client) repo.Asset {
	return &Asset{client: client.WithCollection("asset")}
}

func (r *Asset) Init() error {
	return createIndexes2(
		context.Background(),
		r.client,
		append(
			mongox.IndexFromKeys(assetUniqueIndexes, true),
			mongox.IndexFromKeys(assetIndexes, false)...,
		)...,
	)
}

func (r *Asset) Filtered(f repo.ProjectFilter) repo.Asset {
	return &Asset{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *Asset) FindByID(ctx context.Context, id id.AssetID) (*asset.Asset, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *Asset) FindByIDs(ctx context.Context, ids id.AssetIDList) ([]*asset.Asset, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	filter := bson.M{
		"id": bson.M{"$in": ids.Strings()},
	}
	res, err := r.find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return filterAssets(ids, res), nil
}

func (r *Asset) FindByProject(ctx context.Context, id id.ProjectID, uFilter repo.AssetFilter) ([]*asset.Asset, *usecasex.PageInfo, error) {
	if !r.f.CanRead(id) {
		return nil, usecasex.EmptyPageInfo(), nil
	}

	var filter interface{} = bson.M{
		"project": id.String(),
	}

	if uFilter.Keyword != nil {
		filter = mongox.And(filter, "", bson.M{
			"filename": bson.M{
				"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%s.*", regexp.QuoteMeta(*uFilter.Keyword)), Options: "i"},
			},
		})
	}

	return r.paginate(ctx, filter, uFilter.Sort, uFilter.Pagination)
}

func (r *Asset) UpdateProject(ctx context.Context, from, to id.ProjectID) error {
	if !r.f.CanWrite(from) || !r.f.CanWrite(to) {
		return repo.ErrOperationDenied
	}

	return r.client.UpdateMany(ctx, bson.M{
		"project": from.String(),
	}, bson.M{
		"project": to.String(),
	})
}

func (r *Asset) Save(ctx context.Context, asset *asset.Asset) error {
	if !r.f.CanWrite(asset.Project()) {
		return repo.ErrOperationDenied
	}

	doc, id := mongodoc.NewAsset(asset)
	_, err := r.client.Client().UpdateOne(ctx, bson.M{
		"id": id,
	}, bson.M{
		"$set": doc,
	}, options.Update().SetUpsert(true))
	if err != nil {
		return rerror.ErrInternalBy(err)
	}

	return nil
}

func (r *Asset) Delete(ctx context.Context, id id.AssetID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{
		"id": id.String(),
	}))
}

func (r *Asset) paginate(ctx context.Context, filter interface{}, sort *usecasex.Sort, pagination *usecasex.Pagination) ([]*asset.Asset, *usecasex.PageInfo, error) {
	c := mongodoc.NewAssetConsumer()
	pageInfo, err := r.client.Paginate(ctx, r.readFilter(filter), sort, pagination, c, options.Find().SetProjection(bson.M{"file": 0}))
	if err != nil {
		return nil, nil, rerror.ErrInternalBy(err)
	}
	return c.Result, pageInfo, nil
}

func (r *Asset) find(ctx context.Context, filter interface{}) ([]*asset.Asset, error) {
	c := mongodoc.NewAssetConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c, options.Find().SetProjection(bson.M{"file": 0})); err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	return c.Result, nil
}

func (r *Asset) findOne(ctx context.Context, filter interface{}) (*asset.Asset, error) {
	c := mongodoc.NewAssetConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c, options.FindOne().SetProjection(bson.M{"file": 0})); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func filterAssets(ids []id.AssetID, rows []*asset.Asset) []*asset.Asset {
	res := make([]*asset.Asset, 0, len(ids))
	for _, id := range ids {
		var r2 *asset.Asset
		for _, r := range rows {
			if r.ID() == id {
				r2 = r
				break
			}
		}
		res = append(res, r2)
	}
	return res
}

func (r *Asset) readFilter(filter interface{}) interface{} {
	return applyProjectFilter(filter, r.f.Readable)
}

func (r *Asset) writeFilter(filter interface{}) interface{} {
	return applyProjectFilter(filter, r.f.Writable)
}
