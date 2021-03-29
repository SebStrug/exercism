package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if a string is a valid number per the Luhn formula
// Luhn algorithm: https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(n string) bool {
	noSpaces := strings.ReplaceAll(n, " ", "")
	noSpacesLen := len(noSpaces)
	if noSpacesLen <= 1 {
		return false
	}

	// Convert string to array of integers
	ints := make([]int, noSpacesLen)
	for i := 0; i < noSpacesLen; i++ {
		val, err := strconv.Atoi(string(noSpaces[i]))
		if err != nil {
			return false
		}
		ints[i] = val
	}

	// We want to double every second number starting from the right
	var indStart int
	if noSpacesLen%2 == 0 {
		indStart = 0
	} else {
		indStart = 1
	}

	var sum int
	for i := indStart; i < noSpacesLen; i++ {
		if (i+indStart)%2 == 0 {
			ints[i] *= 2
			if ints[i] > 9 {
				ints[i] -= 9
			}
		}
		sum += ints[i]
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
