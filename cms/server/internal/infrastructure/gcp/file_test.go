package gcp

import (
	"net/url"
	"path"
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestFile_GetURL(t *testing.T) {
	bucketname := "asset.cms.test"
	host := "https://localhost:8080"
	r, err := NewFile(bucketname, host, "")
	assert.NoError(t, err)

	u := newUUID()
	n := "xxx.yyy"
	a := asset.New().NewID().
		Project(id.NewProjectID()).
		CreatedByUser(id.NewUserID()).
		Size(1000).
		FileName(n).
		UUID(u).
		Thread(id.NewThreadID()).
		MustBuild()

	expected, err := url.JoinPath(host, gcsAssetBasePath, u[:2], u[2:], n)
	assert.NoError(t, err)
	actual := r.GetURL(a)
	assert.Equal(t, expected, actual)
}

func TestFile_GetFSObjectPath(t *testing.T) {
	u := newUUID()
	n := "xxx.yyy"
	assert.Equal(t, path.Join(gcsAssetBasePath, u[:2], u[2:], "xxx.yyy"), getGCSObjectPath(u, n))

	n = "ああああ.yyy"
	assert.Equal(t, path.Join(gcsAssetBasePath, u[:2], u[2:], "ああああ.yyy"), getGCSObjectPath(u, n))

	assert.Equal(t, "", getGCSObjectPath("", ""))
}

func TestFile_IsValidUUID(t *testing.T) {
	u := newUUID()
	assert.Equal(t, true, IsValidUUID(u))

	u1 := "xxxxxx"
	assert.Equal(t, false, IsValidUUID(u1))
}
