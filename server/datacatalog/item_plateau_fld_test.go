package datacatalog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortRivers(t *testing.T) {
	rivers := []river{
		{dic: &DicEntry{Name: "b", Scale: "L2（最大規模）"}},
		{dic: &DicEntry{Name: "a", Scale: "L2（最大規模）"}},
		{dic: &DicEntry{Scale: "L1（計画規模）"}},
	}
	sortRivers(rivers)
	assert.Equal(t, []river{
		{dic: &DicEntry{Scale: "L1（計画規模）"}},
		{dic: &DicEntry{Name: "b", Scale: "L2（最大規模）"}},
		{dic: &DicEntry{Name: "a", Scale: "L2（最大規模）"}},
	}, rivers)

	rivers2 := []river{
		{dic: &DicEntry{Name: "a", Scale: "想定最大規模"}},
		{dic: &DicEntry{Name: "a", Scale: "計画規模"}},
	}
	sortRivers(rivers2)
	assert.Equal(t, []river{
		{dic: &DicEntry{Name: "a", Scale: "計画規模"}},
		{dic: &DicEntry{Name: "a", Scale: "想定最大規模"}},
	}, rivers2)
}
