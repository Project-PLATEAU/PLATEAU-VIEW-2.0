package mongo

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(ctx context.Context, mc *mongo.Client, databaseName string, useTransaction bool) (*repo.Container, error) {
	if databaseName == "" {
		databaseName = "reearth_cms"
	}

	lock, err := NewLock(mc.Database(databaseName).Collection("locks"))
	if err != nil {
		return nil, err
	}

	client := mongox.NewClient(databaseName, mc)
	if useTransaction {
		client = client.WithTransaction()
	}

	c := &repo.Container{
		Asset:       NewAsset(client),
		AssetFile:   NewAssetFile(client),
		Workspace:   NewWorkspace(client),
		User:        NewUser(client),
		Transaction: client.Transaction(),
		Lock:        lock,
		Project:     NewProject(client),
		Request:     NewRequest(client),
		Item:        NewItem(client),
		Model:       NewModel(client),
		Schema:      NewSchema(client),
		Thread:      NewThread(client),
		Integration: NewIntegration(client),
		Event:       NewEvent(client),
	}

	// init
	if err := Init(c); err != nil {
		return nil, err
	}

	return c, nil
}

func NewWithDB(ctx context.Context, db *mongo.Database, useTransaction bool) (*repo.Container, error) {
	return New(ctx, db.Client(), db.Name(), useTransaction)
}

func Init(r *repo.Container) error {
	if r == nil {
		return nil
	}

	return util.Try(
		r.Asset.(*Asset).Init,
		r.Workspace.(*Workspace).Init,
		r.User.(*User).Init,
		r.Project.(*ProjectRepo).Init,
		r.Item.(*Item).Init,
		r.Model.(*Model).Init,
		r.Request.(*Request).Init,
		r.Schema.(*Schema).Init,
		r.Thread.(*ThreadRepo).Init,
		r.Integration.(*Integration).Init,
		r.Event.(*Event).Init,
	)
}

func createIndexes(ctx context.Context, c *mongox.Collection, keys, uniqueKeys []string) error {
	created, deleted, err := c.Indexes(ctx, keys, uniqueKeys)
	if len(created) > 0 || len(deleted) > 0 {
		log.Infof("mongo: %s: index deleted: %v, created: %v", c.Client().Name(), deleted, created)
	}
	return err
}

func createIndexes2(ctx context.Context, c *mongox.Collection, inputs ...mongox.Index) error {
	res, err := c.Indexes2(ctx, inputs...)
	if err == nil {
		logIndexResult(c.Client().Name(), res)
	}
	return err
}

func logIndexResult(name string, r mongox.IndexResult) {
	d := r.DeletedNames()
	u := r.UpdatedNames()
	a := r.AddedNames()
	if len(d) == 0 && len(u) == 0 && len(a) == 0 {
		return
	}
	log.Infof("mongo: %s: index deleted: %v, updated: %v, created: %v", name, d, u, a)
}

func applyWorkspaceFilter(filter interface{}, ids id.WorkspaceIDList) interface{} {
	if ids == nil {
		return filter
	}
	return mongox.And(filter, "workspace", bson.M{"$in": ids.Strings()})
}

func applyProjectFilter(filter interface{}, ids id.ProjectIDList) interface{} {
	if ids == nil {
		return filter
	}
	return mongox.And(filter, "project", bson.M{"$in": ids.Strings()})
}
