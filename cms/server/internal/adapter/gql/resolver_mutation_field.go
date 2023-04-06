package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateField(ctx context.Context, input gqlmodel.CreateFieldInput) (*gqlmodel.FieldPayload, error) {
	mId, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	m, err := usecases(ctx).Model.FindByIDs(ctx, []id.ModelID{mId}, getOperator(ctx))
	if err != nil || len(m) != 1 {
		return nil, err
	}

	tp, dv, err := gqlmodel.FromSchemaTypeProperty(input.TypeProperty, input.Type, input.Multiple)
	if err != nil {
		return nil, err
	}

	f, err := usecases(ctx).Schema.CreateField(ctx, interfaces.CreateFieldParam{
		SchemaId:     m[0].Schema(),
		Type:         value.Type(input.Type),
		Name:         input.Title,
		Description:  input.Description,
		Key:          input.Key,
		Multiple:     input.Multiple,
		Unique:       input.Unique,
		Required:     input.Required,
		DefaultValue: dv,
		TypeProperty: tp,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.FieldPayload{
		Field: gqlmodel.ToSchemaField(f),
	}, nil
}

func (r *mutationResolver) UpdateField(ctx context.Context, input gqlmodel.UpdateFieldInput) (*gqlmodel.FieldPayload, error) {
	fId, err := gqlmodel.ToID[id.Field](input.FieldID)
	if err != nil {
		return nil, err
	}

	mId, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	ms, err := usecases(ctx).Model.FindByIDs(ctx, []id.ModelID{mId}, getOperator(ctx))
	if err != nil || len(ms) != 1 || ms[0].ID() != mId {
		if err == nil {
			return nil, rerror.NewE(i18n.T("not found"))
		}
		return nil, err
	}
	m := ms[0]

	s, err := usecases(ctx).Schema.FindByID(ctx, m.Schema(), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	dbField := s.Field(fId)

	tp, dv, err := gqlmodel.FromSchemaTypeProperty(input.TypeProperty, gqlmodel.ToValueType(dbField.Type()), dbField.Multiple())
	if err != nil {
		return nil, err
	}

	f, err := usecases(ctx).Schema.UpdateField(ctx, interfaces.UpdateFieldParam{
		SchemaId:     m.Schema(),
		FieldId:      fId,
		Name:         input.Title,
		Description:  input.Description,
		Key:          input.Key,
		Multiple:     input.Multiple,
		Order:        input.Order,
		Unique:       input.Unique,
		Required:     input.Required,
		DefaultValue: dv,
		TypeProperty: tp,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.FieldPayload{
		Field: gqlmodel.ToSchemaField(f),
	}, nil
}

func (r *mutationResolver) DeleteField(ctx context.Context, input gqlmodel.DeleteFieldInput) (*gqlmodel.DeleteFieldPayload, error) {
	fId, err := gqlmodel.ToID[id.Field](input.FieldID)
	if err != nil {
		return nil, err
	}

	mId, err := gqlmodel.ToID[id.Model](input.ModelID)
	if err != nil {
		return nil, err
	}

	m, err := usecases(ctx).Model.FindByIDs(ctx, []id.ModelID{mId}, getOperator(ctx))
	if err != nil || len(m) != 1 {
		return nil, err
	}

	if err := usecases(ctx).Schema.DeleteField(ctx, m[0].Schema(), fId, getOperator(ctx)); err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteFieldPayload{
		FieldID: input.FieldID,
	}, nil
}

func (r *mutationResolver) UpdateFields(ctx context.Context, input []*gqlmodel.UpdateFieldInput) (*gqlmodel.FieldsPayload, error) {
	mId, err := gqlmodel.ToID[id.Model](input[0].ModelID)
	if err != nil {
		return nil, err
	}

	ms, err := usecases(ctx).Model.FindByIDs(ctx, []id.ModelID{mId}, getOperator(ctx))
	if err != nil || len(ms) != 1 || ms[0].ID() != mId {
		if err == nil {
			return nil, rerror.NewE(i18n.T("not found"))
		}
		return nil, err
	}

	s, err := usecases(ctx).Schema.FindByID(ctx, ms[0].Schema(), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	params, err := util.TryMap(input, func(ipt *gqlmodel.UpdateFieldInput) (interfaces.UpdateFieldParam, error) {
		fid, err := gqlmodel.ToID[id.Field](ipt.FieldID)
		if err != nil {
			return interfaces.UpdateFieldParam{}, err
		}
		dbField := s.Field(fid)

		tp, dv, err := gqlmodel.FromSchemaTypeProperty(ipt.TypeProperty, gqlmodel.ToValueType(dbField.Type()), dbField.Multiple())
		if err != nil {
			return interfaces.UpdateFieldParam{}, err
		}
		return interfaces.UpdateFieldParam{
			SchemaId:     s.ID(),
			FieldId:      fid,
			Name:         ipt.Title,
			Description:  ipt.Description,
			Key:          ipt.Key,
			Multiple:     ipt.Multiple,
			Order:        ipt.Order,
			Unique:       ipt.Unique,
			Required:     ipt.Required,
			DefaultValue: dv,
			TypeProperty: tp,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	fl, err := usecases(ctx).Schema.UpdateFields(ctx, ms[0].Schema(), params, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.FieldsPayload{
		Fields: lo.Map(fl, func(sf *schema.Field, _ int) *gqlmodel.SchemaField {
			return gqlmodel.ToSchemaField(sf)
		}),
	}, nil
}
