package searchindex

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAssetBase(t *testing.T) {
	u, _ := url.Parse("https://example.com/aaa/a.zip")
	assert.Equal(t, "https://example.com/aaa/a", getAssetBase(u))
}
