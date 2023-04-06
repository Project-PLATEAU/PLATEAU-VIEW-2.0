package item

import "strings"

type Direction string

var (
	AscDirection  Direction = "asc"
	DescDirection Direction = "desc"
)

func DirectionFrom(s string) Direction {
	ss := strings.ToLower(s)
	switch ss {
	case "asc":
		return AscDirection
	case "desc":
		return DescDirection
	default:
		return ""
	}
}
func DirectionFromRef(s *string) Direction {
	if s == nil {
		return ""
	}
	return DirectionFrom(*s)
}
