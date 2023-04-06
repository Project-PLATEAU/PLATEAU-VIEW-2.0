package event

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	u := user.New().NewID().Email("hoge@example.com").Name("John").MustBuild()
	a := asset.New().NewID().Thread(id.NewThreadID()).NewUUID().
		Project(project.NewID()).Size(100).CreatedByUser(u.ID()).MustBuild()
	now := time.Now()
	eID := NewID()
	ev := New[*asset.Asset]().ID(eID).Timestamp(now).Type(AssetCreate).
		Operator(operator.OperatorFromUser(u.ID())).Object(a).MustBuild()

	assert.Equal(t, eID, ev.ID())
	assert.Equal(t, Type(AssetCreate), ev.Type())
	assert.Equal(t, operator.OperatorFromUser(u.ID()), ev.Operator())
	assert.Equal(t, a, ev.Object())
	assert.Equal(t, now, ev.Timestamp())
	assert.Equal(t, ev, ev.Clone())
	assert.NotSame(t, ev, ev.Clone())
}
