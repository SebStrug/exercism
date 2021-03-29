package hamming

import "errors"

// Distance calculates the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("`a` and `b` must be the same length")
	}

	// compare runes not bytes so we don't error on unicode
	runeA := []rune(a)
	runeB := []rune(b)

	dist := 0
	for ind, char := range runeA {
		if char != runeB[ind] {
			dist++
		}
	}
	return dist, nil
}
