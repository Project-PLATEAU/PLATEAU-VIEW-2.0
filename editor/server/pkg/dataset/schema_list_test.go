package dataset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatasetSchemaMapGraphSearchByFields(t *testing.T) {
	did1 := NewSchemaID()
	did2 := NewSchemaID()
	did3 := NewSchemaID()
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	sid := NewSceneID()
	f1, _ := NewSchemaField().ID(fid1).Type(ValueTypeString).Ref(&did2).Build()
	f2, _ := NewSchemaField().ID(fid2).Type(ValueTypeString).Ref(&did3).Build()
	f3, _ := NewSchemaField().ID(fid3).Type(ValueTypeString).Build()
	d1, _ := NewSchema().ID(did1).Scene(sid).Fields([]*SchemaField{
		f1,
	}).Build()
	d2, _ := NewSchema().ID(did2).Scene(sid).Fields([]*SchemaField{
		f2,
	}).Build()
	d3, _ := NewSchema().ID(did3).Scene(sid).Fields([]*SchemaField{
		f3,
	}).Build()

	m := SchemaList{d1, d2, d3}.Map()

	res, resf := m.GraphSearchByFields(did1, fid1, fid2, fid3)
	assert.Equal(t, SchemaList{d1, d2, d3}, res)
	assert.Equal(t, f3, resf)

	res2, resf2 := m.GraphSearchByFields(did1, fid1, fid3, fid2)
	assert.Equal(t, SchemaList{d1, d2}, res2)
	assert.Nil(t, resf2)
}
