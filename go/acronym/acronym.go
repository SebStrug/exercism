// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
"strings")

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	r, _ := regexp.Compile("[ |-][_]?[A-Za-z]")
	acronym := s[0:1] + strings.Join(r.FindAllString(s, -1), "")
	acronym = strings.ReplaceAll(acronym, " ", "")
	acronym = strings.ReplaceAll(acronym, "-", "")
	acronym = strings.ReplaceAll(acronym, "_", "")
	return strings.ToUpper(acronym)
}
