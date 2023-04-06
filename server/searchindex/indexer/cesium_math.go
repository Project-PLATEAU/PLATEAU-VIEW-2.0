package indexer

import (
	"math"
)

const (
	Pi               = math.Pi
	TwoPi            = 2 * math.Pi
	DegreesPerRadian = 180.0 / math.Pi
	Epsilon1         = 0.1
	Epsilon12        = 0.000000000001
	Epsilon14        = 0.00000000000001
)

func zeroToTwoPi(angle float64) float64 {
	if angle >= 0 && angle <= TwoPi {
		// Early exit if the input is already inside the range. This avoids
		// unnecessary math which could introduce floating point error.
		return angle
	}
	mod := math.Mod(angle, TwoPi)
	if math.Abs(mod) < Epsilon14 && math.Abs(angle) > Epsilon14 {
		return TwoPi
	}
	return mod
}

func negativePiToPi(angle float64) float64 {
	if angle >= -Pi && angle <= Pi {
		// Early exit if the input is already inside the range. This avoids
		// unnecessary math which could introduce floating point error.
		return angle
	}
	return zeroToTwoPi(angle+Pi) - Pi
}

func toDegrees(radians float64) float64 {
	return radians * DegreesPerRadian
}
