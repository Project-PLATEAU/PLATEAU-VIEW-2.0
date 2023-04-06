package mongo

import (
	"context"
	"testing"

	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/stretchr/testify/assert"
)

func init() {
	mongotest.Env = "REEARTH_CMS_WORKER_DB"
}

func TestWebhook(t *testing.T) {
	ctx := context.Background()
	db := mongotest.Connect(t)(t)
	w := NewWebhook(db)

	assert.NoError(t, w.InitIndex(ctx))
	assert.NoError(t, w.InitIndex(ctx)) // second

	err := w.Delete(ctx, "hogehoge")
	assert.NoError(t, err)

	f, err := w.Get(ctx, "hogehoge")
	assert.False(t, f)
	assert.NoError(t, err)

	f, err = w.GetAndSet(ctx, "hogehoge")
	assert.False(t, f)
	assert.NoError(t, err)

	f, err = w.GetAndSet(ctx, "hogehoge")
	assert.True(t, f)
	assert.NoError(t, err)

	err = w.Delete(ctx, "hogehoge")
	assert.NoError(t, err)

	f, err = w.Get(ctx, "hogehoge")
	assert.False(t, f)
	assert.NoError(t, err)

	f, err = w.GetAndSet(ctx, "hogehoge")
	assert.False(t, f)
	assert.NoError(t, err)

	f, err = w.Get(ctx, "hogehoge")
	assert.True(t, f)
	assert.NoError(t, err)
}
