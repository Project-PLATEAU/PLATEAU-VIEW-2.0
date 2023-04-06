package mongo

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	schemaIndexes = []string{"id"}
)

type Schema struct {
	client *mongox.Collection
	f      repo.WorkspaceFilter
}

func NewSchema(client *mongox.Client) repo.Schema {
	return &Schema{client: client.WithCollection("schema")}
}

func (r *Schema) Init() error {
	return createIndexes(context.Background(), r.client, schemaIndexes, nil)
}

func (r *Schema) Filtered(f repo.WorkspaceFilter) repo.Schema {
	return &Schema{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *Schema) FindByID(ctx context.Context, schemaID id.SchemaID) (*schema.Schema, error) {
	return r.findOne(ctx, bson.M{
		"id": schemaID.String(),
	})
}

func (r *Schema) FindByIDs(ctx context.Context, ids id.SchemaIDList) (schema.List, error) {
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

	// prepare filters the results and sorts them according to original ids list
	return util.Map(ids, func(sid id.SchemaID) *schema.Schema {
		s, ok := lo.Find(res, func(s *schema.Schema) bool {
			return s.ID() == sid
		})
		if !ok {
			return nil
		}
		return s
	}), nil
}

func (r *Schema) Save(ctx context.Context, schema *schema.Schema) error {
	if !r.f.CanWrite(schema.Workspace()) {
		return repo.ErrOperationDenied
	}
	doc, sId := mongodoc.NewSchema(schema)
	return r.client.SaveOne(ctx, sId, doc)
}

func (r *Schema) Remove(ctx context.Context, schemaID id.SchemaID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{"id": schemaID.String()}))
}

func (r *Schema) findOne(ctx context.Context, filter any) (*schema.Schema, error) {
	c := mongodoc.NewSchemaConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *Schema) find(ctx context.Context, filter any) (schema.List, error) {
	c := mongodoc.NewSchemaConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result, nil
}

func (r *Schema) readFilter(filter any) any {
	return applyWorkspaceFilter(filter, r.f.Readable)
}

func (r *Schema) writeFilter(filter any) any {
	return applyWorkspaceFilter(filter, r.f.Writable)
}
