package schema

import (
	"fmt"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestField_ID(t *testing.T) {
	id := id.NewFieldID()
	assert.Equal(t, id, (&Field{id: id}).ID())
}

func TestField_TypeProperty(t *testing.T) {
	tp := &TypeProperty{}
	assert.Same(t, tp, (&Field{typeProperty: tp}).TypeProperty())
}

func TestField_Type(t *testing.T) {
	assert.Equal(t, value.TypeText, (&Field{typeProperty: &TypeProperty{t: value.TypeText}}).Type())
}

func TestField_CreatedAt(t *testing.T) {
	id := id.NewFieldID()
	assert.Equal(t, id.Timestamp(), (&Field{id: id}).CreatedAt())
}

func TestField_UpdatedAt(t *testing.T) {
	now := time.Now()
	fId := NewFieldID()
	tests := []struct {
		name  string
		field Field
		want  time.Time
	}{
		{
			name: "success",
			field: Field{
				updatedAt: now,
			},
			want: now,
		},
		{
			name: "success",
			field: Field{
				id: fId,
			},
			want: fId.Timestamp(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.field.UpdatedAt())
		})
	}
}

func TestField_Clone(t *testing.T) {
	s := &Field{
		id:           NewFieldID(),
		name:         "a",
		description:  "b",
		key:          key.Random(),
		unique:       true,
		multiple:     true,
		required:     true,
		typeProperty: NewText(nil).TypeProperty(),
		defaultValue: value.TypeText.Value("aa").AsMultiple(),
		updatedAt:    time.Now(),
	}
	c := s.Clone()
	assert.Equal(t, s, c)
	assert.NotSame(t, s, c)

	s = nil
	c = s.Clone()
	assert.Nil(t, c)
}

func TestField_SetRequired(t *testing.T) {
	f := &Field{required: false}
	f.SetRequired(true)
	assert.Equal(t, &Field{required: true}, f)
	assert.Equal(t, true, f.Required())
}

func TestField_SetUnique(t *testing.T) {
	f := &Field{unique: false}
	f.SetUnique(true)
	assert.Equal(t, &Field{unique: true}, f)
	assert.Equal(t, true, f.Unique())
}

func TestField_SetMultiple(t *testing.T) {
	f := &Field{multiple: false}
	f.SetMultiple(true)
	assert.Equal(t, &Field{multiple: true}, f)
	assert.Equal(t, true, f.Multiple())
}

func TestField_SetName(t *testing.T) {
	f := &Field{name: ""}
	f.SetName("a")
	assert.Equal(t, &Field{name: "a"}, f)
	assert.Equal(t, "a", f.Name())
}

func TestField_SetDescription(t *testing.T) {
	f := &Field{description: ""}
	f.SetDescription("a")
	assert.Equal(t, &Field{description: "a"}, f)
	assert.Equal(t, "a", f.Description())
}

func TestField_SetKey(t *testing.T) {
	f := &Field{}
	k := key.Random()
	assert.NoError(t, f.SetKey(k))
	assert.Equal(t, &Field{key: k}, f)
	assert.Equal(t, k, f.Key())

	assert.Equal(t, &rerror.Error{
		Label: ErrInvalidKey,
		Err:   fmt.Errorf("%s", ""),
	}, f.SetKey(key.New("")))
}

func TestField_SetTypeProperty(t *testing.T) {
	tp := NewText(lo.ToPtr(1)).TypeProperty()
	f := &Field{}
	assert.NoError(t, f.SetTypeProperty(tp))
	assert.Equal(t, &Field{typeProperty: tp}, f)

	f = &Field{defaultValue: value.TypeText.Value("aaa").AsMultiple()}
	assert.ErrorContains(t, f.SetTypeProperty(tp), "it sholud be shorter than 1 characters")
	assert.Equal(t, &Field{defaultValue: value.TypeText.Value("aaa").AsMultiple()}, f)

	assert.Same(t, ErrInvalidType, f.SetTypeProperty(nil))
	assert.Equal(t, &Field{defaultValue: value.TypeText.Value("aaa").AsMultiple()}, f)
}

func TestField_SetDefautValue(t *testing.T) {
	f := &Field{typeProperty: NewText(lo.ToPtr(1)).TypeProperty()}
	assert.NoError(t, f.SetDefaultValue(value.TypeText.Value("a").AsMultiple()))
	assert.Equal(t, value.TypeText.Value("a").AsMultiple(), f.defaultValue)

	assert.NoError(t, f.SetDefaultValue(nil))
	assert.Nil(t, f.defaultValue)

	assert.ErrorContains(t, f.SetDefaultValue(value.TypeText.Value("aaa").AsMultiple()), "it sholud be shorter than 1 characters")
	assert.Nil(t, f.defaultValue)
}

func TestField_Validate(t *testing.T) {
	f := &Field{typeProperty: NewText(lo.ToPtr(1)).TypeProperty()}
	assert.NoError(t, f.Validate(value.TypeText.Value("a").AsMultiple()))
	assert.NoError(t, f.Validate(nil))

	f.required = true
	assert.Same(t, ErrValueRequired, f.Validate(nil))

	assert.ErrorContains(t, f.Validate(value.TypeText.Value("aaa").AsMultiple()), "it sholud be shorter than 1 characters")
}
