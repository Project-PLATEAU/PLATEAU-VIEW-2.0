package integrationapi

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type ItemModelSchema struct {
	Item   Item   `json:"item"`
	Model  Model  `json:"model"`
	Schema Schema `json:"schema"`
}

func NewItemModelSchema(i item.ItemModelSchema, assets *AssetContext) ItemModelSchema {
	return ItemModelSchema{
		Item:   NewItem(i.Item, i.Schema, assets),
		Model:  NewModel(i.Model, time.Time{}),
		Schema: NewSchema(i.Schema),
	}
}

func NewModel(m *model.Model, lastModified time.Time) Model {
	return Model{
		Id:           m.ID().Ref(),
		Key:          util.ToPtrIfNotEmpty(m.Key().String()),
		Name:         util.ToPtrIfNotEmpty(m.Name()),
		Description:  util.ToPtrIfNotEmpty(m.Description()),
		Public:       util.ToPtrIfNotEmpty(m.Public()),
		ProjectId:    m.Project().Ref(),
		SchemaId:     m.Schema().Ref(),
		CreatedAt:    lo.ToPtr(m.ID().Timestamp()),
		UpdatedAt:    lo.ToPtr(m.UpdatedAt()),
		LastModified: util.ToPtrIfNotEmpty(lastModified),
	}
}

func NewSchema(i *schema.Schema) Schema {
	fs := lo.Map(i.Fields(), func(f *schema.Field, _ int) SchemaField {
		return SchemaField{
			Id:       f.ID().Ref(),
			Type:     lo.ToPtr(ValueType(f.Type())),
			Key:      lo.ToPtr(f.Key().String()),
			Required: lo.ToPtr(f.Required()),
		}
	})

	return Schema{
		Id:        i.ID().Ref(),
		ProjectId: i.Project().Ref(),
		Fields:    &fs,
		CreatedAt: lo.ToPtr(i.ID().Timestamp()),
	}
}
