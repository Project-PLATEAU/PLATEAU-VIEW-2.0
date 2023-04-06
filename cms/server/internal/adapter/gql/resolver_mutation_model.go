package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateModel(ctx context.Context, input gqlmodel.CreateModelInput) (*gqlmodel.ModelPayload, error) {
	pId, err := gqlmodel.ToID[id.Project](input.ProjectID)
	if err != nil {
		return nil, err
	}
	res, err := usecases(ctx).Model.Create(ctx, interfaces.CreateModelParam{
		ProjectId:   pId,
		Name:        input.Name,
		Description: input.Description,
		Key:         input.Key,
		Public:      nil,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.ModelPayload{
		Model: gqlmodel.ToModel(res),
	}, nil
}

func (r *mutationResolver) UpdateModel(ctx context.Context, input gqlmodel.UpdateModelInput) (*gqlmodel.ModelPayload, error) {
	mId, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Model.Update(ctx, interfaces.UpdateModelParam{
		ModelId:     mId,
		Name:        input.Name,
		Description: input.Description,
		Key:         input.Key,
		Public:      lo.ToPtr(input.Public),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.ModelPayload{
		Model: gqlmodel.ToModel(res),
	}, nil
}

func (r *mutationResolver) DeleteModel(ctx context.Context, input gqlmodel.DeleteModelInput) (*gqlmodel.DeleteModelPayload, error) {
	mid, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	err = usecases(ctx).Model.Delete(ctx, mid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteModelPayload{
		ModelID: input.ModelID,
	}, nil
}

func (r *mutationResolver) PublishModel(ctx context.Context, input gqlmodel.PublishModelInput) (*gqlmodel.PublishModelPayload, error) {
	mid, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	s, err := usecases(ctx).Model.Publish(ctx, mid, input.Status, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.PublishModelPayload{
		ModelID: input.ModelID,
		Status:  s,
	}, nil
}
