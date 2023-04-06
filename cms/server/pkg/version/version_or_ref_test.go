package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionOrRef_IsZero(t *testing.T) {
	assert.False(t, New().OrRef().IsZero())
	assert.False(t, Ref("x").OrVersion().IsZero())
	assert.True(t, Zero.OrRef().IsZero())
	assert.True(t, Ref("").OrVersion().IsZero())
	assert.True(t, VersionOrRef{}.IsZero())
}

func TestVersionOrRef_IsRef(t *testing.T) {
	assert.False(t, New().OrRef().IsRef("x"))
	assert.True(t, Ref("x").OrVersion().IsRef("x"))
	assert.False(t, Ref("y").OrVersion().IsRef("x"))
	assert.False(t, Zero.OrRef().IsRef("x"))
	assert.False(t, Ref("").OrVersion().IsRef("x"))
	assert.False(t, VersionOrRef{}.IsRef("x"))
}

func TestVersionOrRef_IsSpecialRef(t *testing.T) {
	assert.False(t, New().OrRef().IsSpecialRef())
	assert.True(t, Latest.OrVersion().IsSpecialRef())
	assert.False(t, Ref("y").OrVersion().IsSpecialRef())
	assert.False(t, Zero.OrRef().IsSpecialRef())
	assert.False(t, Ref("").OrVersion().IsSpecialRef())
	assert.False(t, VersionOrRef{}.IsSpecialRef())
}

func TestVersionOrRef_Match(t *testing.T) {
	v1, v2 := New(), New()
	called := 0
	v1.OrRef().Match(func(v Version) {
		assert.Equal(t, v1, v)
		called++
	}, func(_ Ref) {
		panic("this function should not be called!")
	})
	assert.Equal(t, 1, called)
	v1.OrRef().Match(nil, nil)

	Ref("y").OrVersion().Match(func(_ Version) {
		panic("this function should not be called!")
	}, func(r Ref) {
		assert.Equal(t, Ref("y"), r)
		called++
	})
	assert.Equal(t, 2, called)
	v2.OrRef().Match(nil, nil)

	Zero.OrRef().Match(func(_ Version) {
		panic("this function should not be called!")
	}, func(r Ref) {
		panic("this function should not be called!")
	})

	Ref("").OrVersion().Match(func(_ Version) {
		panic("this function should not be called!")
	}, func(r Ref) {
		panic("this function should not be called!")
	})

	VersionOrRef{}.Match(func(_ Version) {
		panic("this function should not be called!")
	}, func(_ Ref) {
		panic("this function should not be called!")
	})
}

func TestMatchVersionOrRef(t *testing.T) {
	v1, v2 := New(), New()

	assert.Equal(t, 1, MatchVersionOrRef(v1.OrRef(), func(v Version) int {
		assert.Equal(t, v1, v)
		return 1
	}, func(_ Ref) int { panic("this function should not be called!") }))
	v1.OrRef().Match(nil, nil)

	assert.Equal(t, 1, MatchVersionOrRef(Ref("y").OrVersion(), func(_ Version) int {
		panic("this function should not be called!")
	}, func(r Ref) int {
		assert.Equal(t, Ref("y"), r)
		return 1
	}))
	v2.OrRef().Match(nil, nil)

	assert.Equal(t, 0, MatchVersionOrRef(Zero.OrRef(), func(_ Version) int {
		panic("this function should not be called!")
	}, func(r Ref) int {
		panic("this function should not be called!")
	}))

	assert.Equal(t, 0, MatchVersionOrRef(Ref("").OrVersion(), func(_ Version) int {
		panic("this function should not be called!")
	}, func(r Ref) int {
		panic("this function should not be called!")
	}))

	assert.Equal(t, 0, MatchVersionOrRef(VersionOrRef{}, func(_ Version) int {
		panic("this function should not be called!")
	}, func(_ Ref) int {
		panic("this function should not be called!")
	}))
}
