package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMultiple(t *testing.T) {
	m := NewMultiple(TypeUnknown, []any{})
	assert.Nil(t, m)

	m = NewMultiple(TypeBool, []any{true, "test"})
	assert.Equal(t, TypeBool, m.t)
	assert.Equal(t, []*Value{New(TypeBool, true)}, m.v)

	v := []*Value{New(TypeBool, true), New(TypeBool, false)}
	m = NewMultiple(TypeBool, []any{true, false})
	assert.NotNil(t, m)
	assert.Equal(t, TypeBool, m.t)
	assert.Equal(t, v, m.v)
	assert.NotSame(t, v, m.v)
}

func TestMultipleFrom(t *testing.T) {
	m := MultipleFrom(TypeUnknown, []*Value{})
	assert.Nil(t, m)

	m = MultipleFrom(TypeBool, []*Value{New(TypeBool, true), New(TypeText, "test")})
	assert.Nil(t, m)

	v := []*Value{New(TypeBool, true), New(TypeBool, false)}
	m = MultipleFrom(TypeBool, v)
	assert.NotNil(t, m)
	assert.Equal(t, TypeBool, m.t)
	assert.Equal(t, v, m.v)
	assert.NotSame(t, v, m.v)
}

func TestMultiple_IsEmpty(t *testing.T) {
	var m *Multiple = nil
	assert.True(t, m.IsEmpty())

	m = &Multiple{}
	assert.True(t, m.IsEmpty())

	m.t = TypeBool
	assert.True(t, m.IsEmpty())

	m.v = nil
	assert.True(t, m.IsEmpty())

	m.v = []*Value{}
	assert.True(t, m.IsEmpty())

	m.v = []*Value{New(TypeBool, true)}
	assert.False(t, m.IsEmpty())
}

func TestMultiple_Len(t *testing.T) {
	var m *Multiple = nil
	assert.Equal(t, 0, m.Len())

	m = &Multiple{}
	assert.Equal(t, 0, m.Len())

	m.t = TypeBool
	assert.Equal(t, 0, m.Len())

	m.v = nil
	assert.Equal(t, 0, m.Len())

	m.v = []*Value{}
	assert.Equal(t, 0, m.Len())

	m.v = []*Value{New(TypeBool, true)}
	assert.Equal(t, 1, m.Len())

	m.v = []*Value{New(TypeBool, true), New(TypeBool, true)}
	assert.Equal(t, 2, m.Len())
}

func TestMultiple_First(t *testing.T) {
	var m *Multiple = nil
	assert.Nil(t, m.First())

	m = &Multiple{}
	assert.Nil(t, m.First())

	m.t = TypeBool
	assert.Nil(t, m.First())

	m.v = nil
	assert.Nil(t, m.First())

	m.v = []*Value{}
	assert.Nil(t, m.First())

	m.v = []*Value{New(TypeBool, true)}
	assert.Equal(t, New(TypeBool, true), m.First())
}

func TestMultiple_Clone(t *testing.T) {
	m := &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}

	assert.Equal(t, m, m.Clone())
	assert.NotSame(t, m, m.Clone())

	m = nil
	assert.Equal(t, m, m.Clone())
}

func TestMultiple_Values(t *testing.T) {
	v := []*Value{New(TypeBool, true), New(TypeBool, false)}
	m := &Multiple{
		t: TypeBool,
		v: v,
	}

	assert.Equal(t, v, m.Values())
	assert.NotSame(t, v, m.Values())

	m = nil
	assert.Nil(t, m.Values())
}

func TestMultiple_Type(t *testing.T) {
	m := &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}

	assert.Equal(t, m.Type(), TypeBool)

	m = nil
	assert.Equal(t, m.Type(), TypeUnknown)
}

func TestMultiple_Interface(t *testing.T) {
	m := &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}
	assert.Equal(t, []any{true, false}, m.Interface())

	m = nil
	assert.Equal(t, []any{}, m.Interface())
}

func TestMultiple_Validate(t *testing.T) {
	m := &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}

	assert.Equal(t, m.Validate(), true)

	m = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, "test")},
	}

	assert.Equal(t, m.Validate(), false)
}

func TestMultiple_Equal(t *testing.T) {
	var m, w *Multiple = nil, nil
	assert.Equal(t, m.Equal(w), true)

	m = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}
	assert.Equal(t, m.Equal(w), false)

	w = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true)},
	}
	assert.Equal(t, m.Equal(w), false)

	w = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false), New(TypeBool, false)},
	}
	assert.Equal(t, m.Equal(w), false)

	w = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, false), New(TypeBool, false)},
	}
	assert.Equal(t, m.Equal(w), false)

	w = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}
	assert.Equal(t, m.Equal(w), true)
}

func TestMultiple_Cast(t *testing.T) {
	m := &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}
	w := &Multiple{
		t: TypeText,
		v: []*Value{New(TypeText, "true"), New(TypeText, "false")},
	}
	assert.Equal(t, w, m.Cast(TypeText))

	assert.Equal(t, m.Cast(TypeBool), m.Clone())
	assert.NotSame(t, m.Cast(TypeBool), m.Clone())

	m = nil
	assert.Nil(t, m.Cast(TypeText))
}
