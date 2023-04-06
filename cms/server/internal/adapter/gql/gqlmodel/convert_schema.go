package gqlmodel

import (
	"reflect"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

func ToSchema(s *schema.Schema) *Schema {
	if s == nil {
		return nil
	}

	return &Schema{
		ID:        IDFrom(s.ID()),
		ProjectID: IDFrom(s.Project()),
		Fields: lo.Map(s.Fields(), func(sf *schema.Field, _ int) *SchemaField {
			return ToSchemaField(sf)
		}),
	}
}

func ToSchemaField(sf *schema.Field) *SchemaField {
	if sf == nil {
		return nil
	}

	return &SchemaField{
		ID:           IDFrom(sf.ID()),
		Type:         ToValueType(sf.Type()),
		TypeProperty: ToSchemaFieldTypeProperty(sf.TypeProperty(), sf.DefaultValue(), sf.Multiple()),
		Key:          sf.Key().String(),
		Title:        sf.Name(),
		Order:        lo.ToPtr(sf.Order()),
		Description:  lo.ToPtr(sf.Description()),
		Multiple:     sf.Multiple(),
		Unique:       sf.Unique(),
		Required:     sf.Required(),
		CreatedAt:    sf.CreatedAt(),
		UpdatedAt:    sf.UpdatedAt(),
	}
}

func ToSchemaFieldTypeProperty(tp *schema.TypeProperty, dv *value.Multiple, multiple bool) (res SchemaFieldTypeProperty) {
	tp.Match(schema.TypePropertyMatch{
		Text: func(f *schema.FieldText) {
			res = &SchemaFieldText{
				DefaultValue: valueString(dv, multiple),
				MaxLength:    f.MaxLength(),
			}
		},
		TextArea: func(f *schema.FieldTextArea) {
			res = &SchemaFieldTextArea{
				DefaultValue: valueString(dv, multiple),
				MaxLength:    f.MaxLength(),
			}
		},
		RichText: func(f *schema.FieldRichText) {
			res = &SchemaFieldRichText{
				DefaultValue: valueString(dv, multiple),
				MaxLength:    f.MaxLength(),
			}
		},
		Markdown: func(f *schema.FieldMarkdown) {
			res = &SchemaFieldMarkdown{
				DefaultValue: valueString(dv, multiple),
				MaxLength:    f.MaxLength(),
			}
		},
		Select: func(f *schema.FieldSelect) {
			res = &SchemaFieldSelect{
				DefaultValue: valueString(dv, multiple),
				Values:       f.Values(),
			}
		},
		Asset: func(f *schema.FieldAsset) {
			var v any = nil
			if dv != nil {
				if multiple {
					v, _ = dv.ValuesAsset()
				} else {
					v, _ = dv.First().ValueAsset()
				}
			}
			res = &SchemaFieldAsset{
				DefaultValue: v,
			}
		},
		DateTime: func(f *schema.FieldDateTime) {
			var v any = nil
			if dv != nil {
				if multiple {
					v, _ = dv.ValuesDateTime()
				} else {
					v, _ = dv.First().ValueDateTime()
				}
			}
			res = &SchemaFieldDate{
				DefaultValue: v,
			}
		},
		Bool: func(f *schema.FieldBool) {
			var v any = nil
			if dv != nil {
				if multiple {
					v, _ = dv.ValuesBool()
				} else {
					v, _ = dv.First().ValueBool()
				}
			}
			res = &SchemaFieldBool{
				DefaultValue: v,
			}
		},
		Number: func(f *schema.FieldNumber) {
			var v any = nil
			if dv != nil {
				if multiple {
					v, _ = dv.ValuesNumber()
				} else {
					v, _ = dv.First().ValueNumber()
				}
			}
			res = &SchemaFieldInteger{
				DefaultValue: v,
				Min:          util.ToPtrIfNotEmpty(int(lo.FromPtr(f.Min()))),
				Max:          util.ToPtrIfNotEmpty(int(lo.FromPtr(f.Max()))),
			}
		},
		Integer: func(f *schema.FieldInteger) {
			var v any = nil
			if dv != nil {
				if multiple {
					v, _ = dv.ValuesInteger()
				} else {
					v, _ = dv.First().ValueInteger()
				}
			}
			res = &SchemaFieldInteger{
				DefaultValue: v,
				Min:          util.ToPtrIfNotEmpty(int(lo.FromPtr(f.Min()))),
				Max:          util.ToPtrIfNotEmpty(int(lo.FromPtr(f.Max()))),
			}
		},
		Reference: func(f *schema.FieldReference) {
			res = &SchemaFieldReference{
				ModelID: IDFrom(f.Model()),
			}
		},
		URL: func(f *schema.FieldURL) {
			var v any = nil
			if dv != nil {
				if multiple {
					urls, _ := dv.ValuesURL()
					v = lo.Map(urls, func(v value.URL, _ int) string { return v.String() })
				} else {
					url, _ := dv.First().ValueURL()
					v = url.String()
				}
			}
			res = &SchemaFieldURL{
				DefaultValue: v,
			}
		},
	})
	return
}

func valueString(dv *value.Multiple, multiple bool) any {
	if dv == nil {
		return nil
	}
	var v any
	if multiple {
		v, _ = dv.ValuesString()
	} else {
		v, _ = dv.First().ValueString()
	}
	return v
}

var ErrInvalidTypeProperty = rerror.NewE(i18n.T("invalid type property"))

func FromSchemaTypeProperty(tp *SchemaFieldTypePropertyInput, t SchemaFieldType, multiple bool) (tpRes *schema.TypeProperty, dv *value.Multiple, err error) {
	if tp == nil {
		return nil, nil, nil
	}
	switch t {
	case SchemaFieldTypeText:
		x := tp.Text
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeText, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeText, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewText(x.MaxLength).TypeProperty()
	case SchemaFieldTypeTextArea:
		x := tp.TextArea
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeTextArea, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeTextArea, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewTextArea(x.MaxLength).TypeProperty()
	case SchemaFieldTypeRichText:
		x := tp.RichText
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeRichText, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeRichText, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewRichText(x.MaxLength).TypeProperty()
	case SchemaFieldTypeMarkdownText:
		x := tp.MarkdownText
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeMarkdown, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeMarkdownText, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewMarkdown(x.MaxLength).TypeProperty()
	case SchemaFieldTypeAsset:
		x := tp.Asset
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeAsset, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeAsset, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewAsset().TypeProperty()
	case SchemaFieldTypeDate:
		x := tp.Date
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeDateTime, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeDate, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewDateTime().TypeProperty()
	case SchemaFieldTypeBool:
		x := tp.Bool
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeBool, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeBool, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewBool().TypeProperty()
	case SchemaFieldTypeSelect:
		x := tp.Select
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeSelect, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeSelect, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewSelect(x.Values).TypeProperty()
	case SchemaFieldTypeInteger:
		x := tp.Integer
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeInteger, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeInteger, x.DefaultValue).AsMultiple()
		}
		var min, max *int64
		if x.Min != nil {
			min = lo.ToPtr(int64(*x.Min))
		}
		if x.Max != nil {
			max = lo.ToPtr(int64(*x.Max))
		}
		tpi, err2 := schema.NewInteger(min, max)
		if err2 != nil {
			err = err2
		}
		tpRes = tpi.TypeProperty()
	case SchemaFieldTypeReference:
		x := tp.Reference
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		mId, err := ToID[id.Model](x.ModelID)
		if err != nil {
			return nil, nil, err
		}
		tpRes = schema.NewReference(mId).TypeProperty()
	case SchemaFieldTypeURL:
		x := tp.URL
		if x == nil {
			return nil, nil, ErrInvalidTypeProperty
		}
		if multiple {
			dv = value.NewMultiple(value.TypeURL, unpackArray(x.DefaultValue))
		} else {
			dv = FromValue(SchemaFieldTypeURL, x.DefaultValue).AsMultiple()
		}
		tpRes = schema.NewURL().TypeProperty()
	default:
		return nil, nil, ErrInvalidTypeProperty
	}
	return
}

// TODO: move to util
func unpackArray(s any) []any {
	if s == nil {
		return nil
	}
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}
