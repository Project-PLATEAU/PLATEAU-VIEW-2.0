package gqlmodel

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestValueType(t *testing.T) {
	tests := []struct {
		name string
		t    value.Type
		want SchemaFieldType
	}{
		{
			name: "TypeText",
			t:    value.TypeText,
			want: SchemaFieldTypeText,
		},
		{
			name: "TypeTextArea",
			t:    value.TypeTextArea,
			want: SchemaFieldTypeTextArea,
		},
		{
			name: "TypeRichText",
			t:    value.TypeRichText,
			want: SchemaFieldTypeRichText,
		},
		{
			name: "TypeMarkdown",
			t:    value.TypeMarkdown,
			want: SchemaFieldTypeMarkdownText,
		},
		{
			name: "TypeAsset",
			t:    value.TypeAsset,
			want: SchemaFieldTypeAsset,
		},
		{
			name: "TypeDate",
			t:    value.TypeDateTime,
			want: SchemaFieldTypeDate,
		},
		{
			name: "TypeBool",
			t:    value.TypeBool,
			want: SchemaFieldTypeBool,
		},
		{
			name: "TypeSelect",
			t:    value.TypeSelect,
			want: SchemaFieldTypeSelect,
		},
		{
			name: "TypeInteger",
			t:    value.TypeInteger,
			want: SchemaFieldTypeInteger,
		},
		{
			name: "TypeReference",
			t:    value.TypeReference,
			want: SchemaFieldTypeReference,
		},
		{
			name: "TypeURL",
			t:    value.TypeURL,
			want: SchemaFieldTypeURL,
		},
		{
			name: "invalid",
			t:    "some value",
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToValueType(tt.t))
			if tt.want != "" {
				assert.Equal(t, tt.t, FromValueType(tt.want))
			} else {
				assert.Equal(t, value.TypeUnknown, FromValueType(tt.want))
			}
		})
	}
}

func TestToValue(t *testing.T) {
	now := time.Now()
	aid := id.NewAssetID()
	iid := id.NewItemID()

	tests := []struct {
		name string
		v    *value.Value
		want any
	}{
		{
			name: "TypeText",
			v:    value.TypeText.Value("aaa"),
			want: "aaa",
		},
		{
			name: "TypeTextArea",
			v:    value.TypeTextArea.Value("aaa"),
			want: "aaa",
		},
		{
			name: "TypeRichText",
			v:    value.TypeRichText.Value("aaa"),
			want: "aaa",
		},
		{
			name: "TypeMarkdown",
			v:    value.TypeMarkdown.Value("aaa"),
			want: "aaa",
		},
		{
			name: "TypeAsset",
			v:    value.TypeAsset.Value(aid),
			want: aid.String(),
		},
		{
			name: "TypeDate",
			v:    value.TypeDateTime.Value(now),
			want: now.Format(time.RFC3339),
		},
		{
			name: "TypeBool",
			v:    value.TypeBool.Value(true),
			want: true,
		},
		{
			name: "TypeSelect",
			v:    value.TypeSelect.Value("aaa"),
			want: "aaa",
		},
		{
			name: "TypeInteger",
			v:    value.TypeInteger.Value(lo.ToPtr(100)),
			want: int64(100),
		},
		{
			name: "TypeReference",
			v:    value.TypeReference.Value(iid),
			want: iid.String(),
		},
		{
			name: "TypeURL",
			v:    value.TypeURL.Value("https://example.com"),
			want: "https://example.com",
		},
		{
			name: "nil",
			v:    nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToValue(tt.v.AsMultiple(), false))
		})
	}
}

func TestFromValue(t *testing.T) {
	now := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	aid := id.NewAssetID()
	iid := id.NewItemID()

	tests := []struct {
		name string
		t    SchemaFieldType
		v    any
		want *value.Value
	}{
		{
			name: "TypeText",
			t:    SchemaFieldTypeText,
			v:    "aaa",
			want: value.TypeText.Value("aaa"),
		},
		{
			name: "TypeTextArea",
			t:    SchemaFieldTypeTextArea,
			v:    "aaa",
			want: value.TypeTextArea.Value("aaa"),
		},
		{
			name: "TypeRichText",
			t:    SchemaFieldTypeRichText,
			v:    "aaa",
			want: value.TypeRichText.Value("aaa"),
		},
		{
			name: "TypeMarkdown",
			t:    SchemaFieldTypeMarkdownText,
			v:    "aaa",
			want: value.TypeMarkdown.Value("aaa"),
		},
		{
			name: "TypeAsset",
			t:    SchemaFieldTypeAsset,
			v:    aid,
			want: value.TypeAsset.Value(aid),
		},
		{
			name: "TypeDate",
			t:    SchemaFieldTypeDate,
			v:    now.Format(time.RFC3339),
			want: value.TypeDateTime.Value(now),
		},
		{
			name: "TypeBool",
			t:    SchemaFieldTypeBool,
			v:    true,
			want: value.TypeBool.Value(true),
		},
		{
			name: "TypeSelect",
			t:    SchemaFieldTypeSelect,
			v:    "aaa",
			want: value.TypeSelect.Value("aaa"),
		},
		{
			name: "TypeInteger",
			t:    SchemaFieldTypeInteger,
			v:    int64(100),
			want: value.TypeInteger.Value(lo.ToPtr(100)),
		},
		{
			name: "TypeReference",
			t:    SchemaFieldTypeReference,
			v:    iid,
			want: value.TypeReference.Value(iid),
		},
		{
			name: "TypeURL",
			t:    SchemaFieldTypeURL,
			v:    "https://example.com",
			want: value.TypeURL.Value("https://example.com"),
		},
		{
			name: "nil",
			t:    SchemaFieldTypeBool,
			v:    nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, FromValue(tt.t, tt.v))
		})
	}
}
