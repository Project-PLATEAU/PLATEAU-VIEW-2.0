package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_Match(t *testing.T) {
	var res any
	(&Value{t: TypeText, v: "aaa"}).Match(Match{Text: func(v string) { res = v }})
	assert.Equal(t, "aaa", res)

	res = nil
	(&Value{t: TypeBool, v: true}).Match(Match{Text: func(v string) { res = v }})
	assert.Nil(t, res)

	res = nil
	(&Value{t: TypeBool}).Match(Match{Default: func() { res = "default" }})
	assert.Equal(t, "default", res)
}

func TestOptional_Match(t *testing.T) {
	var res any
	(&Optional{t: TypeText, v: &Value{t: TypeText, v: "aaa"}}).Match(OptionalMatch{Match: Match{Text: func(v string) { res = v }}})
	assert.Equal(t, "aaa", res)

	res = nil
	(&Optional{t: TypeBool}).Match(OptionalMatch{None: func() { res = "none" }})
	assert.Equal(t, "none", res)

	res = nil
	(&Optional{t: TypeBool}).Match(OptionalMatch{Match: Match{Default: func() { res = "default" }}})
	assert.Equal(t, "default", res)
}
