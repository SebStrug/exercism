package pangram

import "strings"

// IsPangram checks if a case-insensitive input contains all alphabetic letters.
func IsPangram(s string) bool {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	sLower := strings.ToLower(s)
	for _, l := range letters {
		if !strings.Contains(sLower, l) {
			return false
		}
	}
	return true
}
