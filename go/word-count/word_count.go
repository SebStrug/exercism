package wordcount

import (
	"regexp"
	"strings"
)

// Frequency contains the number of occurences of each word in a phrase
type Frequency map[string]int

var digitCheck = regexp.MustCompile("^[0-9]+$")

// alphanumeric characters and the quote mark
var alphanumericQuote = regexp.MustCompile("[^a-zA-Z0-9']+")

// regex for an apostrophe not use din a contraction
// e.g. `That is 'big'`, vs. `I don't believe it`
var startApostrophe = regexp.MustCompile("^'.+")
var endApostrophe = regexp.MustCompile(".+'$")

// WordCount calculates the number of occurrences of each word in a phrase
// words are numbers, simple words, or contractions of simple words
func WordCount(phrase string) Frequency {
	phrase = strings.ReplaceAll(phrase, ",", " ")
	phrase = strings.ReplaceAll(phrase, "\n", " ")
	freq := make(Frequency)

	for _, subStr := range strings.Split(phrase, " ") {
		if subStr == "" {
			continue
		}
		if startApostrophe.MatchString(subStr) || endApostrophe.MatchString(subStr) {
			subStr = strings.ReplaceAll(subStr, "'", "")
		}
		cleanSubStr := alphanumericQuote.ReplaceAllString(subStr, "")
		normSubStr := strings.ToLower(cleanSubStr)

		// If a word is a number, each number is added individually to Frequency
		if digitCheck.MatchString(normSubStr) {
			for _, num := range normSubStr {
				freq[string(num)]++
			}
		} else {
			freq[normSubStr]++
		}
	}
	return freq
}
