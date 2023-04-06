package interactor

import (
	"context"
	"errors"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
)

type Integration struct {
	repos    *repo.Container
	gateways *gateway.Container
}

func NewIntegration(r *repo.Container, g *gateway.Container) interfaces.Integration {
	return &Integration{
		repos:    r,
		gateways: g,
	}
}

func (i Integration) FindByMe(ctx context.Context, operator *usecase.Operator) (integration.List, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (integration.List, error) {
			in, err := i.repos.Integration.FindByUser(ctx, *operator.User)
			if err != nil {
				return nil, err
			}
			return in, nil
		})
}

func (i Integration) FindByIDs(ctx context.Context, ids id.IntegrationIDList, operator *usecase.Operator) (integration.List, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (integration.List, error) {
			in, err := i.repos.Integration.FindByIDs(ctx, ids)
			if err != nil {
				return nil, err
			}
			return in, err
		})
}

func (i Integration) Create(ctx context.Context, param interfaces.CreateIntegrationParam, operator *usecase.Operator) (*integration.Integration, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (*integration.Integration, error) {
			in, err := integration.New().
				NewID().
				Type(param.Type).
				Developer(*operator.User).
				Name(param.Name).
				Description(lo.FromPtr(param.Description)).
				GenerateToken().
				LogoUrl(&param.Logo).
				Build()
			if err != nil {
				return nil, err
			}

			if err := i.repos.Integration.Save(ctx, in); err != nil {
				return nil, err
			}

			return in, nil
		})
}

func (i Integration) Update(ctx context.Context, iId id.IntegrationID, param interfaces.UpdateIntegrationParam, operator *usecase.Operator) (*integration.Integration, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (*integration.Integration, error) {
			in, err := i.repos.Integration.FindByID(ctx, iId)
			if err != nil {
				return nil, err
			}

			if in.Developer() != *operator.User {
				return nil, interfaces.ErrOperationDenied
			}

			if param.Name != nil {
				in.SetName(*param.Name)
			}

			if param.Description != nil {
				in.SetDescription(*param.Description)
			}

			if param.Logo != nil {
				in.SetLogoUrl(param.Logo)
			}

			in.SetUpdatedAt(time.Now())
			if err := i.repos.Integration.Save(ctx, in); err != nil {
				return nil, err
			}

			return in, nil
		})
}

func (i Integration) Delete(ctx context.Context, integrationId id.IntegrationID, operator *usecase.Operator) error {
	if operator.User == nil {
		return interfaces.ErrInvalidOperator
	}
	return Run0(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) error {
			in, err := i.repos.Integration.FindByID(ctx, integrationId)
			if err != nil {
				return err
			}
			if in.Developer() != *operator.User {
				return interfaces.ErrOperationDenied
			}

			// remove the integration from the connected workspaces
			ws, err := i.repos.Workspace.FindByIntegration(ctx, integrationId)
			if err != nil && !errors.Is(err, rerror.ErrNotFound) {
				return err
			}
			for _, w := range ws {
				if err := w.Members().DeleteIntegration(integrationId); err != nil {
					return err
				}
			}
			if err := i.repos.Workspace.SaveAll(ctx, ws); err != nil {
				return err
			}

			return i.repos.Integration.Remove(ctx, integrationId)
		})
}

func (i Integration) CreateWebhook(ctx context.Context, iId id.IntegrationID, param interfaces.CreateWebhookParam, operator *usecase.Operator) (*integration.Webhook, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (*integration.Webhook, error) {
			in, err := i.repos.Integration.FindByID(ctx, iId)
			if err != nil {
				return nil, err
			}

			if in.Developer() != *operator.User {
				return nil, interfaces.ErrOperationDenied
			}

			w, err := integration.NewWebhookBuilder().
				NewID().
				Name(param.Name).
				Url(&param.URL).
				Active(param.Active).
				Secret(param.Secret).
				Trigger(integration.WebhookTrigger(*param.Trigger)).
				Build()

			if err != nil {
				return nil, err
			}

			in.AddWebhook(w)

			in.SetUpdatedAt(time.Now())
			if err := i.repos.Integration.Save(ctx, in); err != nil {
				return nil, err
			}

			return w, nil
		})
}

func (i Integration) UpdateWebhook(ctx context.Context, iId id.IntegrationID, wId id.WebhookID, param interfaces.UpdateWebhookParam, operator *usecase.Operator) (*integration.Webhook, error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) (*integration.Webhook, error) {
			in, err := i.repos.Integration.FindByID(ctx, iId)
			if err != nil {
				return nil, err
			}

			if in.Developer() != *operator.User {
				return nil, interfaces.ErrOperationDenied
			}

			w, ok := in.Webhook(wId)
			if !ok {
				return nil, rerror.ErrNotFound
			}

			if param.Name != nil {
				w.SetName(*param.Name)
			}

			if param.URL != nil {
				w.SetURL(param.URL)
			}

			if param.Active != nil {
				w.SetActive(*param.Active)
			}

			if param.Trigger != nil {
				w.SetTrigger(integration.WebhookTrigger(*param.Trigger))
			}

			if param.Secret != nil {
				w.SetSecret(*param.Secret)
			}

			w.SetUpdatedAt(time.Now())

			in.UpdateWebhook(wId, w)

			in.SetUpdatedAt(time.Now())
			if err := i.repos.Integration.Save(ctx, in); err != nil {
				return nil, err
			}

			return w, nil
		})
}

func (i Integration) DeleteWebhook(ctx context.Context, iId id.IntegrationID, wId id.WebhookID, operator *usecase.Operator) error {
	if operator.User == nil {
		return interfaces.ErrInvalidOperator
	}
	return Run0(ctx, operator, i.repos, Usecase().Transaction(),
		func(ctx context.Context) error {
			in, err := i.repos.Integration.FindByID(ctx, iId)
			if err != nil {
				return err
			}

			if in.Developer() != *operator.User {
				return interfaces.ErrOperationDenied
			}

			ok := in.DeleteWebhook(wId)
			if !ok {
				return rerror.ErrNotFound
			}

			in.SetUpdatedAt(time.Now())
			if err := i.repos.Integration.Save(ctx, in); err != nil {
				return err
			}

			return nil
		})
}
