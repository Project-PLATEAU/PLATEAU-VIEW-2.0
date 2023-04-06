package mongo

import (
	"context"
	"fmt"
	"regexp"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	requestIndexes       = []string{"project", "items.item"}
	requestUniqueIndexes = []string{"id"}
)

type Request struct {
	client *mongox.Collection
	f      repo.ProjectFilter
}

func NewRequest(client *mongox.Client) repo.Request {
	return &Request{client: client.WithCollection("request")}
}

func (r *Request) Init() error {
	return createIndexes(context.Background(), r.client, requestIndexes, requestUniqueIndexes)
}
func (r *Request) Filtered(f repo.ProjectFilter) repo.Request {
	return &Request{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *Request) FindByID(ctx context.Context, id id.RequestID) (*request.Request, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *Request) FindByIDs(ctx context.Context, ids id.RequestIDList) (request.List, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	filter := bson.M{
		"id": bson.M{
			"$in": ids.Strings(),
		},
	}
	res, err := r.find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return filterRequests(ids, res), nil
}

func (r *Request) FindByItems(ctx context.Context, list id.ItemIDList) (request.List, error) {

	filter := bson.M{
		"items.item": bson.M{
			"$in": list.Strings(),
		},
	}

	return r.find(ctx, filter)
}

func (r *Request) FindByProject(ctx context.Context, id id.ProjectID, uFilter repo.RequestFilter, sort *usecasex.Sort, page *usecasex.Pagination) (request.List, *usecasex.PageInfo, error) {
	if !r.f.CanRead(id) {
		return nil, usecasex.EmptyPageInfo(), nil
	}

	filter := bson.M{
		"project": id.String(),
	}

	if uFilter.Keyword != nil {
		filter["title"] = bson.M{
			"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%s.*", regexp.QuoteMeta(*uFilter.Keyword)), Options: "i"},
		}
	}

	if uFilter.State != nil {
		filter["state"] = bson.M{
			"$in": lo.Map(uFilter.State, func(s request.State, _ int) string {
				return s.String()
			}),
		}
	}

	if uFilter.CreatedBy != nil {
		filter["createdby"] = uFilter.CreatedBy.String()
	}

	if uFilter.Reviewer != nil {
		filter["reviewers"] = uFilter.Reviewer.String()
	}

	rl, p, err := r.paginate(ctx, &filter, sort, page)
	return rl.Ordered(), p, err
}

func (r *Request) Save(ctx context.Context, request *request.Request) error {
	if !r.f.CanWrite(request.Project()) {
		return repo.ErrOperationDenied
	}
	doc, id := mongodoc.NewRequest(request)
	return r.client.SaveOne(ctx, id, doc)
}

func (r *Request) SaveAll(ctx context.Context, pid id.ProjectID, requests request.List) error {
	if len(requests) == 0 {
		return nil
	}

	if !r.f.CanWrite(pid) {
		return repo.ErrOperationDenied
	}

	docs, ids := mongodoc.NewRequests(requests)
	docs2 := make([]any, 0, len(requests))
	for _, d := range docs {
		docs2 = append(docs2, d)
	}
	return r.client.SaveAll(ctx, ids, docs2)
}

func (r *Request) Remove(ctx context.Context, id id.RequestID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{
		"id": id.String()}))
}

func (r *Request) paginate(ctx context.Context, filter any, sort *usecasex.Sort, pagination *usecasex.Pagination) (request.List, *usecasex.PageInfo, error) {
	c := mongodoc.NewRequestConsumer()

	pageInfo, err := r.client.Paginate(ctx, r.readFilter(filter), sort, pagination, c)
	if err != nil {
		return nil, nil, rerror.ErrInternalBy(err)
	}
	return c.Result, pageInfo, nil
}

func (r *Request) find(ctx context.Context, filter any) (request.List, error) {
	c := mongodoc.NewRequestConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	return c.Result, nil
}

func (r *Request) findOne(ctx context.Context, filter any) (*request.Request, error) {
	c := mongodoc.NewRequestConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func filterRequests(ids id.RequestIDList, rows request.List) request.List {
	res := make(request.List, 0, len(ids))
	for _, id := range ids {
		for _, r := range rows {
			if r.ID() == id {
				res = append(res, r)
				break
			}
		}
	}
	return res
}

func (r *Request) readFilter(filter any) interface{} {
	return applyProjectFilter(filter, r.f.Readable)
}

func (r *Request) writeFilter(filter any) interface{} {
	return applyProjectFilter(filter, r.f.Writable)
}
