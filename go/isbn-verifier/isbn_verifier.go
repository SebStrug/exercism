package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN determines if a string represents a valid ISBN number
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	lenISBN := len(isbn)
	if lenISBN != 10 {
		// Valid ISBNs have 9 digits plus a check character
		return false
	}

	sum := 0
	var num int
	for ind, val := range isbn {
		// check character is only valid if it's at the end
		if (val == 'X') && (ind == 9) {
			num = 10
		} else if val == 'X' {
			return false
		} else {
			convVal, err := strconv.Atoi(string(val))
			if err != nil {
				return false
			}
			num = convVal
		}
		sum += num * (10 - ind)
	}
	if sum%11 == 0 {
		return true
	}
	return false
}
