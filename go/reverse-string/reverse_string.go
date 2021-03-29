package reverse

import "strings"

// Reverse reverses a string
func Reverse(s string) string {
	stringLen := len(s)
	reversed := make([]string, stringLen)
	for ind, char := range s {
		reversed[stringLen - ind - 1] = string(char)
	}
	return strings.Join(reversed, "")
}