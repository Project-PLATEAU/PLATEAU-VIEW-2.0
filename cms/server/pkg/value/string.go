package value

import (
	"fmt"
	"strconv"
	"time"

	"github.com/samber/lo"
)

const TypeText Type = "text"
const TypeTextArea Type = "textArea"
const TypeRichText Type = "richText"
const TypeMarkdown Type = "markdown"
const TypeSelect Type = "select"

type propertyString struct{}

type String = string

func (p *propertyString) ToValue(i any) (any, bool) {
	if v, ok := i.(string); ok {
		return v, true
	} else if v, ok := i.(float64); ok {
		return strconv.FormatFloat(v, 'f', -1, 64), true
	} else if v, ok := i.(bool); ok && v {
		return "true", true
	} else if v, ok := i.(bool); ok && !v {
		return "false", true
	} else if v, ok := i.(time.Time); ok {
		return v.Format(time.RFC3339), true
	} else if v, ok := i.(*string); ok && v != nil {
		return p.ToValue(*v)
	} else if v, ok := i.(*float64); ok && v != nil {
		return p.ToValue(*v)
	} else if v, ok := i.(*bool); ok && v != nil {
		return p.ToValue(*v)
	} else if v, ok := i.(*time.Time); ok && v != nil {
		return p.ToValue(*v)
	} else if v, ok := i.(fmt.Stringer); ok && v != nil {
		return v.String(), true
	}
	return nil, false
}

func (*propertyString) ToInterface(v any) (any, bool) {
	return v, true
}

func (*propertyString) Validate(i any) bool {
	_, ok := i.(String)
	return ok
}

func (*propertyString) Equal(v, w any) bool {
	vv := v.(String)
	ww := w.(String)
	return vv == ww
}

func (*propertyString) IsEmpty(v any) bool {
	return v.(String) == ""
}

func (v *Value) ValueString() (vv String, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(String)
	return
}

func (m *Multiple) ValuesString() (vv []String, ok bool) {
	if m == nil {
		return
	}
	vv = lo.FilterMap(m.v, func(v *Value, _ int) (string, bool) {
		return v.ValueString()
	})
	if len(vv) != len(m.v) {
		return nil, false
	}
	return
}
