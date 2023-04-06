package memory

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestEvent_FindByID(t *testing.T) {
	now := time.Now()
	u := user.New().NewID().Email("hoge@example.com").Name("John").MustBuild()
	a := asset.New().NewID().Project(project.NewID()).Size(100).NewUUID().
		CreatedByUser(u.ID()).Thread(id.NewThreadID()).MustBuild()
	eID1 := event.NewID()
	ev := event.New[any]().ID(eID1).Timestamp(now).Type(event.AssetCreate).Operator(operator.OperatorFromUser(u.ID())).Object(a).MustBuild()

	r := NewEvent()
	ctx := context.Background()
	// seed
	err := r.Save(ctx, ev)
	assert.NoError(t, err)

	// found
	got, err := r.FindByID(ctx, eID1)
	assert.NoError(t, err)
	assert.Equal(t, ev, got)

	// not found
	eID2 := event.NewID()
	got2, err := r.FindByID(ctx, eID2)
	assert.Nil(t, got2)
	assert.Equal(t, rerror.ErrNotFound, err)
}

func TestEvent_Save(t *testing.T) {
	now := time.Now()
	u := user.New().NewID().Email("hoge@example.com").Name("John").MustBuild()
	a := asset.New().NewID().Project(project.NewID()).Size(100).NewUUID().
		CreatedByUser(u.ID()).Thread(id.NewThreadID()).MustBuild()
	eID1 := event.NewID()
	ev := event.New[any]().ID(eID1).Timestamp(now).Type(event.AssetCreate).Operator(operator.OperatorFromUser(u.ID())).Object(a).MustBuild()

	r := NewEvent()
	ctx := context.Background()
	err := r.Save(ctx, ev)
	assert.NoError(t, err)
	assert.Equal(t, ev, r.(*Event).data.Values()[0])

	// already exist
	_ = r.Save(ctx, ev)
	assert.Equal(t, 1, len(r.(*Event).data.Values()))
}
