package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
)

type Schema interface {
	Filtered(filter WorkspaceFilter) Schema
	FindByIDs(context.Context, id.SchemaIDList) (schema.List, error)
	FindByID(context.Context, id.SchemaID) (*schema.Schema, error)
	Save(context.Context, *schema.Schema) error
	Remove(context.Context, id.SchemaID) error
}
