package schema

import "github.com/reearth/reearth-cms/server/pkg/value"

type FieldMarkdown struct {
	s *FieldString
}

func NewMarkdown(maxLength *int) *FieldMarkdown {
	return &FieldMarkdown{
		s: NewString(value.TypeMarkdown, maxLength),
	}
}

func (f *FieldMarkdown) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:        f.Type(),
		markdown: f,
	}
}

func (f *FieldMarkdown) MaxLength() *int {
	return f.s.MaxLength()
}

func (f *FieldMarkdown) Type() value.Type {
	return f.s.Type()
}

func (f *FieldMarkdown) Clone() *FieldMarkdown {
	if f == nil {
		return nil
	}
	return &FieldMarkdown{
		s: f.s.Clone(),
	}
}

func (f *FieldMarkdown) Validate(v *value.Value) error {
	return f.s.Validate(v)
}
