package schema

import "github.com/reearth/reearth-cms/server/pkg/value"

type FieldTextArea struct {
	s *FieldString
}

func NewTextArea(maxLength *int) *FieldTextArea {
	return &FieldTextArea{
		s: NewString(value.TypeTextArea, maxLength),
	}
}

func (f *FieldTextArea) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:        f.Type(),
		textArea: f,
	}
}

func (f *FieldTextArea) MaxLength() *int {
	return f.s.MaxLength()
}

func (f *FieldTextArea) Type() value.Type {
	return f.s.Type()
}

func (f *FieldTextArea) Clone() *FieldTextArea {
	if f == nil {
		return nil
	}
	return &FieldTextArea{
		s: f.s.Clone(),
	}
}

func (f *FieldTextArea) Validate(v *value.Value) error {
	return f.s.Validate(v)
}
