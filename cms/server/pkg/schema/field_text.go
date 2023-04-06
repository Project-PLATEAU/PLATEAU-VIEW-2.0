package schema

import "github.com/reearth/reearth-cms/server/pkg/value"

type FieldText struct {
	s *FieldString
}

func NewText(maxLength *int) *FieldText {
	return &FieldText{
		s: NewString(value.TypeText, maxLength),
	}
}

func (f *FieldText) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:    f.Type(),
		text: f,
	}
}

func (f *FieldText) MaxLength() *int {
	return f.s.MaxLength()
}

func (f *FieldText) Type() value.Type {
	return f.s.Type()
}

func (f *FieldText) Clone() *FieldText {
	if f == nil {
		return nil
	}
	return &FieldText{
		s: f.s.Clone(),
	}
}

func (f *FieldText) Validate(v *value.Value) error {
	return f.s.Validate(v)
}
