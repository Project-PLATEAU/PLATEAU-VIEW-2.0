package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
)

var (
	projectIndexes       = []string{"alias", "workspace"}
	projectUniqueIndexes = []string{"id"}
)

type ProjectRepo struct {
	client *mongox.Collection
	f      repo.WorkspaceFilter
}

func NewProject(client *mongox.Client) repo.Project {
	return &ProjectRepo{client: client.WithCollection("project")}
}

func (r *ProjectRepo) Init() error {
	return createIndexes(context.Background(), r.client, projectIndexes, projectUniqueIndexes)
}

func (r *ProjectRepo) Filtered(f repo.WorkspaceFilter) repo.Project {
	return &ProjectRepo{
		client: r.client,
		f:      r.f.Merge(f),
	}
}

func (r *ProjectRepo) FindByID(ctx context.Context, id id.ProjectID) (*project.Project, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	})
}

func (r *ProjectRepo) FindByIDs(ctx context.Context, ids id.ProjectIDList) (project.List, error) {
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

func (r *ProjectRepo) FindByWorkspaces(ctx context.Context, ids id.WorkspaceIDList, pagination *usecasex.Pagination) (project.List, *usecasex.PageInfo, error) {
	return r.paginate(ctx, bson.M{
		"workspace": bson.M{
			"$in": ids.Strings(),
		},
	}, pagination)
}

func (r *ProjectRepo) FindByIDOrAlias(ctx context.Context, id project.IDOrAlias) (*project.Project, error) {
	pid := id.ID()
	alias := id.Alias()
	if pid == nil && (alias == nil || *alias == "") {
		return nil, rerror.ErrNotFound
	}

	filter := bson.M{}
	if pid != nil {
		filter["id"] = pid.String()
	}
	if alias != nil {
		filter["alias"] = *alias
	}

	return r.findOne(ctx, filter)
}

func (r *ProjectRepo) FindByPublicName(ctx context.Context, name string) (*project.Project, error) {
	if name == "" {
		return nil, rerror.ErrNotFound
	}
	return r.findOne(ctx, bson.M{
		"alias": name,
	})
}

func (r *ProjectRepo) CountByWorkspace(ctx context.Context, workspace id.WorkspaceID) (int, error) {
	count, err := r.client.Count(ctx, bson.M{
		"workspace": workspace.String(),
	})
	return int(count), err
}

func (r *ProjectRepo) Save(ctx context.Context, project *project.Project) error {
	if !r.f.CanWrite(project.Workspace()) {
		return repo.ErrOperationDenied
	}
	doc, id := mongodoc.NewProject(project)
	return r.client.SaveOne(ctx, id, doc)
}

func (r *ProjectRepo) Remove(ctx context.Context, id id.ProjectID) error {
	return r.client.RemoveOne(ctx, r.writeFilter(bson.M{"id": id.String()}))
}

func (r *ProjectRepo) find(ctx context.Context, filter any) (project.List, error) {
	c := mongodoc.NewProjectConsumer()
	if err := r.client.Find(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result, nil
}

func (r *ProjectRepo) findOne(ctx context.Context, filter any) (*project.Project, error) {
	c := mongodoc.NewProjectConsumer()
	if err := r.client.FindOne(ctx, r.readFilter(filter), c); err != nil {
		return nil, err
	}
	return c.Result[0], nil
}

func (r *ProjectRepo) paginate(ctx context.Context, filter bson.M, pagination *usecasex.Pagination) (project.List, *usecasex.PageInfo, error) {
	c := mongodoc.NewProjectConsumer()
	pageInfo, err := r.client.Paginate(ctx, r.readFilter(filter), nil, pagination, c)
	if err != nil {
		return nil, nil, rerror.ErrInternalBy(err)
	}
	return c.Result, pageInfo, nil
}

func filterProjects(ids []id.ProjectID, rows project.List) project.List {
	res := make(project.List, 0, len(ids))
	for _, id := range ids {
		for _, r := range rows {
			if r.ID() == id {
				res = append(res, r)
				break
			}
		}
	}
	return res
}

func (r *ProjectRepo) readFilter(filter any) any {
	return applyWorkspaceFilter(filter, r.f.Readable)
}

func (r *ProjectRepo) writeFilter(filter any) any {
	return applyWorkspaceFilter(filter, r.f.Writable)
}
