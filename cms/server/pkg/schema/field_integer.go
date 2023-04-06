package schema

import (
	"fmt"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

var (
	ErrInvalidMinMax = rerror.NewE(i18n.T("max must be larger then min"))
)

type FieldInteger struct {
	min *int64
	max *int64
}

func NewInteger(min, max *int64) (*FieldInteger, error) {
	if min != nil && max != nil && *min > *max {
		return nil, ErrInvalidMinMax
	}
	return &FieldInteger{
		min: min,
		max: max,
	}, nil
}

func (f *FieldInteger) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:       f.Type(),
		integer: f,
	}
}

func (f *FieldInteger) Min() *int64 {
	return util.CloneRef(f.min)
}

func (f *FieldInteger) Max() *int64 {
	return util.CloneRef(f.max)
}

func (f *FieldInteger) Type() value.Type {
	return value.TypeInteger
}

func (f *FieldInteger) Clone() *FieldInteger {
	if f == nil {
		return nil
	}
	return &FieldInteger{
		min: util.CloneRef(f.min),
		max: util.CloneRef(f.max),
	}
}

func (f *FieldInteger) Validate(v *value.Value) (err error) {
	v.Match(value.Match{
		Integer: func(a value.Integer) {
			if f.min != nil && a < *f.min {
				err = fmt.Errorf("value should be larger than %d", *f.min)
			}
			if f.max != nil && a > *f.max {
				err = fmt.Errorf("value should be smaller than %d", *f.max)
			}
		},
		Default: func() {
			err = ErrInvalidValue
		},
	})
	return
}
