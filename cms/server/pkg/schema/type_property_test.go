package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestTypeProperty_Type(t *testing.T) {
	assert.Equal(t, value.TypeText, (&TypeProperty{t: value.TypeText}).Type())
}

func TestMatchTypeProperty1(t *testing.T) {
	m := TypePropertyMatch1[string]{
		Text:      func(_ *FieldText) string { return "Text" },
		TextArea:  func(_ *FieldTextArea) string { return "TextArea" },
		RichText:  func(_ *FieldRichText) string { return "RichText" },
		Markdown:  func(_ *FieldMarkdown) string { return "Markdown" },
		Asset:     func(_ *FieldAsset) string { return "Asset" },
		DateTime:  func(_ *FieldDateTime) string { return "DateTime" },
		Bool:      func(_ *FieldBool) string { return "Bool" },
		Select:    func(_ *FieldSelect) string { return "Select" },
		Integer:   func(_ *FieldInteger) string { return "Integer" },
		Number:    func(_ *FieldNumber) string { return "Number" },
		Reference: func(_ *FieldReference) string { return "Reference" },
		URL:       func(_ *FieldURL) string { return "URL" },
		Default:   func() string { return "Default" },
	}

	type args struct {
		tp *TypeProperty
		m  TypePropertyMatch1[string]
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "text",
			args: args{
				tp: &TypeProperty{t: value.TypeText, text: &FieldText{}},
				m:  m,
			},
			want: "Text",
		},
		{
			name: "textArea",
			args: args{
				tp: &TypeProperty{t: value.TypeTextArea, textArea: &FieldTextArea{}},
				m:  m,
			},
			want: "TextArea",
		},
		{
			name: "RichText",
			args: args{
				tp: &TypeProperty{t: value.TypeRichText, richText: &FieldRichText{}},
				m:  m,
			},
			want: "RichText",
		},
		{
			name: "Markdown",
			args: args{
				tp: &TypeProperty{t: value.TypeMarkdown, markdown: &FieldMarkdown{}},
				m:  m,
			},
			want: "Markdown",
		},
		{
			name: "Asset",
			args: args{
				tp: &TypeProperty{t: value.TypeAsset, asset: &FieldAsset{}},
				m:  m,
			},
			want: "Asset",
		},
		{
			name: "DateTime",
			args: args{
				tp: &TypeProperty{t: value.TypeDateTime, dateTime: &FieldDateTime{}},
				m:  m,
			},
			want: "DateTime",
		},
		{
			name: "Bool",
			args: args{
				tp: &TypeProperty{t: value.TypeBool, bool: &FieldBool{}},
				m:  m,
			},
			want: "Bool",
		},
		{
			name: "Select",
			args: args{
				tp: &TypeProperty{t: value.TypeSelect, selectt: &FieldSelect{}},
				m:  m,
			},
			want: "Select",
		},
		{
			name: "Number",
			args: args{
				tp: &TypeProperty{t: value.TypeNumber, number: &FieldNumber{}},
				m:  m,
			},
			want: "Number",
		},
		{
			name: "Integer",
			args: args{
				tp: &TypeProperty{t: value.TypeInteger, integer: &FieldInteger{}},
				m:  m,
			},
			want: "Integer",
		},
		{
			name: "Reference",
			args: args{
				tp: &TypeProperty{t: value.TypeReference, reference: &FieldReference{}},
				m:  m,
			},
			want: "Reference",
		},
		{
			name: "URL",
			args: args{
				tp: &TypeProperty{t: value.TypeURL, url: &FieldURL{}},
				m:  m,
			},
			want: "URL",
		},
		{
			name: "Default",
			args: args{
				tp: &TypeProperty{t: value.TypeAsset, asset: &FieldAsset{}},
				m: TypePropertyMatch1[string]{
					Default: func() string { return "Default" },
				},
			},
			want: "Default",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, MatchTypeProperty1(tc.args.tp, tc.args.m))
		})
	}
}
