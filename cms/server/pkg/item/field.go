package item

import (
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
)

type Field struct {
	field FieldID
	value *value.Multiple
}

func NewField(field FieldID, v *value.Multiple) *Field {
	if v == nil {
		return nil
	}
	return &Field{field: field, value: v}
}

func (f *Field) FieldID() schema.FieldID {
	return f.field
}

func (f *Field) Type() value.Type {
	return f.value.Type()
}

func (f *Field) Value() *value.Multiple {
	return f.value
}
