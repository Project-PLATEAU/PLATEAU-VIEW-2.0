package cms

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestItem_Unmarshal(t *testing.T) {
	type str string

	type S struct {
		ID  string         `cms:"id"`
		AAA str            `cms:"aaa,"`
		BBB []string       `cms:"bbb"`
		CCC []str          `cms:"ccc"`
		DDD map[string]any `cms:"ddd"`
	}
	s := S{}

	Item{
		ID: "xxx",
		Fields: []Field{
			{Key: "aaa", Value: "bbb"},
			{Key: "bbb", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Value: []string{"a", "b"}},
			{Key: "ddd", Value: map[string]any{"a": "b"}},
		},
	}.Unmarshal(&s)
	assert.Equal(t, S{
		ID:  "xxx",
		AAA: "bbb",
		BBB: []string{"ccc", "bbb"},
		CCC: []str{"a", "b"},
		DDD: map[string]any{"a": "b"},
	}, s)

	// no panic
	Item{}.Unmarshal(nil)
	Item{}.Unmarshal((*S)(nil))
}

func TestMarshal(t *testing.T) {
	type str string

	type S struct {
		ID  string   `cms:"id"`
		AAA string   `cms:"aaa,text"`
		BBB []string `cms:"bbb,select"`
		CCC str      `cms:"ccc"`
		DDD []str    `cms:"ddd"`
		EEE string   `cms:"eee,text"`
	}
	s := S{
		ID:  "xxx",
		AAA: "bbb",
		BBB: []string{"ccc", "bbb"},
		CCC: str("x"),
		DDD: []str{"1", "2"},
	}
	i := &Item{
		ID: "xxx",
		Fields: []Field{
			{Key: "aaa", Type: "text", Value: "bbb"},
			{Key: "bbb", Type: "select", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Type: "", Value: "x"},
			{Key: "ddd", Type: "", Value: []string{"1", "2"}},
			// no field for eee
		},
	}

	item := &Item{}
	Marshal(s, item)
	assert.Equal(t, i, item)

	item2 := &Item{}
	Marshal(&s, item2)
	assert.Equal(t, i, item2)

	// no panic
	Marshal(nil, nil)
	Marshal(nil, item2)
	Marshal((*S)(nil), item2)
	Marshal(s, nil)
}

func TestItem_Field(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Value: "ccc", Type: "string",
	}, Item{
		Fields: []Field{
			{ID: "aaa", Value: "bbb", Type: "string"},
			{ID: "bbb", Value: "ccc", Type: "string"},
		},
	}.Field("bbb"))
	assert.Nil(t, Item{
		Fields: []Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}.Field("ccc"))
}

func TestItem_FieldByKey(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Key: "ccc", Type: "string",
	}, Item{
		Fields: []Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}.FieldByKey("ccc"))
	assert.Nil(t, Item{
		Fields: []Field{
			{ID: "aaa", Key: "aaa", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}.FieldByKey("bbb"))
}

func TestField_ValueString(t *testing.T) {
	assert.Equal(t, lo.ToPtr("ccc"), (&Field{
		Value: "ccc",
	}).ValueString())
	assert.Nil(t, (&Field{
		Value: 1,
	}).ValueString())
}

func TestField_ValueStrings(t *testing.T) {
	assert.Equal(t, []string{"ccc", "ddd"}, (&Field{
		Value: []string{"ccc", "ddd"},
	}).ValueStrings())
	assert.Equal(t, []string{"ccc", "ddd"}, (&Field{
		Value: []any{"ccc", "ddd", 1},
	}).ValueStrings())
	assert.Nil(t, (&Field{
		Value: "ccc",
	}).ValueStrings())
}

func TestField_ValueInt(t *testing.T) {
	assert.Equal(t, lo.ToPtr(100), (&Field{
		Value: 100,
	}).ValueInt())
	assert.Nil(t, (&Field{
		Value: "100",
	}).ValueInt())
}

func TestField_ValueJSON(t *testing.T) {
	r, err := (&Field{
		Value: `{"foo":"bar"}`,
	}).ValueJSON()
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"foo": "bar"}, r)
}

func TestItems_HasNext(t *testing.T) {
	assert.True(t, Items{Page: 1, PerPage: 50, TotalCount: 100}.HasNext())
	assert.False(t, Items{Page: 2, PerPage: 50, TotalCount: 100}.HasNext())
	assert.True(t, Items{Page: 1, PerPage: 10, TotalCount: 11}.HasNext())
	assert.False(t, Items{Page: 2, PerPage: 10, TotalCount: 11}.HasNext())
}

func TestFile_Paths(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, File{
		Path: "_",
		Children: []File{
			{Path: "a"},
			{Path: "_", Children: []File{{Path: "b"}}},
			{Path: "c"},
		},
	}.Paths())
}
