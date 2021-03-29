package etl

import "strings"

// Transforms performs a simple Extract-Transform-Load process
// converting a legacy scrabble score format to a new format
func Transform(input map[int][]string) output map[string]int {
	output := make(map[string]int, 26)
	for point, letters := range input {
		for _, char := range letters {
			output[strings.ToLower(char)] = point
		}
	}
	return
}