package mongo

import (
	"context"

	"github.com/reearth/reearth/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth/server/pkg/workspace"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	policyIndexes       = []string{}
	policyUniqueIndexes = []string{"id"}
)

type Policy struct {
	client *mongox.ClientCollection
}

func NewPolicy(c *mongox.Client) *Policy {
	return &Policy{
		client: c.WithCollection("policy"),
	}
}

func (r *Policy) Init() error {
	return createIndexes(context.Background(), r.client, policyIndexes, policyUniqueIndexes)
}

func (r *Policy) FindByID(ctx context.Context, id workspace.PolicyID) (*workspace.Policy, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *Policy) FindByIDs(ctx context.Context, ids []workspace.PolicyID) ([]*workspace.Policy, error) {
	return r.find(ctx, bson.M{
		"id": bson.M{
			"$in": util.Map(ids, func(id workspace.PolicyID) string { return id.String() }),
		},
	})
}

func (r *Policy) findOne(ctx context.Context, filter interface{}) (*workspace.Policy, error) {
	c := mongodoc.NewPolicyConsumer()
	if err := r.client.FindOne(ctx, filter, c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *Policy) find(ctx context.Context, filter interface{}) ([]*workspace.Policy, error) {
	c := mongodoc.NewPolicyConsumer()
	if err := r.client.Find(ctx, filter, c); err != nil {
		return nil, err
	}
	return c.Result, nil
}
