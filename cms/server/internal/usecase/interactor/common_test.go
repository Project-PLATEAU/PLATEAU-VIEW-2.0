package interactor

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway/gatewaymock"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/task"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCommon_createEvent(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()
	uID := user.NewID()
	a := asset.New().NewID().Thread(asset.NewThreadID()).
		Project(project.NewID()).Size(100).CreatedByUser(uID).NewUUID().MustBuild()
	workspace := user.NewWorkspace().NewID().MustBuild()
	wh := integration.NewWebhookBuilder().NewID().Name("aaa").
		Url(lo.Must(url.Parse("https://example.com"))).Active(true).
		Trigger(integration.WebhookTrigger{event.AssetCreate: true}).MustBuild()
	integration := integration.New().NewID().Developer(uID).Name("xxx").Webhook([]*integration.Webhook{wh}).MustBuild()
	lo.Must0(workspace.Members().AddIntegration(integration.ID(), user.RoleOwner, uID))

	db := memory.New()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mRunner := gatewaymock.NewMockTaskRunner(mockCtrl)
	gw := &gateway.Container{
		TaskRunner: mRunner,
	}

	ctx := context.Background()
	lo.Must0(db.Workspace.Save(ctx, workspace))
	lo.Must0(db.Integration.Save(ctx, integration))
	mRunner.EXPECT().Run(ctx, gomock.Any()).Times(1).Return(nil)

	ev, err := createEvent(ctx, db, gw, Event{
		Workspace: workspace.ID(),
		Type:      event.Type(event.AssetCreate),
		Object:    a,
		Operator:  operator.OperatorFromUser(uID),
	})
	assert.NoError(t, err)
	expectedEv := event.New[any]().ID(ev.ID()).Timestamp(now).Type(event.AssetCreate).Operator(operator.OperatorFromUser(uID)).Object(a).MustBuild()
	assert.Equal(t, expectedEv, ev)
}

func TestCommon_webhook(t *testing.T) {
	now := time.Now()
	uID := user.NewID()
	a := asset.New().NewID().Thread(asset.NewThreadID()).NewUUID().
		Project(project.NewID()).Size(100).CreatedByUser(uID).
		MustBuild()
	workspace := user.NewWorkspace().NewID().MustBuild()
	wh := integration.NewWebhookBuilder().NewID().Name("aaa").
		Url(lo.Must(url.Parse("https://example.com"))).Active(true).
		Trigger(integration.WebhookTrigger{event.AssetCreate: true}).MustBuild()
	integration := integration.New().NewID().Developer(uID).Name("xxx").
		Webhook([]*integration.Webhook{wh}).MustBuild()
	lo.Must0(workspace.Members().AddIntegration(integration.ID(), user.RoleOwner, uID))
	ev := event.New[any]().NewID().Timestamp(now).Type(event.AssetCreate).
		Operator(operator.OperatorFromUser(uID)).Object(a).MustBuild()

	db := memory.New()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mRunner := gatewaymock.NewMockTaskRunner(mockCtrl)
	gw := &gateway.Container{
		TaskRunner: mRunner,
	}

	ctx := context.Background()
	// no workspace
	err := webhook(ctx, db, gw, Event{Workspace: workspace.ID()}, ev)
	assert.Error(t, err)

	lo.Must0(db.Workspace.Save(ctx, workspace))
	// no webhook call since no integrtaion
	mRunner.EXPECT().Run(ctx, task.WebhookPayload{
		Webhook: wh,
		Event:   ev,
	}.Payload()).Times(0).Return(nil)
	err = webhook(ctx, db, gw, Event{Workspace: workspace.ID()}, ev)
	assert.NoError(t, err)

	lo.Must0(db.Integration.Save(ctx, integration))
	mRunner.EXPECT().Run(ctx, task.WebhookPayload{
		Webhook: wh,
		Event:   ev,
	}.Payload()).Times(1).Return(nil)
	err = webhook(ctx, db, gw, Event{Workspace: workspace.ID()}, ev)
	assert.NoError(t, err)
}
