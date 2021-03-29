// Package twofer is for returning a two-fer string
// as defined in https://github.com/exercism/problem-specifications/issues/757
package twofer

import "fmt"

// ShareWith returns a two-fer string given an input
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
