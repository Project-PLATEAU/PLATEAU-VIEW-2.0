package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
)

type Integration interface {
	FindByIDs(context.Context, id.IntegrationIDList) (integration.List, error)
	FindByUser(context.Context, id.UserID) (integration.List, error)
	FindByID(context.Context, id.IntegrationID) (*integration.Integration, error)
	FindByToken(context.Context, string) (*integration.Integration, error)
	Save(context.Context, *integration.Integration) error
	Remove(context.Context, id.IntegrationID) error
}
