package value

import (
	"encoding/json"
	"math"
	"strconv"
	"time"

	"github.com/samber/lo"
)

const TypeInteger Type = "integer"

type propertyInteger struct{}

type Integer = int64

func (p *propertyInteger) ToValue(i any) (any, bool) {
	switch v := i.(type) {
	case float64:
		if !math.IsNaN(v) {
			return Integer(v), true
		}
	case float32:
		if !math.IsNaN(float64(v)) {
			return Integer(v), true
		}
	case int:
		return Integer(v), true
	case int8:
		return Integer(v), true
	case int16:
		return Integer(v), true
	case int32:
		return Integer(v), true
	case int64:
		return Integer(v), true
	case uint:
		return Integer(v), true
	case uint8:
		return Integer(v), true
	case uint16:
		return Integer(v), true
	case uint32:
		return Integer(v), true
	case uint64:
		return Integer(v), true
	case uintptr:
		return Integer(v), true
	case json.Number:
		if i, err := v.Int64(); err == nil {
			return i, true
		}
		if f, err := v.Float64(); err == nil {
			return Integer(f), true
		}
	case string:
		if vv, err := strconv.ParseInt(v, 0, 64); err == nil {
			return Integer(vv), true
		}
		if vv, err := strconv.ParseFloat(v, 64); err == nil {
			return Integer(vv), true
		}
	case bool:
		if v {
			return Integer(1), true
		} else {
			return Integer(0), true
		}
	case time.Time:
		return Integer(v.Unix()), true
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

func (*propertyInteger) ToInterface(v any) (any, bool) {
	return v, true
}

func (*propertyInteger) Validate(i any) bool {
	_, ok := i.(Integer)
	return ok
}

func (*propertyInteger) Equal(v, w any) bool {
	vv := v.(Integer)
	ww := w.(Integer)
	return vv == ww
}

func (*propertyInteger) IsEmpty(_ any) bool {
	return false
}

func (v *Value) ValueInteger() (vv Integer, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(Integer)
	return
}

func (m *Multiple) ValuesInteger() (vv []Integer, ok bool) {
	if m == nil {
		return
	}
	vv = lo.FilterMap(m.v, func(v *Value, _ int) (Integer, bool) {
		return v.ValueInteger()
	})
	if len(vv) != len(m.v) {
		return nil, false
	}
	return
}
