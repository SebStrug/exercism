package darts

import "math"

// Radius calculates the real cartesian radius given (x, y) coordinates
func Radius(x, y float64) float64 {
	return math.Sqrt((x * x) + (y * y))
}

// Score calculates the score for an (x, y) position on a dartboard
func Score(x, y float64) int {
	switch rad := Radius(x, y); {
	case rad <= 1:
		return 10
	case rad <= 5:
		return 5
	case rad <= 10:
		return 1
	default:
		return 0
	}
}
