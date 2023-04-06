package version

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRef_Ref(t *testing.T) {
	assert.Equal(t, lo.ToPtr(Ref("x")), Ref("x").Ref())
}

func TestRef_String(t *testing.T) {
	assert.Equal(t, "x", Ref("x").String())
}

func TestRef_OrVersion(t *testing.T) {
	assert.Equal(t, VersionOrRef{ref: Ref("x")}, Ref("x").OrVersion())
	assert.Equal(t, VersionOrRef{}, Ref("").OrVersion())
}

func TestRef_OrLatest(t *testing.T) {
	assert.Equal(t, lo.ToPtr(Latest), lo.ToPtr(Latest).OrLatest())
	assert.Equal(t, lo.ToPtr(Latest), lo.ToPtr(Ref("")).OrLatest())
	assert.Equal(t, lo.ToPtr(Ref("aaa")), lo.ToPtr(Ref("aaa")).OrLatest())
}

func TestRef_IsSpecial(t *testing.T) {
	assert.False(t, Ref("x").IsSpecial())
	assert.True(t, Ref("").IsSpecial())
	assert.True(t, Latest.IsSpecial())
}

func TestRefsFrom(t *testing.T) {
	s := NewRefs("x", "y")
	assert.True(t, s.Has("x"))
	assert.True(t, s.Has("y"))
	assert.False(t, s.Has("z"))
}
