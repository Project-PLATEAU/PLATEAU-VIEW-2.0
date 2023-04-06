package schema

import (
	"fmt"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/util"
)

type FieldNumber struct {
	min *float64
	max *float64
}

func NewNumber(min, max *float64) (*FieldNumber, error) {
	if min != nil && max != nil && *min > *max {
		return nil, ErrInvalidMinMax
	}
	return &FieldNumber{
		min: min,
		max: max,
	}, nil
}

func (f *FieldNumber) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:      f.Type(),
		number: f,
	}
}

func (f *FieldNumber) Min() *float64 {
	return util.CloneRef(f.min)
}

func (f *FieldNumber) Max() *float64 {
	return util.CloneRef(f.max)
}

func (f *FieldNumber) Type() value.Type {
	return value.TypeNumber
}

func (f *FieldNumber) Clone() *FieldNumber {
	if f == nil {
		return nil
	}
	return &FieldNumber{
		min: util.CloneRef(f.min),
		max: util.CloneRef(f.max),
	}
}

func (f *FieldNumber) Validate(v *value.Value) (err error) {
	v.Match(value.Match{
		Number: func(a value.Number) {
			if f.min != nil && a < *f.min {
				err = fmt.Errorf("value should be larger than %f", *f.min)
			}
			if f.max != nil && a > *f.max {
				err = fmt.Errorf("value should be smaller than %f", *f.max)
			}
		},
		Default: func() {
			err = ErrInvalidValue
		},
	})
	return
}
