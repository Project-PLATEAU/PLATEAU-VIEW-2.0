package indexer

import (
	"math"
)

type Rectangle struct {
	West  float64
	South float64
	East  float64
	North float64
}

func (rectangle *Rectangle) center() *Cartographic {
	east := rectangle.East
	west := rectangle.West

	if east < west {
		east += TwoPi
	}

	longitude := negativePiToPi((west + east) * 0.5)
	latitude := (rectangle.South + rectangle.North) * 0.5

	res := Cartographic{}
	res.Longitude = longitude
	res.Latitude = latitude
	res.Height = 0.0

	return &res
}

func rectangleFromCartographicArray(cartographic []Cartographic) *Rectangle {

	west := math.MaxFloat64
	east := -math.MaxFloat64
	westOverIDL := math.MaxFloat64
	eastOverIDL := -math.MaxFloat64
	south := math.MaxFloat64
	north := -math.MaxFloat64

	for _, position := range cartographic {
		west = math.Min(west, position.Longitude)
		east = math.Max(east, position.Longitude)
		south = math.Min(south, position.Latitude)
		north = math.Max(north, position.Latitude)

		lonAdjusted := position.Longitude + TwoPi
		if position.Longitude >= 0 {
			lonAdjusted = position.Longitude
		}

		westOverIDL = math.Min(westOverIDL, lonAdjusted)
		eastOverIDL = math.Max(eastOverIDL, lonAdjusted)
	}

	if east-west > eastOverIDL-westOverIDL {
		west = westOverIDL
		east = eastOverIDL

		if east > Pi {
			east = east - TwoPi
		}
		if west > Pi {
			west = west - TwoPi
		}
	}

	res := Rectangle{}
	res.East = east
	res.West = west
	res.North = north
	res.South = south

	return &res
}
