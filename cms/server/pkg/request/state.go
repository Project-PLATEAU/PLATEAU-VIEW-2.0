package request

import "strings"

type State string

var StateApproved State = "approved"
var StateClosed State = "closed"
var StateWaiting State = "waiting"
var StateDraft State = "draft"

func (s State) String() string {
	return string(s)
}

func StateFrom(s string) State {
	ss := strings.ToLower(s)
	switch State(ss) {
	case StateDraft:
		return StateDraft
	case StateWaiting:
		return StateWaiting
	case StateApproved:
		return StateApproved
	case StateClosed:
		return StateClosed
	default:
		return State("")
	}
}
