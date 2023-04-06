package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var ErrInvalidValue = rerror.NewE(i18n.T("invalid value"))

// TypeProperty Represent special attributes for some field
// only one of the type properties should be not nil
type TypeProperty struct {
	t         value.Type
	asset     *FieldAsset
	text      *FieldText
	textArea  *FieldTextArea
	richText  *FieldRichText
	markdown  *FieldMarkdown
	dateTime  *FieldDateTime
	bool      *FieldBool
	selectt   *FieldSelect
	integer   *FieldInteger
	number    *FieldNumber
	reference *FieldReference
	url       *FieldURL
}

type TypePropertyMatch struct {
	Text      func(*FieldText)
	TextArea  func(*FieldTextArea)
	RichText  func(text *FieldRichText)
	Markdown  func(*FieldMarkdown)
	Asset     func(*FieldAsset)
	DateTime  func(*FieldDateTime)
	Bool      func(*FieldBool)
	Select    func(*FieldSelect)
	Integer   func(*FieldInteger)
	Number    func(*FieldNumber)
	Reference func(*FieldReference)
	URL       func(*FieldURL)
	Default   func()
}

type TypePropertyMatch1[T any] struct {
	Text      func(*FieldText) T
	TextArea  func(*FieldTextArea) T
	RichText  func(text *FieldRichText) T
	Markdown  func(*FieldMarkdown) T
	Asset     func(*FieldAsset) T
	DateTime  func(*FieldDateTime) T
	Bool      func(*FieldBool) T
	Select    func(*FieldSelect) T
	Integer   func(*FieldInteger) T
	Number    func(*FieldNumber) T
	Reference func(*FieldReference) T
	URL       func(*FieldURL) T
	Default   func() T
}

func (t *TypeProperty) Type() value.Type {
	return t.t
}

func (t *TypeProperty) Validate(v *value.Value) error {
	return MatchTypeProperty1(t, TypePropertyMatch1[error]{
		Text: func(f *FieldText) error {
			return f.Validate(v)
		},
		TextArea: func(f *FieldTextArea) error {
			return f.Validate(v)
		},
		RichText: func(f *FieldRichText) error {
			return f.Validate(v)
		},
		Markdown: func(f *FieldMarkdown) error {
			return f.Validate(v)
		},
		Asset: func(f *FieldAsset) error {
			return f.Validate(v)
		},
		Bool: func(f *FieldBool) error {
			return f.Validate(v)
		},
		DateTime: func(f *FieldDateTime) error {
			return f.Validate(v)
		},
		Number: func(f *FieldNumber) error {
			return f.Validate(v)
		},
		Integer: func(f *FieldInteger) error {
			return f.Validate(v)
		},
		Reference: func(f *FieldReference) error {
			return f.Validate(v)
		},
		Select: func(f *FieldSelect) error {
			return f.Validate(v)
		},
		URL: func(f *FieldURL) error {
			return f.Validate(v)
		},
	})
}

func (t *TypeProperty) Match(m TypePropertyMatch) {
	if t == nil || t.t == value.TypeUnknown {
		if m.Default != nil {
			m.Default()
		}
		return
	}

	switch t.t {
	case value.TypeText:
		if m.Text != nil {
			m.Text(t.text)
			return
		}
	case value.TypeTextArea:
		if m.TextArea != nil {
			m.TextArea(t.textArea)
			return
		}
	case value.TypeRichText:
		if m.RichText != nil {
			m.RichText(t.richText)
			return
		}
	case value.TypeMarkdown:
		if m.Markdown != nil {
			m.Markdown(t.markdown)
			return
		}
	case value.TypeAsset:
		if m.Asset != nil {
			m.Asset(t.asset)
			return
		}
	case value.TypeDateTime:
		if m.DateTime != nil {
			m.DateTime(t.dateTime)
			return
		}
	case value.TypeReference:
		if m.Reference != nil {
			m.Reference(t.reference)
			return
		}
	case value.TypeNumber:
		if m.Number != nil {
			m.Number(t.number)
			return
		}
	case value.TypeInteger:
		if m.Integer != nil {
			m.Integer(t.integer)
			return
		}
	case value.TypeSelect:
		if m.Select != nil {
			m.Select(t.selectt)
			return
		}
	case value.TypeBool:
		if m.Bool != nil {
			m.Bool(t.bool)
			return
		}
	case value.TypeURL:
		if m.URL != nil {
			m.URL(t.url)
			return
		}
	}

	if m.Default != nil {
		m.Default()
	}
}

func (t *TypeProperty) Clone() *TypeProperty {
	if t == nil {
		return nil
	}

	return &TypeProperty{
		t:         t.t,
		text:      t.text.Clone(),
		textArea:  t.textArea.Clone(),
		richText:  t.richText.Clone(),
		markdown:  t.markdown.Clone(),
		asset:     t.asset.Clone(),
		dateTime:  t.dateTime.Clone(),
		bool:      t.bool.Clone(),
		selectt:   t.selectt.Clone(),
		number:    t.number.Clone(),
		integer:   t.integer.Clone(),
		reference: t.reference.Clone(),
		url:       t.url.Clone(),
	}
}

func MatchTypeProperty1[T any](t *TypeProperty, m TypePropertyMatch1[T]) (res T) {
	if t == nil || t.t == value.TypeUnknown {
		if m.Default != nil {
			return m.Default()
		}
		return
	}

	switch t.t {
	case value.TypeText:
		if m.Text != nil {
			return m.Text(t.text)
		}
	case value.TypeTextArea:
		if m.TextArea != nil {
			return m.TextArea(t.textArea)
		}
	case value.TypeRichText:
		if m.RichText != nil {
			return m.RichText(t.richText)
		}
	case value.TypeMarkdown:
		if m.Markdown != nil {
			return m.Markdown(t.markdown)
		}
	case value.TypeAsset:
		if m.Asset != nil {
			return m.Asset(t.asset)
		}
	case value.TypeDateTime:
		if m.DateTime != nil {
			return m.DateTime(t.dateTime)
		}
	case value.TypeReference:
		if m.Reference != nil {
			return m.Reference(t.reference)
		}
	case value.TypeNumber:
		if m.Number != nil {
			return m.Number(t.number)
		}
	case value.TypeInteger:
		if m.Integer != nil {
			return m.Integer(t.integer)
		}
	case value.TypeSelect:
		if m.Select != nil {
			return m.Select(t.selectt)
		}
	case value.TypeBool:
		if m.Bool != nil {
			return m.Bool(t.bool)
		}
	case value.TypeURL:
		if m.URL != nil {
			return m.URL(t.url)
		}
	}

	if m.Default != nil {
		return m.Default()
	}
	return
}
