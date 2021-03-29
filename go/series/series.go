package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var strings []string
	for ind := 0; ind+n-1 < len(s); ind++ {
		strings = append(strings, s[ind:ind+n])
	}
	return strings
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	first, ok := First(n, s)
	if !ok {
		panic(nil)
	}
	return first	
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	strings := All(n, s)
	return strings[0], true
}