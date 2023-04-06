package mongo

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	threadIndexes       = []string{"workspace", "author"}
	threadUniqueIndexes = []string{"id"}
)

type ThreadRepo struct {
	client *mongox.Collection
	f      repo.WorkspaceFilter
}

func NewThread(client *mongox.Client) repo.Thread {
	return &ThreadRepo{client: client.WithCollection("thread")}
}

func (r *ThreadRepo) Init() error {
	return createIndexes(context.Background(), r.client, threadIndexes, threadUniqueIndexes)

}

func (r *ThreadRepo) Save(ctx context.Context, thread *thread.Thread) error {
	if !r.f.CanWrite(thread.Workspace()) {
		return repo.ErrOperationDenied
	}
	doc, id := mongodoc.NewThread(thread)
	return r.client.SaveOne(ctx, id, doc)
}

func (r *ThreadRepo) Filtered(f repo.WorkspaceFilter) repo.Thread {
	return &ThreadRepo{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *ThreadRepo) FindByID(ctx context.Context, id id.ThreadID) (*thread.Thread, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *ThreadRepo) FindByIDs(ctx context.Context, ids id.ThreadIDList) ([]*thread.Thread, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	filter := bson.M{
		"id": bson.M{"$in": ids.Strings()},
	}
	res, err := r.find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return filterThreads(ids, res), nil
}

func (r *ThreadRepo) findOne(ctx context.Context, filter any) (*thread.Thread, error) {
	c := mongodoc.NewThreadConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *ThreadRepo) find(ctx context.Context, filter interface{}) ([]*thread.Thread, error) {
	c := mongodoc.NewThreadConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	return c.Result, nil
}

func filterThreads(ids []id.ThreadID, rows []*thread.Thread) []*thread.Thread {
	res := make([]*thread.Thread, 0, len(ids))
	for _, id := range ids {
		var r2 *thread.Thread
		for _, r := range rows {
			if r.ID() == id {
				r2 = r
				break
			}
		}
		res = append(res, r2)
	}
	return res
}

func (r *ThreadRepo) readFilter(filter any) any {
	return applyWorkspaceFilter(filter, r.f.Readable)
}

// func (r *ThreadRepo) writeFilter(filter any) any {
// 	return applyWorkspaceFilter(filter, r.f.Writable)
// }
