package raindrops

import "strconv"

// Convert a number into a string containing raindrop sounds
// according to factors. Return number if criteria not met.
// Problem specified in Exercism Go track, Raindrops exercise.
func Convert(i int) string {
	res := ""
	if i%3 == 0 {
		res += "Pling"
	}
	if i%5 == 0 {
		res += "Plang"
	}
	if i%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return strconv.Itoa(i)
	}
	return res
}
