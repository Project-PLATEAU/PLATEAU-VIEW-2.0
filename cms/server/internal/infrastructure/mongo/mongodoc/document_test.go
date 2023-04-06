package mongodoc

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewDocument(t *testing.T) {
	u := user.New().NewID().Email("hoge@example.com").Name("John").MustBuild()
	a := asset.New().NewID().
		Project(project.NewID()).
		Size(100).
		CreatedByUser(u.ID()).
		NewUUID().
		Thread(id.NewThreadID()).
		MustBuild()

	// should success
	assetDoc, _ := NewAsset(a)
	expected := Document{Object: lo.Must(bson.Marshal(assetDoc)), Type: "asset"}
	doc, id, err := NewDocument(a)
	assert.Equal(t, expected, doc)
	assert.Equal(t, assetDoc.ID, id)
	assert.NoError(t, err)

	// should return error
	unsupportedDoc := struct {
		Hoge string
		Fuga int64
	}{
		Hoge: "hoge",
		Fuga: 0,
	}
	doc2, id2, err := NewDocument(unsupportedDoc)
	assert.Equal(t, Document{Type: "", Object: bson.Raw(nil)}, doc2)
	assert.Equal(t, ErrInvalidObject, err)
	assert.Zero(t, id2)
}

func TestModelFrom(t *testing.T) {
	u := user.New().NewID().Email("hoge@example.com").Name("John").MustBuild()
	now := time.Now().Truncate(time.Millisecond).UTC()
	a := asset.New().NewID().Project(project.NewID()).Size(100).CreatedAt(now).NewUUID().CreatedByUser(u.ID()).
		Thread(id.NewThreadID()).MustBuild()

	// should success
	doc, _, err := NewDocument(a)
	assert.NoError(t, err)
	a2, err := ModelFrom(doc)
	assert.NoError(t, err)
	assert.Equal(t, a, a2)

	// should retunr error
	doc2 := Document{Type: "hoge", Object: lo.Must(bson.Marshal(struct{}{}))}
	got, err := ModelFrom(doc2)
	assert.Equal(t, ErrInvalidDoc, err)
	assert.Nil(t, got)
}
