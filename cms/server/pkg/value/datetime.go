package value

import (
	"time"

	"github.com/samber/lo"
)

const TypeDateTime Type = "datetime"

type propertyDateTime struct{}

type DateTime = time.Time

var timeLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
}

func (p *propertyDateTime) ToValue(i any) (any, bool) {
	switch v := i.(type) {
	case time.Time:
		return v, true
	case string:
		for _, l := range timeLayouts {
			if tt, err := time.Parse(l, v); err == nil {
				return tt, true
			}
		}
	case *time.Time:
		if v != nil {
			return p.ToValue(*v)
		}
	case *string:
		if v != nil {
			return p.ToValue(*v)
		}
	}

	if _, ok := i.(bool); ok {
		return nil, false
	}
	if _, ok := i.(*bool); ok {
		return nil, false
	}

	if v, ok := defaultTypes.Get(TypeInteger).ToValue(i); ok {
		return time.Unix(v.(Integer), 0), true
	}

	return nil, false
}

func (*propertyDateTime) ToInterface(v any) (any, bool) {
	return v.(DateTime).Format(time.RFC3339), true
}

func (*propertyDateTime) Validate(i any) bool {
	_, ok := i.(DateTime)
	return ok
}

func (*propertyDateTime) Equal(v, w any) bool {
	vv := v.(DateTime)
	ww := w.(DateTime)
	return vv.Equal(ww)
}

func (*propertyDateTime) IsEmpty(v any) bool {
	return v.(DateTime).IsZero()
}

func (v *Value) ValueDateTime() (vv DateTime, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(DateTime)
	return
}

func (m *Multiple) ValuesDateTime() (vv []DateTime, ok bool) {
	if m == nil {
		return
	}
	vv = lo.FilterMap(m.v, func(v *Value, _ int) (DateTime, bool) {
		return v.ValueDateTime()
	})
	if len(vv) != len(m.v) {
		return nil, false
	}
	return vv, true
}
