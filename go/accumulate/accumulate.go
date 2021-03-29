package accumulate

// Accumulate iteratively applies a passed function to an array of strings
func Accumulate(input []string, conv func(string) string) []string {
	output := []string{}
	for _, word := range input {
		output = append(output, conv(word))
	}
	return output
}