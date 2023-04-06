package value

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Multiple struct {
	t Type
	v []*Value
}

func NewMultiple(t Type, v []any) *Multiple {
	if t == TypeUnknown {
		return nil
	}
	vs := lo.FilterMap(v, func(w any, _ int) (*Value, bool) {
		vv := New(t, w)
		return vv, vv != nil
	})
	return &Multiple{
		t: t,
		v: vs,
	}
}

func MultipleFrom(t Type, v []*Value) *Multiple {
	if t == TypeUnknown {
		return nil
	}
	ok := lo.EveryBy(v, func(w *Value) bool {
		return w.Type() == t
	})
	if !ok {
		return nil
	}
	return &Multiple{
		t: t,
		v: slices.Clone(v),
	}
}

func (m *Multiple) IsEmpty() bool {
	return m == nil || m.t == TypeUnknown || lo.EveryBy(m.v, func(w *Value) bool {
		return w.IsEmpty()
	})
}

func (m *Multiple) Clone() *Multiple {
	if m == nil {
		return nil
	}
	v := lo.Map(m.v, func(v *Value, _ int) *Value {
		return v.Clone()
	})
	return &Multiple{
		t: m.t,
		v: v,
	}
}

func (m *Multiple) Values() []*Value {
	if m == nil {
		return nil
	}
	return slices.Clone(m.v)
}

func (m *Multiple) Len() int {
	if m == nil {
		return 0
	}
	return len(m.v)
}

func (m *Multiple) First() *Value {
	if m == nil || len(m.v) == 0 {
		return nil
	}
	return m.v[0]
}

func (m *Multiple) Type() Type {
	if m == nil {
		return TypeUnknown
	}
	return m.t
}

// Interface converts the value into generic representation
func (m *Multiple) Interface() []any {
	if m.IsEmpty() {
		return []any{}
	}

	return lo.Map(m.v, func(v *Value, i int) any {
		return v.Interface()
	})
}

func (m *Multiple) Validate() bool {
	return lo.EveryBy(m.v, func(v *Value) bool {
		return v.Validate()
	})
}

func (m *Multiple) Equal(w *Multiple) bool {
	if m == nil && w == nil {
		return true
	}
	if m == nil || w == nil || m.t != w.t || len(m.v) != len(w.v) {
		return false
	}
	z := lo.Zip2(m.v, w.v)
	return lo.EveryBy(z, func(vv lo.Tuple2[*Value, *Value]) bool {
		return vv.A.Equal(vv.B)
	})
}

func (m *Multiple) Cast(t Type) *Multiple {
	if m == nil {
		return nil
	}
	if m.t == t {
		return m.Clone()
	}
	v := lo.FilterMap(m.v, func(v *Value, _ int) (*Value, bool) {
		c := v.Cast(t)
		return c, c != nil
	})
	return &Multiple{
		t: t,
		v: v,
	}
}
