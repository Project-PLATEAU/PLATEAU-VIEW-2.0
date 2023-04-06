package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/reearth/reearth/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth/server/internal/usecase/repo"
	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/project"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
)

var (
	projectIndexes       = []string{"alias", "alias,publishmentstatus", "team"}
	projectUniqueIndexes = []string{"id"}
)

type Project struct {
	client *mongox.ClientCollection
	f      repo.WorkspaceFilter
}

func NewProject(client *mongox.Client) *Project {
	return &Project{
		client: client.WithCollection("project"),
	}
}

func (r *Project) Init() error {
	return createIndexes(context.Background(), r.client, projectIndexes, projectUniqueIndexes)
}

func (r *Project) Filtered(f repo.WorkspaceFilter) repo.Project {
	return &Project{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *Project) FindByID(ctx context.Context, id id.ProjectID) (*project.Project, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *Project) FindByIDs(ctx context.Context, ids id.ProjectIDList) ([]*project.Project, error) {
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
	return filterProjects(ids, res), nil
}

func (r *Project) FindByWorkspace(ctx context.Context, id id.WorkspaceID, pagination *usecasex.Pagination) ([]*project.Project, *usecasex.PageInfo, error) {
	if !r.f.CanRead(id) {
		return nil, usecasex.EmptyPageInfo(), nil
	}
	return r.paginate(ctx, bson.M{
		"team": id.String(),
	}, pagination)
}

func (r *Project) FindByPublicName(ctx context.Context, name string) (*project.Project, error) {
	if name == "" {
		return nil, rerror.ErrNotFound
	}
	return r.findOne(ctx, bson.M{
		"$or": []bson.M{
			{"alias": name, "publishmentstatus": "limited"},
			{"domains.domain": name, "publishmentstatus": "public"},
			{"alias": name, "publishmentstatus": "public"},
		},
	})
}

func (r *Project) CountByWorkspace(ctx context.Context, ws id.WorkspaceID) (int, error) {
	if !r.f.CanRead(ws) {
		return 0, repo.ErrOperationDenied
	}

	count, err := r.client.Count(ctx, bson.M{
		"team": ws.String(),
	})
	return int(count), err
}

func (r *Project) CountPublicByWorkspace(ctx context.Context, ws id.WorkspaceID) (int, error) {
	if !r.f.CanRead(ws) {
		return 0, repo.ErrOperationDenied
	}

	count, err := r.client.Count(ctx, bson.M{
		"team": ws.String(),
		"publishmentstatus": bson.M{
			"$in": []string{"public", "limited"},
		},
	})
	return int(count), err
}

func (r *Project) Save(ctx context.Context, project *project.Project) error {
	if !r.f.CanWrite(project.Workspace()) {
		return repo.ErrOperationDenied
	}
	doc, id := mongodoc.NewProject(project)
	return r.client.SaveOne(ctx, id, doc)
}

func (r *Project) Remove(ctx context.Context, id id.ProjectID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{"id": id.String()}))
}

func (r *Project) find(ctx context.Context, filter interface{}) ([]*project.Project, error) {
	c := mongodoc.NewProjectConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result, nil
}

func (r *Project) findOne(ctx context.Context, filter interface{}) (*project.Project, error) {
	c := mongodoc.NewProjectConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *Project) paginate(ctx context.Context, filter bson.M, pagination *usecasex.Pagination) ([]*project.Project, *usecasex.PageInfo, error) {
	c := mongodoc.NewProjectConsumer()
	pageInfo, err := r.client.Paginate(ctx, r.readFilter(filter), nil, pagination, c)
	if err != nil {
		return nil, nil, rerror.ErrInternalBy(err)
	}
	return c.Result, pageInfo, nil
}

func filterProjects(ids []id.ProjectID, rows []*project.Project) []*project.Project {
	res := make([]*project.Project, 0, len(ids))
	for _, id := range ids {
		var r2 *project.Project
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

func (r *Project) readFilter(filter interface{}) interface{} {
	return applyWorkspaceFilter(filter, r.f.Readable)
}

func (r *Project) writeFilter(filter interface{}) interface{} {
	return applyWorkspaceFilter(filter, r.f.Writable)
}
