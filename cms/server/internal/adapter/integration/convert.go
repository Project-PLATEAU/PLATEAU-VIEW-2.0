package integration

import (
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearthx/usecasex"
)

const maxPerPage = 100
const defaultPerPage int64 = 50

func fromPagination(page, perPage *integrationapi.PageParam) *usecasex.Pagination {
	p := int64(1)
	if page != nil && *page > 0 {
		p = int64(*page)
	}

	pp := defaultPerPage
	if perPage != nil {
		if ppr := *perPage; 1 <= ppr {
			if ppr > maxPerPage {
				pp = int64(maxPerPage)
			} else {
				pp = int64(ppr)
			}
		}
	}

	return usecasex.OffsetPagination{
		Offset: (p - 1) * pp,
		Limit:  pp,
	}.Wrap()
}

func fromItemFieldParam(f integrationapi.Field) interfaces.ItemFieldParam {
	var v any = f.Value
	if f.Value != nil {
		v = *f.Value
	}

	var k *key.Key
	if f.Key != nil {
		k = key.New(*f.Key).Ref()
	}

	return interfaces.ItemFieldParam{
		Field: f.Id,
		Key:   k,
		Type:  integrationapi.FromValueType(f.Type),
		Value: v,
	}
}
