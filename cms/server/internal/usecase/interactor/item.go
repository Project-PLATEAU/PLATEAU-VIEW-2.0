package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Item struct {
	repos       *repo.Container
	gateways    *gateway.Container
	ignoreEvent bool
}

func NewItem(r *repo.Container, g *gateway.Container) *Item {
	return &Item{
		repos:    r,
		gateways: g,
	}
}

func (i Item) FindByID(ctx context.Context, itemID id.ItemID, _ *usecase.Operator) (item.Versioned, error) {
	return i.repos.Item.FindByID(ctx, itemID, nil)
}

func (i Item) FindPublicByID(ctx context.Context, itemID id.ItemID, _ *usecase.Operator) (item.Versioned, error) {
	return i.repos.Item.FindByID(ctx, itemID, version.Public.Ref())
}

func (i Item) FindByIDs(ctx context.Context, ids id.ItemIDList, _ *usecase.Operator) (item.VersionedList, error) {
	return i.repos.Item.FindByIDs(ctx, ids, nil)
}

func (i Item) ItemStatus(ctx context.Context, itemsIds id.ItemIDList, _ *usecase.Operator) (map[id.ItemID]item.Status, error) {
	requests, err := i.repos.Request.FindByItems(ctx, itemsIds)
	if err != nil {
		return nil, err
	}
	items, err := i.repos.Item.FindAllVersionsByIDs(ctx, itemsIds)
	if err != nil {
		return nil, err
	}
	res := map[id.ItemID]item.Status{}
	for _, itemId := range itemsIds {
		s := item.StatusDraft
		latest, _ := lo.Find(items, func(v item.Versioned) bool {
			return v.Value().ID() == itemId && v.Refs().Has(version.Latest)
		})
		hasPublicVersion := lo.ContainsBy(items, func(v item.Versioned) bool {
			return v.Value().ID() == itemId && v.Refs().Has(version.Public)
		})
		if hasPublicVersion {
			s = s.Wrap(item.StatusPublic)
		}
		hasApprovedRequest, hasWaitingRequest := false, false
		for _, r := range requests {
			if !r.Items().IDs().Has(itemId) {
				continue
			}
			switch r.State() {
			case request.StateApproved:
				hasApprovedRequest = true
			case request.StateWaiting:
				hasWaitingRequest = true
			}
			if hasApprovedRequest && hasWaitingRequest {
				break
			}
		}

		if hasPublicVersion && !latest.Refs().Has(version.Public) {
			s = s.Wrap(item.StatusChanged)
		}
		if hasWaitingRequest {
			s = s.Wrap(item.StatusReview)
		}
		res[itemId] = s
	}
	return res, nil
}

func (i Item) FindByProject(ctx context.Context, projectID id.ProjectID, p *usecasex.Pagination, operator *usecase.Operator) (item.VersionedList, *usecasex.PageInfo, error) {
	if !operator.IsReadableProject(projectID) {
		return nil, nil, rerror.ErrNotFound
	}
	// TODO: check operation for projects that publication type is limited
	return i.repos.Item.FindByProject(ctx, projectID, nil, p)
}

func (i Item) FindByModel(ctx context.Context, modelID id.ModelID, p *usecasex.Pagination, operator *usecase.Operator) (item.VersionedList, *usecasex.PageInfo, error) {
	m, err := i.repos.Model.FindByID(ctx, modelID)
	if err != nil {
		return nil, nil, err
	}
	if !operator.IsReadableProject(m.Project()) {
		return nil, nil, rerror.ErrNotFound
	}
	return i.repos.Item.FindByModel(ctx, m.ID(), nil, p)
}

func (i Item) FindPublicByModel(ctx context.Context, modelID id.ModelID, p *usecasex.Pagination, _ *usecase.Operator) (item.VersionedList, *usecasex.PageInfo, error) {
	m, err := i.repos.Model.FindByID(ctx, modelID)
	if err != nil {
		return nil, nil, err
	}
	// TODO: check operation for projects that publication type is limited
	return i.repos.Item.FindByModel(ctx, m.ID(), version.Public.Ref(), p)
}

func (i Item) FindBySchema(ctx context.Context, schemaID id.SchemaID, sort *usecasex.Sort, p *usecasex.Pagination, _ *usecase.Operator) (item.VersionedList, *usecasex.PageInfo, error) {
	s, err := i.repos.Schema.FindByID(ctx, schemaID)
	if err != nil {
		return nil, nil, err
	}

	sfIds := s.Fields().IDs()
	res, page, err := i.repos.Item.FindBySchema(ctx, schemaID, nil, sort, p)
	return res.FilterFields(sfIds), page, err
}

func (i Item) FindByAssets(ctx context.Context, list id.AssetIDList, _ *usecase.Operator) (map[id.AssetID]item.VersionedList, error) {
	itms, err := i.repos.Item.FindByAssets(ctx, list, nil)
	if err != nil {
		return nil, err
	}
	res := map[id.AssetID]item.VersionedList{}
	for _, aid := range list {
		for _, itm := range itms {
			if itm.Value().AssetIDs().Has(aid) && !slices.Contains(res[aid], itm) {
				res[aid] = append(res[aid], itm)
			}
		}
	}
	return res, nil
}

func (i Item) FindAllVersionsByID(ctx context.Context, itemID id.ItemID, _ *usecase.Operator) (item.VersionedList, error) {
	return i.repos.Item.FindAllVersionsByID(ctx, itemID)
}

func (i Item) Search(ctx context.Context, q *item.Query, sort *usecasex.Sort, p *usecasex.Pagination, _ *usecase.Operator) (item.VersionedList, *usecasex.PageInfo, error) {
	return i.repos.Item.Search(ctx, q, sort, p)
}

func (i Item) Create(ctx context.Context, param interfaces.CreateItemParam, operator *usecase.Operator) (item.Versioned, error) {
	if operator.User == nil && operator.Integration == nil {
		return nil, interfaces.ErrInvalidOperator
	}

	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (item.Versioned, error) {
		s, err := i.repos.Schema.FindByID(ctx, param.SchemaID)
		if err != nil {
			return nil, err
		}

		prj, err := i.repos.Project.FindByID(ctx, s.Project())
		if err != nil {
			return nil, err
		}

		m, err := i.repos.Model.FindByID(ctx, param.ModelID)
		if err != nil {
			return nil, err
		}

		if !operator.IsWritableWorkspace(s.Workspace()) {
			return nil, interfaces.ErrOperationDenied
		}

		fields, err := itemFieldsFromParams(param.Fields, s)
		if err != nil {
			return nil, err
		}

		if err := i.checkUnique(ctx, fields, s, m.ID(), nil); err != nil {
			return nil, err
		}

		th, err := thread.New().NewID().Workspace(s.Workspace()).Build()

		if err != nil {
			return nil, err
		}
		if err := i.repos.Thread.Save(ctx, th); err != nil {
			return nil, err
		}

		ib := item.New().
			NewID().
			Schema(s.ID()).
			Project(s.Project()).
			Model(m.ID()).
			Thread(th.ID()).
			Fields(fields)

		if operator.User != nil {
			ib = ib.User(*operator.User)
		}
		if operator.Integration != nil {
			ib = ib.Integration(*operator.Integration)
		}

		it, err := ib.Build()
		if err != nil {
			return nil, err
		}

		if err := i.repos.Item.Save(ctx, it); err != nil {
			return nil, err
		}

		vi, err := i.repos.Item.FindByID(ctx, it.ID(), nil)
		if err != nil {
			return nil, err
		}

		if err := i.event(ctx, Event{
			Project:   prj,
			Workspace: s.Workspace(),
			Type:      event.ItemCreate,
			Object:    vi,
			WebhookObject: item.ItemModelSchema{
				Item:   vi.Value(),
				Model:  m,
				Schema: s,
			},
			Operator: operator.Operator(),
		}); err != nil {
			return nil, err
		}

		return vi, nil
	})
}

func (i Item) LastModifiedByModel(ctx context.Context, model id.ModelID, op *usecase.Operator) (time.Time, error) {
	return i.repos.Item.LastModifiedByModel(ctx, model)
}

func (i Item) Update(ctx context.Context, param interfaces.UpdateItemParam, operator *usecase.Operator) (item.Versioned, error) {
	if operator.User == nil && operator.Integration == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	if len(param.Fields) == 0 {
		return nil, interfaces.ErrItemFieldRequired
	}

	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (item.Versioned, error) {
		itm, err := i.repos.Item.FindByID(ctx, param.ItemID, nil)
		if err != nil {
			return nil, err
		}

		itv := itm.Value()
		if !operator.CanUpdate(itv) {
			return nil, interfaces.ErrOperationDenied
		}

		m, err := i.repos.Model.FindByID(ctx, itv.Model())
		if err != nil {
			return nil, err
		}

		s, err := i.repos.Schema.FindByID(ctx, itv.Schema())
		if err != nil {
			return nil, err
		}

		prj, err := i.repos.Project.FindByID(ctx, s.Project())
		if err != nil {
			return nil, err
		}

		fields, err := itemFieldsFromParams(param.Fields, s)
		if err != nil {
			return nil, err
		}

		if err := i.checkUnique(ctx, fields, s, itv.Model(), itv); err != nil {
			return nil, err
		}

		itv.UpdateFields(fields)
		if err := i.repos.Item.Save(ctx, itv); err != nil {
			return nil, err
		}

		if err := i.event(ctx, Event{
			Project:   prj,
			Workspace: s.Workspace(),
			Type:      event.ItemUpdate,
			Object:    itm,
			WebhookObject: item.ItemModelSchema{
				Item:   itv,
				Model:  m,
				Schema: s,
			},
			Operator: operator.Operator(),
		}); err != nil {
			return nil, err
		}

		return itm, nil
	})
}

func (i Item) Delete(ctx context.Context, itemID id.ItemID, operator *usecase.Operator) error {
	if operator.User == nil && operator.Integration == nil {
		return interfaces.ErrInvalidOperator
	}

	return Run0(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) error {
		itm, err := i.repos.Item.FindByID(ctx, itemID, nil)
		if err != nil {
			return err
		}

		if !operator.CanUpdate(itm.Value()) {
			return interfaces.ErrOperationDenied
		}

		return i.repos.Item.Remove(ctx, itemID)
	})
}

func (i Item) Unpublish(ctx context.Context, itemIDs id.ItemIDList, operator *usecase.Operator) (item.VersionedList, error) {
	if operator.User == nil && operator.Integration == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (item.VersionedList, error) {
		items, err := i.repos.Item.FindByIDs(ctx, itemIDs, nil)
		if err != nil {
			return nil, err
		}

		// check all items were found
		if len(items) != len(itemIDs) {
			return nil, interfaces.ErrItemMissing
		}

		// check all items on the same models
		s := lo.CountBy(items, func(itm item.Versioned) bool {
			return itm.Value().Model() == items[0].Value().Model()
		})
		if s != len(items) {
			return nil, interfaces.ErrItemsShouldBeOnSameModel
		}

		m, err := i.repos.Model.FindByID(ctx, items[0].Value().Model())
		if err != nil {
			return nil, err
		}

		prj, err := i.repos.Project.FindByID(ctx, m.Project())
		if err != nil {
			return nil, err
		}

		sch, err := i.repos.Schema.FindByID(ctx, m.Schema())
		if err != nil {
			return nil, err
		}

		if !operator.IsMaintainingWorkspace(prj.Workspace()) {
			return nil, interfaces.ErrInvalidOperator
		}

		// remove public ref from the items
		for _, itm := range items {
			if err := i.repos.Item.UpdateRef(ctx, itm.Value().ID(), version.Public, nil); err != nil {
				return nil, err
			}
		}

		for _, itm := range items {
			if err := i.event(ctx, Event{
				Project:   prj,
				Workspace: prj.Workspace(),
				Type:      event.ItemUnpublish,
				Object:    itm,
				WebhookObject: item.ItemModelSchema{
					Item:   itm.Value(),
					Model:  m,
					Schema: sch,
				},
				Operator: operator.Operator(),
			}); err != nil {
				return nil, err
			}
		}

		return items, nil
	})
}

func (i Item) checkUnique(ctx context.Context, itemFields []*item.Field, s *schema.Schema, mid id.ModelID, itm *item.Item) error {
	var fieldsArg []repo.FieldAndValue
	for _, f := range itemFields {
		if itm != nil {
			oldF := itm.Field(f.FieldID())
			if oldF != nil && f.Value().Equal(oldF.Value()) {
				continue
			}
		}

		sf := s.Field(f.FieldID())
		if sf == nil {
			return interfaces.ErrInvalidField
		}

		newV := f.Value()
		if !sf.Unique() || newV.IsEmpty() {
			continue
		}

		fieldsArg = append(fieldsArg, repo.FieldAndValue{
			Field: f.FieldID(),
			Value: newV,
		})
	}

	exists, err := i.repos.Item.FindByModelAndValue(ctx, mid, fieldsArg, nil)
	if err != nil {
		return err
	}

	if len(exists) > 0 && (itm == nil || len(exists) != 1 || exists[0].Value().ID() != itm.ID()) {
		return interfaces.ErrDuplicatedItemValue
	}

	return nil
}

func itemFieldsFromParams(fields []interfaces.ItemFieldParam, s *schema.Schema) ([]*item.Field, error) {
	return util.TryMap(fields, func(f interfaces.ItemFieldParam) (*item.Field, error) {
		sf := s.FieldByIDOrKey(f.Field, f.Key)
		if sf == nil {
			return nil, interfaces.ErrFieldNotFound
		}

		if !sf.Multiple() {
			f.Value = []any{f.Value}
		}

		as, ok := f.Value.([]any)
		if !ok {
			return nil, interfaces.ErrInvalidValue
		}

		m := value.NewMultiple(sf.Type(), as)
		if err := sf.Validate(m); err != nil {
			return nil, fmt.Errorf("field %s: %w", sf.Name(), err)
		}

		return item.NewField(sf.ID(), m), nil
	})
}

func (i Item) event(ctx context.Context, e Event) error {
	if i.ignoreEvent {
		return nil
	}

	_, err := createEvent(ctx, i.repos, i.gateways, e)
	return err
}
