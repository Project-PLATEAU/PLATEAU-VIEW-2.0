package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	got := New()
	assert.NotNil(t, got)
	assert.NotNil(t, got.User)
	assert.NotNil(t, got.Workspace)
	assert.NotNil(t, got.Lock)
	assert.NotNil(t, got.Transaction)
}
