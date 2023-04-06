package fme

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	i := ID{ItemID: "item", AssetID: "asset", ProjectID: "project"}
	assert.Equal(t, i, lo.Must(ParseID(i.String("aaa"), "aaa")))
	_, err := ParseID(i.String("aaa"), "aaa2")
	assert.Same(t, ErrInvalidID, err)
}
