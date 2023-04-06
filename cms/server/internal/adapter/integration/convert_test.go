package integration

import (
	"testing"

	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestFromPagination(t *testing.T) {
	assert.Equal(t, &usecasex.Pagination{
		Offset: &usecasex.OffsetPagination{
			Offset: 0,
			Limit:  50,
		},
	}, fromPagination(nil, nil))

	assert.Equal(t, &usecasex.Pagination{
		Offset: &usecasex.OffsetPagination{
			Offset: 0,
			Limit:  100,
		},
	}, fromPagination(lo.ToPtr(1), lo.ToPtr(100)))

	assert.Equal(t, &usecasex.Pagination{
		Offset: &usecasex.OffsetPagination{
			Offset: 100,
			Limit:  100,
		},
	}, fromPagination(lo.ToPtr(2), lo.ToPtr(200)))
}
