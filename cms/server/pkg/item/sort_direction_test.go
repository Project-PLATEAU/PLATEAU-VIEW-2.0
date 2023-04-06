package item

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectionFrom(t *testing.T) {
	input1 := "Desc"
	input2 := "Asc"
	want1 := DescDirection
	want2 := AscDirection
	assert.Equal(t, want1, DirectionFrom(input1))
	assert.Equal(t, want2, DirectionFrom(input2))
	assert.Equal(t, Direction(""), DirectionFrom("xxx"))
}
