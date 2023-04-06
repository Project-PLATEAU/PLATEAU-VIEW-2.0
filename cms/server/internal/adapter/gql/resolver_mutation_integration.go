package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateIntegration(ctx context.Context, input gqlmodel.CreateIntegrationInput) (*gqlmodel.IntegrationPayload, error) {
	op := getOperator(ctx)

	res, err := usecases(ctx).Integration.Create(
		ctx,
		interfaces.CreateIntegrationParam{
			Name:        input.Name,
			Description: input.Description,
			Type:        integration.TypeFrom(input.Type.String()),
			Logo:        input.LogoURL,
		},
		op,
	)
	if err != nil {
		return nil, err
	}

	return &gqlmodel.IntegrationPayload{
		Integration: gqlmodel.ToIntegration(res, op.User),
	}, nil
}

func (r *mutationResolver) UpdateIntegration(ctx context.Context, input gqlmodel.UpdateIntegrationInput) (*gqlmodel.IntegrationPayload, error) {
	iId, err := gqlmodel.ToID[id.Integration](input.IntegrationID)
	if err != nil {
		return nil, err
	}
	op := getOperator(ctx)

	res, err := usecases(ctx).Integration.Update(
		ctx,
		iId,
		interfaces.UpdateIntegrationParam{
			Name:        input.Name,
			Description: input.Description,
			Logo:        input.LogoURL,
		},
		op,
	)
	if err != nil {
		return nil, err
	}

	return &gqlmodel.IntegrationPayload{
		Integration: gqlmodel.ToIntegration(res, op.User),
	}, nil
}

func (r *mutationResolver) DeleteIntegration(ctx context.Context, input gqlmodel.DeleteIntegrationInput) (*gqlmodel.DeleteIntegrationPayload, error) {
	iId, err := gqlmodel.ToID[id.Integration](input.IntegrationID)
	if err != nil {
		return nil, err
	}

	err = usecases(ctx).Integration.Delete(ctx, iId, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteIntegrationPayload{
		IntegrationID: input.IntegrationID,
	}, nil
}

func (r *mutationResolver) CreateWebhook(ctx context.Context, input gqlmodel.CreateWebhookInput) (*gqlmodel.WebhookPayload, error) {
	iId, err := gqlmodel.ToID[id.Integration](input.IntegrationID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Integration.CreateWebhook(ctx, iId, interfaces.CreateWebhookParam{
		Name:   input.Name,
		URL:    input.URL,
		Active: input.Active,
		Trigger: &interfaces.WebhookTriggerParam{
			event.ItemCreate:      lo.FromPtrOr(input.Trigger.OnItemCreate, false),
			event.ItemUpdate:      lo.FromPtrOr(input.Trigger.OnItemUpdate, false),
			event.ItemDelete:      lo.FromPtrOr(input.Trigger.OnItemDelete, false),
			event.ItemPublish:     lo.FromPtrOr(input.Trigger.OnItemPublish, false),
			event.ItemUnpublish:   lo.FromPtrOr(input.Trigger.OnItemUnPublish, false),
			event.AssetCreate:     lo.FromPtrOr(input.Trigger.OnAssetUpload, false),
			event.AssetDecompress: lo.FromPtrOr(input.Trigger.OnAssetDecompress, false),
			event.AssetDelete:     lo.FromPtrOr(input.Trigger.OnAssetDelete, false),
		},
		Secret: input.Secret,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.WebhookPayload{
		Webhook: gqlmodel.ToWebhook(res),
	}, nil
}

func (r *mutationResolver) UpdateWebhook(ctx context.Context, input gqlmodel.UpdateWebhookInput) (*gqlmodel.WebhookPayload, error) {
	iId, wId, err := gqlmodel.ToID2[id.Integration, id.Webhook](input.IntegrationID, input.WebhookID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Integration.UpdateWebhook(ctx, iId, wId, interfaces.UpdateWebhookParam{
		Name:   input.Name,
		URL:    input.URL,
		Active: input.Active,
		Trigger: &interfaces.WebhookTriggerParam{
			event.ItemCreate:      lo.FromPtrOr(input.Trigger.OnItemCreate, false),
			event.ItemUpdate:      lo.FromPtrOr(input.Trigger.OnItemUpdate, false),
			event.ItemDelete:      lo.FromPtrOr(input.Trigger.OnItemDelete, false),
			event.ItemPublish:     lo.FromPtrOr(input.Trigger.OnItemPublish, false),
			event.ItemUnpublish:   lo.FromPtrOr(input.Trigger.OnItemUnPublish, false),
			event.AssetCreate:     lo.FromPtrOr(input.Trigger.OnAssetUpload, false),
			event.AssetDecompress: lo.FromPtrOr(input.Trigger.OnAssetDecompress, false),
			event.AssetDelete:     lo.FromPtrOr(input.Trigger.OnAssetDelete, false),
		},
		Secret: input.Secret,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.WebhookPayload{
		Webhook: gqlmodel.ToWebhook(res),
	}, nil
}

func (r *mutationResolver) DeleteWebhook(ctx context.Context, input gqlmodel.DeleteWebhookInput) (*gqlmodel.DeleteWebhookPayload, error) {
	iId, wId, err := gqlmodel.ToID2[id.Integration, id.Webhook](input.IntegrationID, input.WebhookID)
	if err != nil {
		return nil, err
	}

	err = usecases(ctx).Integration.DeleteWebhook(ctx, iId, wId, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteWebhookPayload{
		WebhookID: input.WebhookID,
	}, nil
}
