package mongo

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AssetFile struct {
	client *mongox.Collection
}

func NewAssetFile(client *mongox.Client) repo.AssetFile {
	return &AssetFile{client: client.WithCollection("asset")}
}

func (r *AssetFile) Init() error {
	return createIndexes(context.Background(), r.client, assetIndexes, assetUniqueIndexes)
}

func (r *AssetFile) Filtered(f repo.ProjectFilter) repo.Asset {
	return &Asset{
		client: r.client,
	}
}

func (r *AssetFile) FindByID(ctx context.Context, id id.AssetID) (*asset.File, error) {
	c := &mongodoc.AssetAndFileConsumer{}
	if err := r.client.FindOne(ctx, bson.M{
		"id": id.String(),
	}, c, options.FindOne().SetProjection(bson.M{
		"id":   1,
		"file": 1,
	})); err != nil {
		return nil, err
	}
	f := c.Result[0].File.Model()
	if f == nil {
		return nil, rerror.ErrNotFound
	}
	return f, nil
}

func (r *AssetFile) Save(ctx context.Context, id id.AssetID, file *asset.File) error {
	doc := mongodoc.NewFile(file)
	_, err := r.client.Client().UpdateOne(ctx, bson.M{
		"id": id.String(),
	}, bson.M{
		"$set": bson.M{
			"id":   id.String(),
			"file": doc,
		},
	})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return rerror.ErrNotFound
	}
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}
