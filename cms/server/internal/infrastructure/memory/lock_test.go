package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	ctx := context.Background()
	expected := &Lock{}
	got := NewLock()
	assert.Equal(t, expected, got)
	assert.Nil(t, got.Lock(ctx, "hoge"))
	assert.Nil(t, got.Unlock(ctx, "hoge"))
}
