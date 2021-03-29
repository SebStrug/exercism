// Dummy conversations
package bob

import (
	"strings"
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	is_upper_case := strings.ToUpper(remark) == remark
	is_question := strings.HasSuffix(remark, "?")
	empty_string := len(strings.TrimSpace(remark)) == 0
	has_letter := regexp.MustCompile("[a-zA-Z]").MatchString(remark)
	switch true {
	// ends in a question mark
	case is_question && is_upper_case && has_letter:
		return "Calm down, I know what I'm doing!"
	case is_question:
		return "Sure."
	case is_upper_case && has_letter:
		return "Whoa, chill out!"
	case empty_string:
		return "Fine. Be that way!"
	}
	return "Whatever."
}
