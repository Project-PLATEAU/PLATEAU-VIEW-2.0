package mongogit

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	client *mongox.Collection
}

func NewCollection(client *mongox.Collection) *Collection {
	return &Collection{client: client}
}

func (c *Collection) Client() *mongox.Collection {
	return c.client
}

func (c *Collection) FindOne(ctx context.Context, filter any, q version.Query, consumer mongox.Consumer) error {
	return c.client.FindOne(ctx, apply(q, filter), consumer)
}

func (c *Collection) Find(ctx context.Context, filter any, q version.Query, consumer mongox.Consumer) error {
	return c.client.Find(ctx, apply(q, filter), consumer)
}

func (c *Collection) Paginate(ctx context.Context, filter any, q version.Query, s *usecasex.Sort, p *usecasex.Pagination, consumer mongox.Consumer) (*usecasex.PageInfo, error) {
	return c.client.Paginate(ctx, apply(q, filter), s, p, consumer)
}

func (c *Collection) Count(ctx context.Context, filter any, q version.Query) (int64, error) {
	return c.client.Count(ctx, apply(q, filter))
}

func (c *Collection) SaveOne(ctx context.Context, id string, d any, parent *version.VersionOrRef) error {
	q := bson.M{
		"id":    id,
		metaKey: true,
	}
	if archived, err := c.IsArchived(ctx, q); err != nil {
		return err
	} else if archived {
		return version.ErrArchived
	}

	actualVr := lo.FromPtrOr(parent, version.Latest.OrVersion())
	meta, err := c.meta(ctx, id, actualVr.Ref())
	if err != nil {
		return err
	}

	var refs []version.Ref
	actualVr.Match(nil, func(r version.Ref) { refs = []version.Ref{r} })
	newmeta := Meta{
		ObjectID: primitive.NewObjectIDFromTimestamp(util.Now()),
		Version:  version.New(),
		Refs:     refs,
	}
	if meta == nil {
		if !actualVr.IsRef(version.Latest) {
			return rerror.ErrNotFound // invalid dest
		}
	} else {
		newmeta.Parents = []version.Version{meta.Version}
	}

	if err := version.MatchVersionOrRef(actualVr, nil, func(r version.Ref) error {
		return c.UpdateRef(ctx, id, r, nil)
	}); err != nil {
		return err
	}

	if _, err := c.client.Client().InsertOne(ctx, &Document[any]{
		Data: d,
		Meta: newmeta,
	}); err != nil {
		return rerror.ErrInternalBy(err)
	}

	return nil
}

func (c *Collection) UpdateRef(ctx context.Context, id string, ref version.Ref, dest *version.VersionOrRef) error {
	if _, err := c.client.Client().UpdateMany(ctx, bson.M{
		"id":    id,
		refsKey: bson.M{"$in": []string{ref.String()}},
	}, bson.M{
		"$pull": bson.M{refsKey: ref},
	}); err != nil {
		return rerror.ErrInternalBy(err)
	}

	if dest != nil {
		if _, err := c.client.Client().UpdateOne(ctx, apply(version.Eq(*dest), bson.M{
			"id": id,
		}), bson.M{
			"$push": bson.M{refsKey: ref},
		}); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	return nil
}

func (c *Collection) IsArchived(ctx context.Context, filter any) (bool, error) {
	cons := mongox.SliceConsumer[MetadataDocument]{}
	q := mongox.And(filter, "", bson.M{
		metaKey: true,
	})

	if err := c.client.FindOne(ctx, q, &cons); err != nil {
		if errors.Is(rerror.ErrNotFound, err) || err == io.EOF {
			return false, nil
		}
		return false, err
	}
	return cons.Result[0].Archived, nil
}

func (c *Collection) ArchiveOne(ctx context.Context, filter bson.M, archived bool) error {
	f := mongox.And(filter, "", bson.M{metaKey: true})

	if !archived {
		_, err := c.client.Client().DeleteOne(ctx, f)
		if err != nil {
			return rerror.ErrInternalBy(err)
		}
		return nil
	}

	_, err := c.client.Client().ReplaceOne(ctx, f, lo.Assign(bson.M{
		metaKey:    true,
		"archived": archived,
	}, filter), options.Replace().SetUpsert(true))
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func (c *Collection) Timestamp(ctx context.Context, filter any, q version.Query) (time.Time, error) {
	consumer := mongox.SliceConsumer[Meta]{}
	f := apply(q, filter)
	if err := c.client.Find(ctx, f, &consumer, options.Find().SetLimit(1).SetSort(bson.D{{Key: "_id", Value: -1}})); err != nil {
		return time.Time{}, err
	}
	if len(consumer.Result) == 0 {
		return time.Time{}, rerror.ErrNotFound
	}
	return consumer.Result[0].Timestamp(), nil
}

func (c *Collection) RemoveOne(ctx context.Context, filter any) error {
	return c.client.RemoveAll(ctx, filter)
}

func (c *Collection) Empty(ctx context.Context) error {
	return c.client.Client().Drop(ctx)
}

func (c *Collection) Indexes() []mongox.Index {
	return []mongox.Index{
		{
			Name:   "mongogit_id",
			Key:    bson.D{{Key: "id", Value: 1}, {Key: versionKey, Value: 1}},
			Unique: true,
		},
		{
			Name:   "mongogit_id_meta",
			Key:    bson.D{{Key: "id", Value: 1}, {Key: metaKey, Value: 1}},
			Unique: true,
			Filter: bson.M{metaKey: true},
		},
		{
			Name: "mongogit_id_refs",
			Key:  bson.D{{Key: "id", Value: 1}, {Key: refsKey, Value: 1}},
		},
		{
			Name: "mongogit_refs",
			Key:  bson.D{{Key: refsKey, Value: 1}},
		},
		{
			Name: "mongogit_parents",
			Key:  bson.D{{Key: parentsKey, Value: 1}},
		},
	}
}

func (c *Collection) meta(ctx context.Context, id string, v *version.VersionOrRef) (*Meta, error) {
	consumer := mongox.SliceConsumer[Meta]{}
	q := apply(version.Eq(lo.FromPtrOr(v, version.Latest.OrVersion())), bson.M{
		"id": id,
	})
	if err := c.client.FindOne(ctx, q, &consumer); err != nil {
		if errors.Is(rerror.ErrNotFound, err) && (v == nil || v.IsRef(version.Latest)) {
			return nil, nil
		}
		return nil, err
	}
	return &consumer.Result[0], nil
}
