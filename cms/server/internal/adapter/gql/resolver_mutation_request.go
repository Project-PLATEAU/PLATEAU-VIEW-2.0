package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input gqlmodel.CreateRequestInput) (*gqlmodel.RequestPayload, error) {
	pid, err := gqlmodel.ToID[id.Project](input.ProjectID)
	if err != nil {
		return nil, err
	}
	reviewers, err := util.TryMap(input.ReviewersID, gqlmodel.ToID[id.User])
	if err != nil {
		return nil, err
	}
	items, err := util.TryMap(input.Items, func(i *gqlmodel.RequestItemInput) (*request.Item, error) {
		iid, err := gqlmodel.ToID[id.Item](i.ItemID)
		if err != nil {
			return nil, err
		}

		return request.NewItem(iid)
	})
	if err != nil {
		return nil, err
	}
	uc := usecases(ctx).Request
	params := interfaces.CreateRequestParam{
		ProjectID:   pid,
		Title:       input.Title,
		Description: input.Description,
		State:       lo.ToPtr(request.StateFrom(input.State.String())),
		Reviewers:   reviewers,
		Items:       items,
	}

	res, err := uc.Create(ctx, params, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.RequestPayload{
		Request: gqlmodel.ToRequest(res),
	}, nil
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input gqlmodel.UpdateRequestInput) (*gqlmodel.RequestPayload, error) {
	rid, err := gqlmodel.ToID[id.Request](input.RequestID)
	if err != nil {
		return nil, err
	}
	reviewers, err := gqlmodel.ToIDs[id.User](input.ReviewersID)
	if err != nil {
		return nil, err
	}
	items, err := util.TryMap(input.Items, func(i *gqlmodel.RequestItemInput) (*request.Item, error) {
		iid, err := gqlmodel.ToID[id.Item](i.ItemID)
		if err != nil {
			return nil, err
		}

		return request.NewItem(iid)
	})
	if err != nil {
		return nil, err
	}
	uc := usecases(ctx).Request
	params := interfaces.UpdateRequestParam{
		RequestID:   rid,
		Title:       input.Title,
		Description: input.Description,
		State:       lo.ToPtr(request.StateFrom(input.State.String())),
		Reviewers:   reviewers,
		Items:       items,
	}

	res, err := uc.Update(ctx, params, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.RequestPayload{
		Request: gqlmodel.ToRequest(res),
	}, nil
}

func (r *mutationResolver) ApproveRequest(ctx context.Context, input gqlmodel.ApproveRequestInput) (*gqlmodel.RequestPayload, error) {
	rid, err := gqlmodel.ToID[id.Request](input.RequestID)
	if err != nil {
		return nil, err
	}
	res, err := usecases(ctx).Request.Approve(ctx, rid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.RequestPayload{
		Request: gqlmodel.ToRequest(res),
	}, nil
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, input gqlmodel.DeleteRequestInput) (*gqlmodel.DeleteRequestPayload, error) {
	rids, err := gqlmodel.ToIDs[id.Request](input.RequestsID)
	if err != nil {
		return nil, err
	}
	pid, err := gqlmodel.ToID[id.Project](input.ProjectID)
	if err != nil {
		return nil, err
	}

	err = usecases(ctx).Request.CloseAll(ctx, pid, rids, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteRequestPayload{
		Requests: input.RequestsID,
	}, nil
}
