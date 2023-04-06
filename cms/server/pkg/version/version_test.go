package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_OrRef(t *testing.T) {
	v := New()
	assert.Equal(t, VersionOrRef{version: v}, v.OrRef())
	assert.Equal(t, VersionOrRef{}, Zero.OrRef())
}
