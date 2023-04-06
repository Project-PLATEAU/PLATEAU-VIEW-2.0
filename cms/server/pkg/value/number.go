package value

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/samber/lo"
)

const TypeNumber Type = "number"

type propertyNumber struct{}

type Number = float64

func (p *propertyNumber) ToValue(i any) (any, bool) {
	switch v := i.(type) {
	case float64:
		return Number(v), true
	case float32:
		return Number(v), true
	case int:
		return Number(v), true
	case int8:
		return Number(v), true
	case int16:
		return Number(v), true
	case int32:
		return Number(v), true
	case int64:
		return Number(v), true
	case uint:
		return Number(v), true
	case uint8:
		return Number(v), true
	case uint16:
		return Number(v), true
	case uint32:
		return Number(v), true
	case uint64:
		return Number(v), true
	case uintptr:
		return Number(v), true
	case json.Number:
		if f, err := v.Float64(); err == nil {
			return f, true
		}
	case string:
		if vv, err := strconv.ParseFloat(v, 64); err == nil {
			return vv, true
		}
	case bool:
		if v {
			return Number(1), true
		} else {
			return Number(0), true
		}
	case time.Time:
		return Number(v.Unix()), true
	case *float64:
		if v != nil {
			return p.ToValue(*v)
		}
	case *float32:
		if v != nil {
			return p.ToValue(*v)
		}
	case *int:
		if v != nil {
			return p.ToValue(*v)
		}
	case *int8:
		if v != nil {
			return p.ToValue(*v)
		}
	case *int16:
		if v != nil {
			return p.ToValue(*v)
		}
	case *int32:
		if v != nil {
			return p.ToValue(*v)
		}
	case *int64:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uint:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uint8:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uint16:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uint32:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uint64:
		if v != nil {
			return p.ToValue(*v)
		}
	case *uintptr:
		if v != nil {
			return p.ToValue(*v)
		}
	case *json.Number:
		if v != nil {
			return p.ToValue(*v)
		}
	case *string:
		if v != nil {
			return p.ToValue(*v)
		}
	case *bool:
		if v != nil {
			return p.ToValue(*v)
		}
	case *time.Time:
		if v != nil {
			return p.ToValue(*v)
		}
	}
	return nil, false
}

func (*propertyNumber) ToInterface(v any) (any, bool) {
	return v, true
}

func (*propertyNumber) Validate(i any) bool {
	_, ok := i.(Number)
	return ok
}

func (*propertyNumber) Equal(v, w any) bool {
	vv := v.(Number)
	ww := w.(Number)
	return vv == ww
}

func (*propertyNumber) IsEmpty(_ any) bool {
	return false
}

func (v *Value) ValueNumber() (vv Number, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(Number)
	return
}

func (m *Multiple) ValuesNumber() (vv []Number, ok bool) {
	if m == nil {
		return
	}
	vv = lo.FilterMap(m.v, func(v *Value, _ int) (Number, bool) {
		return v.ValueNumber()
	})
	if len(vv) != len(m.v) {
		return nil, false
	}
	return
}
