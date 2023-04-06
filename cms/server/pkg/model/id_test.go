package model

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestIDOrKey(t *testing.T) {
	i := NewID()
	assert.Equal(t, &i, IDOrKey(i.String()).ID())
	assert.Empty(t, IDOrKey(i.String()).Key())
	assert.Nil(t, IDOrKey("aaa").ID())
	assert.Equal(t, lo.ToPtr("aaa"), IDOrKey("aaa").Key())
}
