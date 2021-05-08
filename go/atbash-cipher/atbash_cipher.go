package atbash

import (
	"strings"
	"unicode"
)

// Atbash encrypts a string using the atbash cipher
// each letter in the alphabet takes its reversed position
// i.e. a -> z, b -> y, etc.
func Atbash(s string) string {
	// *This capacity should be /5 not %5
	encoded := make([]string, 0, len(s)+(len(s)%5))
	var encodedChar = ""
	var numValidChars = 0
	// *Could use bytes throughout instead of runes/strings
	for _, char := range s {
		// *This would be nicer as a switch statement
		if unicode.IsUpper(char) {
			// Atbash expects uppercase chars to be encoded as lowercase
			encodedChar = strings.ToLower(string(65 + (90 - char)))
		} else if unicode.IsLower(char) {
			encodedChar = string(97 + (122 - char))
		} else if unicode.IsNumber(char) {
			// Numbers remain as they are
			encodedChar = string(char)
		} else {
			continue
		}

		// Add a space every five chars
		// *Could just check the length here, no need for numValidChars
		if numValidChars%5 == 0 && numValidChars != 0 {
			encoded = append(encoded, " ")
		}
		numValidChars += 1

		encoded = append(encoded, encodedChar)
	}

	return strings.Join(encoded, "")
}
