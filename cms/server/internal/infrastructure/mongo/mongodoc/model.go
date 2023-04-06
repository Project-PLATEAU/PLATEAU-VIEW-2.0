package mongodoc

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearthx/mongox"
)

type ModelDocument struct {
	ID          string
	Name        string
	Description string
	Key         string
	Public      bool
	Project     string
	Schema      string
	UpdatedAt   time.Time
}

func NewModel(model *model.Model) (*ModelDocument, string) {
	mId := model.ID().String()
	return &ModelDocument{
		ID:          mId,
		Name:        model.Name(),
		Description: model.Description(),
		Key:         model.Key().String(),
		Public:      model.Public(),
		Project:     model.Project().String(),
		Schema:      model.Schema().String(),
		UpdatedAt:   model.UpdatedAt(),
	}, mId
}

func (d *ModelDocument) Model() (*model.Model, error) {
	mId, err := id.ModelIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	pId, err := id.ProjectIDFrom(d.Project)
	if err != nil {
		return nil, err
	}
	sId, err := id.SchemaIDFrom(d.Schema)
	if err != nil {
		return nil, err
	}

	return model.New().
		ID(mId).
		Name(d.Name).
		Description(d.Description).
		UpdatedAt(d.UpdatedAt).
		Key(key.New(d.Key)).
		Public(d.Public).
		Project(pId).
		Schema(sId).
		Build()
}

type ModelConsumer = mongox.SliceFuncConsumer[*ModelDocument, *model.Model]

func NewModelConsumer() *ModelConsumer {
	return NewComsumer[*ModelDocument, *model.Model]()
}
