package triangle

import "math"

// Kind of triangle
type Kind string

const (
    NaT = "not a triangle"
    Equ = "equilateral" 
    Iso = "isoceles" 
    Sca = "scalene" 
)

// KindFromSides determines what type of triangle four length inputs would produce
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 1) {
			return NaT
		}
	}
	if a + b < c || a + c < b || b + c < a {
		return NaT
	}
	ab := a == b
	ac := a == c
	bc := b == c
	if ab && ac || ac && bc {
		return Equ
	} else if ab || ac || bc {
		return Iso
	} else {
		return Sca
	}
}
