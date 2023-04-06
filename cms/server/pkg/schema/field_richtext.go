package schema

import "github.com/reearth/reearth-cms/server/pkg/value"

type FieldRichText struct {
	s *FieldString
}

func NewRichText(maxLength *int) *FieldRichText {
	return &FieldRichText{
		s: NewString(value.TypeRichText, maxLength),
	}
}

func (f *FieldRichText) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:        f.Type(),
		richText: f,
	}
}

func (f *FieldRichText) MaxLength() *int {
	return f.s.MaxLength()
}

func (f *FieldRichText) Type() value.Type {
	return f.s.Type()
}

func (f *FieldRichText) Clone() *FieldRichText {
	if f == nil {
		return nil
	}
	return &FieldRichText{
		s: f.s.Clone(),
	}
}

func (f *FieldRichText) Validate(v *value.Value) error {
	return f.s.Validate(v)
}
