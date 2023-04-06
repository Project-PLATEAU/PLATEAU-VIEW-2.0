package schema

import (
	"strings"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

var (
	ErrFieldValues       = rerror.NewE(i18n.T("invalid values"))
	ErrFieldDefaultValue = rerror.NewE(i18n.T("invalid default values"))
)

type FieldSelect struct {
	values []string
}

func NewSelect(values []string) *FieldSelect {
	return &FieldSelect{
		values: lo.Uniq(lo.FilterMap(values, func(v string, _ int) (string, bool) {
			s := strings.TrimSpace(v)
			return s, len(s) > 0
		})),
	}
}

func (f *FieldSelect) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:       f.Type(),
		selectt: f,
	}
}

func (f *FieldSelect) Values() []string {
	return slices.Clone(f.values)
}

func (*FieldSelect) Type() value.Type {
	return value.TypeSelect
}

func (f *FieldSelect) Clone() *FieldSelect {
	if f == nil {
		return nil
	}
	return &FieldSelect{
		values: slices.Clone(f.values),
	}
}

func (f *FieldSelect) Validate(v *value.Value) (err error) {
	v.Match(value.Match{
		Select: func(a value.String) {
			if !slices.Contains(f.values, a) {
				err = ErrInvalidValue
			}
		},
		Default: func() {
			err = ErrInvalidValue
		},
	})
	return
}
