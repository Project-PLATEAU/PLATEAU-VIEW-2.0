package mongogit

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDocument_MarshalBSON(t *testing.T) {
	reunmarshal := func(d *Document[any]) (res bson.M) {
		lo.Must0(bson.Unmarshal(lo.Must(bson.Marshal(d)), &res))
		return
	}

	objid := primitive.NewObjectID()
	v1, v2 := version.New(), version.New()

	assert.Equal(t, bson.M{
		"a":        "b",
		"_id":      objid,
		versionKey: primitive.Binary{Subtype: 0, Data: v1[:]},
		parentsKey: bson.A{primitive.Binary{Subtype: 0, Data: v2[:]}},
		refsKey:    bson.A{"latest"},
	}, reunmarshal(&Document[any]{
		Data: bson.M{"a": "b"},
		Meta: Meta{
			ObjectID: objid,
			Version:  v1,
			Parents:  []version.Version{v2},
			Refs:     []version.Ref{version.Latest},
		},
	}))
}

func TestDocument_UnmarshalBSON(t *testing.T) {
	type d struct {
		A string
	}

	reunmarshal := func(d any) (res Document[d]) {
		lo.Must0(bson.Unmarshal(lo.Must(bson.Marshal(d)), &res))
		return
	}

	objid := primitive.NewObjectID()
	v1, v2 := version.New(), version.New()

	assert.Equal(t, Document[d]{
		Data: d{A: "b"},
		Meta: Meta{
			ObjectID: objid,
			Version:  v1,
			Parents:  []version.Version{v2},
			Refs:     []version.Ref{version.Latest},
		},
	}, reunmarshal(bson.M{
		"a":        "b",
		"_id":      objid,
		versionKey: primitive.Binary{Subtype: 0, Data: v1[:]},
		parentsKey: bson.A{primitive.Binary{Subtype: 0, Data: v2[:]}},
		refsKey:    bson.A{"latest"},
	}))
}

func TestMeta_Apply(t *testing.T) {
	v1, v2 := version.New(), version.New()
	got, err := Meta{
		Version: v1,
		Parents: []version.Version{v2},
		Refs:    []version.Ref{version.Latest},
	}.apply(bson.M{
		"a": 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, bson.D{
		{Key: "a", Value: int32(1)},
		{Key: versionKey, Value: v1},
		{Key: parentsKey, Value: []version.Version{v2}},
		{Key: refsKey, Value: []version.Ref{version.Latest}},
	}, got)

	type A struct {
		A string
	}
	objid := primitive.NewObjectID()
	got, err = Meta{
		ObjectID: objid,
		Version:  v1,
		Parents:  []version.Version{v2},
		Refs:     []version.Ref{version.Latest},
	}.apply(A{
		A: "hoge",
	})
	assert.NoError(t, err)
	assert.Equal(t, bson.D{
		{Key: "a", Value: "hoge"},
		{Key: "_id", Value: objid},
		{Key: versionKey, Value: v1},
		{Key: parentsKey, Value: []version.Version{v2}},
		{Key: refsKey, Value: []version.Ref{version.Latest}},
	}, got)
}
