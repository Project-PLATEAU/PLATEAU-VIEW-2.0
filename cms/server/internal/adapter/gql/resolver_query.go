package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*gqlmodel.Me, error) {
	u := getUser(ctx)
	if u == nil {
		return nil, nil
	}
	return gqlmodel.ToMe(u), nil
}

func (r *queryResolver) Node(ctx context.Context, i gqlmodel.ID, typeArg gqlmodel.NodeType) (gqlmodel.Node, error) {
	dataloaders := dataloaders(ctx)
	switch typeArg {
	case gqlmodel.NodeTypeUser:
		result, err := dataloaders.User.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeWorkspace:
		result, err := dataloaders.Workspace.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeProject:
		result, err := dataloaders.Project.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeAsset:
		result, err := dataloaders.Asset.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeModel:
		result, err := dataloaders.Model.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeSchema:
		result, err := dataloaders.Schema.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeItem:
		result, err := dataloaders.Item.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeIntegration:
		result, err := dataloaders.Integration.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	case gqlmodel.NodeTypeRequest:
		result, err := dataloaders.Request.Load(i)
		if result == nil {
			return nil, nil
		}
		return result, err
	}
	return nil, nil
}

func (r *queryResolver) Nodes(ctx context.Context, ids []gqlmodel.ID, typeArg gqlmodel.NodeType) ([]gqlmodel.Node, error) {
	dataloaders := dataloaders(ctx)
	switch typeArg {
	case gqlmodel.NodeTypeUser:
		data, err := dataloaders.User.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeWorkspace:
		data, err := dataloaders.Workspace.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeProject:
		data, err := dataloaders.Project.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeAsset:
		data, err := dataloaders.Asset.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeModel:
		data, err := dataloaders.Model.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeSchema:
		data, err := dataloaders.Schema.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeItem:
		data, err := dataloaders.Item.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeRequest:
		data, err := dataloaders.Request.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case gqlmodel.NodeTypeIntegration:
		data, err := dataloaders.Integration.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]gqlmodel.Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	}
	return nil, nil
}

func (r *queryResolver) SearchUser(ctx context.Context, nameOrEmail string) (*gqlmodel.User, error) {
	return loaders(ctx).User.SearchUser(ctx, nameOrEmail)
}

func (r *queryResolver) Projects(ctx context.Context, workspaceID gqlmodel.ID, p *gqlmodel.Pagination) (*gqlmodel.ProjectConnection, error) {
	return loaders(ctx).Project.FindByWorkspace(ctx, workspaceID, p)
}

func (r *queryResolver) CheckProjectAlias(ctx context.Context, alias string) (*gqlmodel.ProjectAliasAvailability, error) {
	return loaders(ctx).Project.CheckAlias(ctx, alias)
}

func (r *queryResolver) AssetFile(ctx context.Context, assetId gqlmodel.ID) (*gqlmodel.AssetFile, error) {
	id, err := id.AssetIDFrom(string(assetId))
	if err != nil {
		return nil, err
	}
	f, err := usecases(ctx).Asset.FindFileByID(ctx, id, getOperator(ctx))
	if err != nil {
		return nil, err
	}
	return gqlmodel.ToAssetFile(f), nil
}

func (r *queryResolver) Models(ctx context.Context, projectID gqlmodel.ID, p *gqlmodel.Pagination) (*gqlmodel.ModelConnection, error) {
	return loaders(ctx).Model.FindByProject(ctx, projectID, p)
}

func (r *queryResolver) CheckModelKeyAvailability(ctx context.Context, projectID gqlmodel.ID, key string) (*gqlmodel.KeyAvailability, error) {
	return loaders(ctx).Model.CheckKey(ctx, projectID, key)
}

func (r *queryResolver) VersionsByItem(ctx context.Context, itemID gqlmodel.ID) ([]*gqlmodel.VersionedItem, error) {
	return loaders(ctx).Item.FindVersionedItems(ctx, itemID)
}

func (r *queryResolver) Items(ctx context.Context, schemaID gqlmodel.ID, sort *gqlmodel.ItemSort, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	return loaders(ctx).Item.FindBySchema(ctx, schemaID, sort, p)
}

func (r *queryResolver) Assets(ctx context.Context, projectId gqlmodel.ID, keyword *string, sort *gqlmodel.AssetSort, pagination *gqlmodel.Pagination) (*gqlmodel.AssetConnection, error) {
	return loaders(ctx).Asset.FindByProject(ctx, projectId, keyword, sort, pagination)
}

func (r *queryResolver) ItemsByProject(ctx context.Context, projectID gqlmodel.ID, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	return loaders(ctx).Item.FindByProject(ctx, projectID, p)
}

func (r *queryResolver) SearchItem(ctx context.Context, query gqlmodel.ItemQuery, sort *gqlmodel.ItemSort, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	return loaders(ctx).Item.Search(ctx, query, sort, p)
}

func (r *queryResolver) Requests(ctx context.Context, projectID gqlmodel.ID, key *string, state []gqlmodel.RequestState, reviewer, createdBy *gqlmodel.ID, p *gqlmodel.Pagination, sort *gqlmodel.Sort) (*gqlmodel.RequestConnection, error) {
	return loaders(ctx).Request.FindByProject(ctx, projectID, key, state, reviewer, createdBy, p, sort)
}
