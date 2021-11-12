package anagram

import (
	"reflect"
	"strings"
)

// makeStringMap converts a string into a map of
// letters and their frequency in the string
func makeStringMap(s string) map[string]int {
	var stringMap = make(map[string]int)
	for _, char := range strings.Split(s, "") {
		if _, ok := stringMap[char]; ok {
			stringMap[char]++
		} else {
			stringMap[char] = 1
		}
	}
	return stringMap
}

func Detect(subject string, candidates []string) []string {
	subjectLower := strings.ToLower(subject)

	subjectMap := makeStringMap(subjectLower)
	anagrams := []string{}

	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}

		candidateLower := strings.ToLower(candidate)
		// words are not anagrams of themselves
		if subjectLower == candidateLower {
			continue
		}

		candidateMap := makeStringMap(candidateLower)
		if reflect.DeepEqual(candidateMap, subjectMap) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
