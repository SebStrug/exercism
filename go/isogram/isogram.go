package isogram

import "strings"

// IsIsogram tests if a word is an isogram
func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	seenChars := map[rune]bool{}
	for _, char := range s {
		// ignore spaces and hyphens
		if char == '-' || char == ' ' {
			continue
		}

		if seenChars[char] {
			return false
		}
		seenChars[char] = true
	}
	return true
}
