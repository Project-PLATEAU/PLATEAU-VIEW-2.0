package item

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortTypeFrom(t *testing.T) {
	assert.Equal(t, SortTypeCreationDate, SortTypeFrom("creation_date"))
	assert.Equal(t, SortTypeModificationDate, SortTypeFrom("modification_date"))
	assert.Equal(t, SortTypeModificationDate, SortTypeFrom("xxx"))
}
