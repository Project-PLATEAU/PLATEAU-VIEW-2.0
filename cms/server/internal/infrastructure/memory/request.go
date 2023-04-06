package memory

import (
	"context"
	"strings"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Request struct {
	data *util.SyncMap[request.ID, *request.Request]
	err  error
	f    repo.ProjectFilter
}

func NewRequest() repo.Request {
	return &Request{
		data: &util.SyncMap[id.RequestID, *request.Request]{},
	}
}

func (r *Request) Filtered(f repo.ProjectFilter) repo.Request {
	return &Request{
		data: r.data,
		f:    r.f.Merge(f),
	}
}

func (r *Request) FindByID(ctx context.Context, id id.RequestID) (*request.Request, error) {
	if r.err != nil {
		return nil, r.err
	}

	return rerror.ErrIfNil(r.data.Find(func(key request.ID, value *request.Request) bool {
		return key == id && r.f.CanRead(value.Project())
	}), rerror.ErrNotFound)
}

func (r *Request) FindByIDs(ctx context.Context, ids id.RequestIDList) (request.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	res := r.data.FindAll(func(key request.ID, value *request.Request) bool {
		return ids.Has(key) && r.f.CanRead(value.Project())
	})
	return res, nil
}

func (r *Request) FindByProject(ctx context.Context, id id.ProjectID, filter repo.RequestFilter, sort *usecasex.Sort, page *usecasex.Pagination) (request.List, *usecasex.PageInfo, error) {
	if !r.f.CanRead(id) {
		return nil, usecasex.EmptyPageInfo(), nil
	}

	if r.err != nil {
		return nil, nil, r.err
	}

	result := r.data.FindAll(func(_ request.ID, v *request.Request) bool {
		res := v.Project() == id
		if filter.State != nil {
			res = res && slices.Contains(filter.State, v.State())
		}
		if filter.Keyword != nil {
			res = res && strings.Contains(v.Title(), *filter.Keyword)
		}

		return res
	})

	var startCursor, endCursor *usecasex.Cursor
	if len(result) > 0 {
		startCursor = lo.ToPtr(usecasex.Cursor(result[0].ID().String()))
		endCursor = lo.ToPtr(usecasex.Cursor(result[len(result)-1].ID().String()))
	}

	return result, usecasex.NewPageInfo(
		int64(len(result)),
		startCursor,
		endCursor,
		true,
		true,
	), nil

}

func (r *Request) Save(ctx context.Context, a *request.Request) error {
	if !r.f.CanWrite(a.Project()) {
		return repo.ErrOperationDenied
	}

	if r.err != nil {
		return r.err
	}

	r.data.Store(a.ID(), a)
	return nil
}

func (r *Request) SaveAll(ctx context.Context, pid id.ProjectID, requests request.List) error {
	if r.err != nil {
		return r.err
	}
	entries := map[id.RequestID]*request.Request{}
	for _, req := range requests {
		entries[req.ID()] = req
	}
	r.data.StoreAll(entries)
	return nil
}

func (r *Request) FindByItems(ctx context.Context, list id.ItemIDList) (request.List, error) {
	if r.err != nil {
		return nil, r.err
	}
	res := r.data.FindAll(func(_ request.ID, value *request.Request) bool {
		return value.Items().IDs().Has(list...) && r.f.CanRead(value.Project())
	})
	return res, nil
}

func SetRequestError(r repo.Request, err error) {
	r.(*Request).err = err
}
