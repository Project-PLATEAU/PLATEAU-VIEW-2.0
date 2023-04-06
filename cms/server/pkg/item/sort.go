package item

import (
	"strings"
)

type Sort struct {
	Direction Direction
	SortBy    SortType
}
type SortType string

var (
	SortTypeCreationDate     SortType = "id"
	SortTypeModificationDate SortType = "timestamp"
)

func SortTypeFrom(s string) SortType {
	ss := strings.ToLower(s)
	switch ss {
	case "creation_date":
		return SortTypeCreationDate
	case "modification_date":
		return SortTypeModificationDate
	default:
		return SortTypeModificationDate
	}
}
