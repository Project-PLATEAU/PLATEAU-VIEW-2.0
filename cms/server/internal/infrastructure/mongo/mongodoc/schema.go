package mongodoc

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
)

type SchemaDocument struct {
	ID        string
	Workspace string
	Project   string
	Fields    []FieldDocument
}

type FieldDocument struct {
	ID           string
	Name         string
	Description  string
	Order        int
	Key          string
	Unique       bool
	Multiple     bool
	Required     bool
	UpdatedAt    time.Time
	DefaultValue *ValueDocument
	TypeProperty TypePropertyDocument
}

type TypePropertyDocument struct {
	Type      string
	Text      *FieldTextPropertyDocument      `bson:",omitempty"`
	TextArea  *FieldTextPropertyDocument      `bson:",omitempty"`
	RichText  *FieldTextPropertyDocument      `bson:",omitempty"`
	Markdown  *FieldTextPropertyDocument      `bson:",omitempty"`
	Select    *FieldSelectPropertyDocument    `bson:",omitempty"`
	Number    *FieldNumberPropertyDocument    `bson:",omitempty"`
	Integer   *FieldIntegerPropertyDocument   `bson:",omitempty"`
	Reference *FieldReferencePropertyDocument `bson:",omitempty"`
}

type FieldTextPropertyDocument struct {
	MaxLength *int
}
type FieldSelectPropertyDocument struct {
	Values []string
}

type FieldNumberPropertyDocument struct {
	Min *float64
	Max *float64
}

type FieldIntegerPropertyDocument struct {
	Min *int64
	Max *int64
}

type FieldReferencePropertyDocument struct {
	Model string
}

func NewSchema(s *schema.Schema) (*SchemaDocument, string) {
	sId := s.ID().String()
	fieldsDoc := util.Map(s.Fields(), func(f *schema.Field) FieldDocument {
		fd := FieldDocument{
			ID:           f.ID().String(),
			Name:         f.Name(),
			Description:  f.Description(),
			Order:        f.Order(),
			Key:          f.Key().String(),
			Unique:       f.Unique(),
			Multiple:     f.Multiple(),
			Required:     f.Required(),
			UpdatedAt:    f.UpdatedAt(),
			DefaultValue: NewMultipleValue(f.DefaultValue()),
			TypeProperty: TypePropertyDocument{
				Type: string(f.Type()),
			},
		}

		f.TypeProperty().Match(schema.TypePropertyMatch{
			Text: func(fp *schema.FieldText) {
				fd.TypeProperty.Text = &FieldTextPropertyDocument{
					MaxLength: fp.MaxLength(),
				}
			},
			TextArea: func(fp *schema.FieldTextArea) {
				fd.TypeProperty.TextArea = &FieldTextPropertyDocument{
					MaxLength: fp.MaxLength(),
				}
			},
			RichText: func(fp *schema.FieldRichText) {
				fd.TypeProperty.RichText = &FieldTextPropertyDocument{
					MaxLength: fp.MaxLength(),
				}
			},
			Markdown: func(fp *schema.FieldMarkdown) {
				fd.TypeProperty.Markdown = &FieldTextPropertyDocument{
					MaxLength: fp.MaxLength(),
				}
			},
			Asset:    func(fp *schema.FieldAsset) {},
			DateTime: func(fp *schema.FieldDateTime) {},
			Bool:     func(fp *schema.FieldBool) {},
			Select: func(fp *schema.FieldSelect) {
				fd.TypeProperty.Select = &FieldSelectPropertyDocument{
					Values: fp.Values(),
				}
			},
			Number: func(fp *schema.FieldNumber) {
				fd.TypeProperty.Number = &FieldNumberPropertyDocument{
					Min: fp.Min(),
					Max: fp.Max(),
				}
			},
			Integer: func(fp *schema.FieldInteger) {
				fd.TypeProperty.Integer = &FieldIntegerPropertyDocument{
					Min: fp.Min(),
					Max: fp.Max(),
				}
			},
			Reference: func(fp *schema.FieldReference) {
				fd.TypeProperty.Reference = &FieldReferencePropertyDocument{
					Model: fp.Model().String(),
				}
			},
			URL: func(fp *schema.FieldURL) {},
		})
		return fd
	})
	return &SchemaDocument{
		ID:        sId,
		Workspace: s.Workspace().String(),
		Project:   s.Project().String(),
		Fields:    fieldsDoc,
	}, sId
}

func (d *SchemaDocument) Model() (*schema.Schema, error) {
	sId, err := id.SchemaIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	wId, err := id.WorkspaceIDFrom(d.Workspace)
	if err != nil {
		return nil, err
	}
	pId, err := id.ProjectIDFrom(d.Project)
	if err != nil {
		return nil, err
	}

	f, err := util.TryMap(d.Fields, func(fd FieldDocument) (*schema.Field, error) {
		tpd := fd.TypeProperty
		var tp *schema.TypeProperty
		switch value.Type(tpd.Type) {
		case value.TypeText:
			tp = schema.NewText(tpd.Text.MaxLength).TypeProperty()
		case value.TypeTextArea:
			tp = schema.NewTextArea(tpd.TextArea.MaxLength).TypeProperty()
		case value.TypeRichText:
			tp = schema.NewRichText(tpd.RichText.MaxLength).TypeProperty()
		case value.TypeMarkdown:
			tp = schema.NewMarkdown(tpd.Markdown.MaxLength).TypeProperty()
		case value.TypeAsset:
			tp = schema.NewAsset().TypeProperty()
		case value.TypeDateTime:
			tp = schema.NewDateTime().TypeProperty()
		case value.TypeBool:
			tp = schema.NewBool().TypeProperty()
		case value.TypeSelect:
			tp = schema.NewSelect(tpd.Select.Values).TypeProperty()
		case value.TypeNumber:
			tpi, err := schema.NewNumber(tpd.Number.Min, tpd.Number.Max)
			if err != nil {
				return nil, err
			}
			tp = tpi.TypeProperty()
		case value.TypeInteger:
			tpi, err := schema.NewInteger(tpd.Integer.Min, tpd.Integer.Max)
			if err != nil {
				return nil, err
			}
			tp = tpi.TypeProperty()
		case value.TypeReference:
			mid, err := id.ModelIDFrom(tpd.Reference.Model)
			if err != nil {
				return nil, err
			}
			tp = schema.NewReference(mid).TypeProperty()
		case value.TypeURL:
			tp = schema.NewURL().TypeProperty()
		}

		fid, err := id.FieldIDFrom(fd.ID)
		if err != nil {
			return nil, err
		}

		return schema.NewField(tp).
			ID(fid).
			Name(fd.Name).
			Unique(fd.Unique).
			Multiple(fd.Multiple).
			Order(fd.Order).
			Required(fd.Required).
			Description(fd.Description).
			Key(key.New(fd.Key)).
			UpdatedAt(fd.UpdatedAt).
			DefaultValue(fd.DefaultValue.MultipleValue()).
			Build()
	})
	if err != nil {
		return nil, err
	}

	return schema.New().
		ID(sId).
		Workspace(wId).
		Project(pId).
		Fields(f).
		Build()
}

type SchemaConsumer = mongox.SliceFuncConsumer[*SchemaDocument, *schema.Schema]

func NewSchemaConsumer() *SchemaConsumer {
	return NewComsumer[*SchemaDocument, *schema.Schema]()
}
