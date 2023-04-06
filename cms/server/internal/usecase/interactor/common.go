package interactor

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/task"
	"github.com/reearth/reearthx/util"
)

type ContainerConfig struct {
	SignupSecret    string
	AuthSrvUIDomain string
}

func New(r *repo.Container, g *gateway.Container, config ContainerConfig) interfaces.Container {
	return interfaces.Container{
		Asset:       NewAsset(r, g),
		Workspace:   NewWorkspace(r, g),
		User:        NewUser(r, g, config.SignupSecret, config.AuthSrvUIDomain),
		Project:     NewProject(r, g),
		Item:        NewItem(r, g),
		Request:     NewRequest(r, g),
		Model:       NewModel(r, g),
		Schema:      NewSchema(r, g),
		Integration: NewIntegration(r, g),
		Thread:      NewThread(r, g),
	}
}

type Event struct {
	Project       *project.Project
	Workspace     id.WorkspaceID
	Type          event.Type
	Operator      operator.Operator
	Object        any
	WebhookObject any
}

func (e *Event) EventProject() *event.Project {
	if e == nil || e.Project == nil {
		return nil
	}
	return &event.Project{
		ID:    e.Project.ID().String(),
		Alias: e.Project.Alias(),
	}
}

func createEvent(ctx context.Context, r *repo.Container, g *gateway.Container, e Event) (*event.Event[any], error) {
	ev, err := event.New[any]().NewID().Object(e.Object).Type(e.Type).Project(e.EventProject()).Timestamp(util.Now()).Operator(e.Operator).Build()
	if err != nil {
		return nil, err
	}

	if err := r.Event.Save(ctx, ev); err != nil {
		return nil, err
	}

	if err := webhook(ctx, r, g, e, ev); err != nil {
		return nil, err
	}

	return ev, nil
}

func webhook(ctx context.Context, r *repo.Container, g *gateway.Container, e Event, ev *event.Event[any]) error {
	if g == nil || g.TaskRunner == nil {
		return nil
	}

	ws, err := r.Workspace.FindByID(ctx, e.Workspace)
	if err != nil {
		return err
	}
	integrationIDs := ws.Members().IntegrationIDs()

	integrations, err := r.Integration.FindByIDs(ctx, integrationIDs)
	if err != nil {
		return err
	}

	for _, w := range integrations.ActiveWebhooks(ev.Type()) {
		if err := g.TaskRunner.Run(ctx, task.WebhookPayload{
			Webhook:  w,
			Event:    ev,
			Override: e.WebhookObject,
		}.Payload()); err != nil {
			return err
		}
	}

	return nil
}
