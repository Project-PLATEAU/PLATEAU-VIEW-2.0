package schema

import (
	"fmt"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var ErrInvalidKey = rerror.NewE(i18n.T("invalid key"))
var ErrInvalidType = rerror.NewE(i18n.T("invalid type"))

type FieldBuilder struct {
	f   *Field
	dv  *value.Multiple
	err error
}

func NewField(tp *TypeProperty) *FieldBuilder {
	return &FieldBuilder{
		f: &Field{
			typeProperty: tp,
		},
	}
}

func (b *FieldBuilder) Build() (*Field, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.f.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.f.typeProperty == nil {
		return nil, ErrInvalidType
	}
	if !b.f.key.IsValid() {
		return nil, &rerror.Error{
			Label: ErrInvalidKey,
			Err:   fmt.Errorf("%s", b.f.key.String()),
		}
	}
	if err := b.f.SetDefaultValue(b.dv); err != nil {
		return nil, err
	}
	return b.f, nil
}

func (b *FieldBuilder) Type(tp *TypeProperty) *FieldBuilder {
	b.f.typeProperty = tp
	return b
}

func (b *FieldBuilder) MustBuild() *Field {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *FieldBuilder) ID(id FieldID) *FieldBuilder {
	b.f.id = id.Clone()
	return b
}

func (b *FieldBuilder) NewID() *FieldBuilder {
	b.f.id = NewFieldID()
	return b
}

func (b *FieldBuilder) Unique(unique bool) *FieldBuilder {
	b.f.unique = unique
	return b
}

func (b *FieldBuilder) Multiple(multiple bool) *FieldBuilder {
	b.f.multiple = multiple
	return b
}

func (b *FieldBuilder) Required(required bool) *FieldBuilder {
	b.f.required = required
	return b
}

func (b *FieldBuilder) Order(o int) *FieldBuilder {
	b.f.order = o
	return b
}

func (b *FieldBuilder) Name(name string) *FieldBuilder {
	b.f.name = name
	return b
}

func (b *FieldBuilder) Description(description string) *FieldBuilder {
	b.f.description = description
	return b
}

func (b *FieldBuilder) Key(key key.Key) *FieldBuilder {
	b.f.key = key
	return b
}

func (b *FieldBuilder) RandomKey() *FieldBuilder {
	b.f.key = key.Random()
	return b
}

func (b *FieldBuilder) UpdatedAt(t time.Time) *FieldBuilder {
	b.f.updatedAt = t
	return b
}

func (b *FieldBuilder) DefaultValue(v *value.Multiple) *FieldBuilder {
	b.dv = v
	return b
}
