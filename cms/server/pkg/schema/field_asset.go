package schema

import "github.com/reearth/reearth-cms/server/pkg/value"

type FieldAsset struct {
}

func NewAsset() *FieldAsset {
	return &FieldAsset{}
}

func (f *FieldAsset) TypeProperty() *TypeProperty {
	return &TypeProperty{
		t:     f.Type(),
		asset: f,
	}
}

func (f *FieldAsset) Type() value.Type {
	return value.TypeAsset
}

func (f *FieldAsset) Clone() *FieldAsset {
	if f == nil {
		return nil
	}
	return &FieldAsset{}
}

func (f *FieldAsset) Validate(v *value.Value) (err error) {
	v.Match(value.Match{
		Asset: func(a value.Asset) {
			// ok
		},
		Default: func() {
			err = ErrInvalidValue
		},
	})
	return
}
