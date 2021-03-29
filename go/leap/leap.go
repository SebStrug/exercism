// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// Test if a given number is a leap year
func IsLeapYear(y int) bool {
	div_400 := y % 400 == 0
	div_100 := y % 100 == 0
	div_4 := y % 4 == 0
	return div_400 || (div_4 && !div_100)
}
