package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearthx/usecasex"
)

type RequestFilter struct {
	State     []request.State
	Keyword   *string
	Reviewer  *id.UserID
	CreatedBy *id.UserID
}

type Request interface {
	Filtered(ProjectFilter) Request
	FindByProject(context.Context, id.ProjectID, RequestFilter, *usecasex.Sort, *usecasex.Pagination) (request.List, *usecasex.PageInfo, error)
	FindByID(context.Context, id.RequestID) (*request.Request, error)
	FindByIDs(context.Context, id.RequestIDList) (request.List, error)
	FindByItems(context.Context, id.ItemIDList) (request.List, error)
	Save(context.Context, *request.Request) error
	SaveAll(context.Context, id.ProjectID, request.List) error
}
