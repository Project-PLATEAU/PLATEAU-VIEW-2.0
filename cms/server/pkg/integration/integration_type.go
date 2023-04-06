package integration

import "strings"

type Type string

const (
	TypePublic Type = "public"

	TypePrivate Type = "private"
)

func TypeFrom(s string) Type {
	switch strings.ToLower(s) {
	case "public":
		return TypePublic
	case "private":
		return TypePrivate
	default:
		return TypePrivate
	}
}
