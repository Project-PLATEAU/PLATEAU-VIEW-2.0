package publicapi

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestListParamFromEchoContext(t *testing.T) {
	e := echo.New()

	p, err := listParamFromEchoContext(e.NewContext(
		httptest.NewRequest("GET", "/?start_cursor=xxx&page_size=100", nil), nil))
	assert.NoError(t, err)
	assert.Equal(t, ListParam{
		Pagination: &usecasex.Pagination{
			Cursor: &usecasex.CursorPagination{
				First: lo.ToPtr(int64(100)),
				After: usecasex.Cursor("xxx").Ref(),
			},
		},
	}, p)

	p, err = listParamFromEchoContext(e.NewContext(
		httptest.NewRequest("GET", "/?offset=101&limit=101", nil), nil))
	assert.NoError(t, err)
	assert.Equal(t, ListParam{
		Pagination: &usecasex.Pagination{
			Offset: &usecasex.OffsetPagination{
				Limit:  100,
				Offset: 101,
			},
		},
	}, p)

	p, err = listParamFromEchoContext(e.NewContext(
		httptest.NewRequest("GET", "/?page=3&perPage=100", nil), nil))
	assert.NoError(t, err)
	assert.Equal(t, ListParam{
		Pagination: &usecasex.Pagination{
			Offset: &usecasex.OffsetPagination{
				Limit:  100,
				Offset: 200,
			},
		},
	}, p)
}
