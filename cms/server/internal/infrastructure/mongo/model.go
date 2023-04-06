package mongo

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	modelIndexes       = []string{"project", "workspace", "key"}
	modelUniqueIndexes = []string{"id"}
)

type Model struct {
	client *mongox.Collection
	f      repo.ProjectFilter
}

func NewModel(client *mongox.Client) repo.Model {
	return &Model{client: client.WithCollection("model")}
}

func (r *Model) Filtered(f repo.ProjectFilter) repo.Model {
	return &Model{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *Model) Init() error {
	return createIndexes(context.Background(), r.client, modelIndexes, modelUniqueIndexes)
}

func (r *Model) FindByID(ctx context.Context, modelID id.ModelID) (*model.Model, error) {
	return r.findOne(ctx, bson.M{
		"id": modelID.String(),
	})
}

func (r *Model) FindByIDs(ctx context.Context, ids id.ModelIDList) (model.List, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	res, err := r.find(ctx, bson.M{
		"id": bson.M{
			"$in": ids.Strings(),
		},
	})
	if err != nil {
		return nil, err
	}
	return prepare(ids, res), nil
}

func (r *Model) FindByProject(ctx context.Context, pid id.ProjectID, pagination *usecasex.Pagination) (model.List, *usecasex.PageInfo, error) {
	if !r.f.CanRead(pid) {
		return nil, usecasex.EmptyPageInfo(), nil
	}

	return r.paginate(ctx, bson.M{
		"project": pid.String(),
	}, pagination)
}

func (r *Model) FindByKey(ctx context.Context, projectID id.ProjectID, key string) (*model.Model, error) {
	if len(key) == 0 {
		return nil, rerror.ErrNotFound
	}
	if !r.f.CanRead(projectID) {
		return nil, repo.ErrOperationDenied
	}

	return r.findOne(ctx, bson.M{
		"key":     key,
		"project": projectID.String(),
	})
}

func (r *Model) FindByIDOrKey(ctx context.Context, projectID id.ProjectID, q model.IDOrKey) (*model.Model, error) {
	mid := q.ID()
	key := q.Key()
	if mid == nil && (key == nil || *key == "") {
		return nil, rerror.ErrNotFound
	}

	filter := bson.M{
		"project": projectID.String(),
	}
	if mid != nil {
		filter["id"] = mid.String()
	}
	if key != nil {
		filter["key"] = *key
	}

	return r.findOne(ctx, filter)
}

func (r *Model) CountByProject(ctx context.Context, projectID id.ProjectID) (int, error) {
	count, err := r.client.Count(ctx, bson.M{
		"project": projectID.String(),
	})
	return int(count), err
}

func (r *Model) Save(ctx context.Context, model *model.Model) error {
	if !r.f.CanWrite(model.Project()) {
		return repo.ErrOperationDenied
	}
	doc, mId := mongodoc.NewModel(model)
	return r.client.SaveOne(ctx, mId, doc)
}

func (r *Model) Remove(ctx context.Context, modelID id.ModelID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{"id": modelID.String()}))
}

func (r *Model) findOne(ctx context.Context, filter any) (*model.Model, error) {
	c := mongodoc.NewModelConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *Model) find(ctx context.Context, filter any) (model.List, error) {
	c := mongodoc.NewModelConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result, nil
}

func (r *Model) paginate(ctx context.Context, filter bson.M, pagination *usecasex.Pagination) (model.List, *usecasex.PageInfo, error) {
	c := mongodoc.NewModelConsumer()
	pageInfo, err := r.client.Paginate(ctx, r.readFilter(filter), nil, pagination, c)
	if err != nil {
		return nil, nil, rerror.ErrInternalBy(err)
	}
	return c.Result, pageInfo, nil
}

// prepare filters the results and sorts them according to original ids list
func prepare(ids id.ModelIDList, rows model.List) model.List {
	res := make(model.List, 0, len(ids))
	for _, mId := range ids {
		for _, r := range rows {
			if r.ID() == mId {
				res = append(res, r)
				break
			}
		}
	}
	return res
}

func (r *Model) readFilter(filter interface{}) interface{} {
	return applyProjectFilter(filter, r.f.Readable)
}

func (r *Model) writeFilter(filter interface{}) interface{} {
	return applyProjectFilter(filter, r.f.Writable)
}
