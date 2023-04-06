package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateFrom(t *testing.T) {
	s := StateFrom("xxx")
	assert.Equal(t, State(""), s)
	s = StateFrom("approved")
	assert.Equal(t, StateApproved, s)
	s = StateFrom("waiting")
	assert.Equal(t, StateWaiting, s)
	s = StateFrom("draft")
	assert.Equal(t, StateDraft, s)
	s = StateFrom("closed")
	assert.Equal(t, StateClosed, s)
}

func TestState_String(t *testing.T) {
	assert.Equal(t, "closed", StateClosed.String())
}
