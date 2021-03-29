package allergies

import (
	"fmt"
)

func Allergens() []string {
	return []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}
}

func Allergies(n uint) []string {
	var s []string
	if n == 0 {
		return s
	}
	bin_n := fmt.Sprintf("%b", n)
	allergens := Allergens()
	idx := 0
	// Iterate over binary number in reverse
	for i := len(bin_n)-1; i >= 0; i-- {
		if idx >= 8 {
			return s
		}
		if bin_n[i] - 48 == 1 {
			s = append(s, allergens[idx])
		}
		idx++
	}
	return s
}

func AllergicTo(n uint, s string) bool {
	allergies := Allergies(n)
	for _, allergy := range allergies {
		if s == allergy {
			return true
		}
	}
	return false
}
