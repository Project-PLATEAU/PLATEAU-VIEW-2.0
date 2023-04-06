package mongo

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	integrationIndexes       = []string{"developer"}
	integrationUniqueIndexes = []string{"id", "token"}
)

type Integration struct {
	client *mongox.Collection
}

func NewIntegration(client *mongox.Client) repo.Integration {
	return &Integration{client: client.WithCollection("integration")}
}

func (r *Integration) Init() error {
	return createIndexes(context.Background(), r.client, integrationIndexes, integrationUniqueIndexes)
}

func (r *Integration) FindByID(ctx context.Context, integrationID id.IntegrationID) (*integration.Integration, error) {
	return r.findOne(ctx, bson.M{
		"id": integrationID.String(),
	})
}

func (r *Integration) FindByToken(ctx context.Context, token string) (*integration.Integration, error) {
	return r.findOne(ctx, bson.M{
		"token": token,
	})
}

func (r *Integration) FindByIDs(ctx context.Context, ids id.IntegrationIDList) (integration.List, error) {
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
	return util.Map(ids, func(sid id.IntegrationID) *integration.Integration {
		s, ok := lo.Find(res, func(s *integration.Integration) bool {
			return s.ID() == sid
		})
		if !ok {
			return nil
		}
		return s
	}), nil
}

func (r *Integration) FindByUser(ctx context.Context, userID id.UserID) (integration.List, error) {
	return r.find(ctx, bson.M{
		"developer": userID.String(),
	})
}

func (r *Integration) Save(ctx context.Context, integration *integration.Integration) error {
	doc, sId := mongodoc.NewIntegration(integration)
	return r.client.SaveOne(ctx, sId, doc)
}

func (r *Integration) Remove(ctx context.Context, integrationID id.IntegrationID) error {
	return r.client.RemoveOne(ctx, bson.M{"id": integrationID.String()})
}

func (r *Integration) findOne(ctx context.Context, filter any) (*integration.Integration, error) {
	c := mongodoc.NewIntegrationConsumer()
	if err := r.client.FindOne(ctx, filter, c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *Integration) find(ctx context.Context, filter any) (integration.List, error) {
	c := mongodoc.NewIntegrationConsumer()
	if err := r.client.Find(ctx, filter, c); err != nil {
		return nil, err
	}
	return c.Result, nil
}
