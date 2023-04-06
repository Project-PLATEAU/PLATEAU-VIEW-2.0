package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/value"
)

type FieldDateTime struct{}

func NewDateTime() *FieldDateTime {
	return &FieldDateTime{}
}

func (f *FieldDateTime) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:        f.Type(),
		dateTime: f,
	}
}

func (f *FieldDateTime) Type() value.Type {
	return value.TypeDateTime
}

func (f *FieldDateTime) Clone() *FieldDateTime {
	if f == nil {
		return nil
	}
	return &FieldDateTime{}
}

func (f *FieldDateTime) Validate(v *value.Value) (err error) {
	v.Match(value.Match{
		DateTime: func(a value.DateTime) {
			// ok
		},
		Default: func() {
			err = ErrInvalidValue
		},
	})
	return
}
