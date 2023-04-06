package schema

import (
	"fmt"
	"unicode/utf8"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/util"
)

type FieldString struct {
	t         value.Type
	maxLength *int
}

func NewString(t value.Type, maxLength *int) *FieldString {
	return &FieldString{
		t:         t,
		maxLength: maxLength,
	}
}

func (f *FieldString) MaxLength() *int {
	return util.CloneRef(f.maxLength)
}

func (f *FieldString) Type() value.Type {
	return f.t
}

func (f *FieldString) Clone() *FieldString {
	if f == nil {
		return nil
	}
	return &FieldString{
		t:         f.t,
		maxLength: util.CloneRef(f.maxLength),
	}
}

func (f *FieldString) Validate(v *value.Value) error {
	if v.Type() != f.t {
		return ErrInvalidValue
	}

	s, ok := v.ValueString()
	if !ok {
		return ErrInvalidValue
	}

	if f.maxLength != nil {
		len := utf8.RuneCountInString(s)
		if len > *f.maxLength {
			return fmt.Errorf("value has %d characters, but it sholud be shorter than %d characters", len, *f.maxLength)
		}
	}

	return nil
}
