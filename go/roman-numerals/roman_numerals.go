package romannumerals

import (
	"errors"
	"strings"
)

// GetQuotientRemainder returns the quotient and remainder
// of a number and its divisor
func GetQuotientRemainder(num int, div int, letter string) (string, int) {
	quot, rem := num/div, num%div
	quotStr := make([]string, quot)
	for i := 0; i < quot; i++ {
		quotStr[i] = letter
	}
	return strings.Join(quotStr, ""), rem
}

// ToRomanNumeral converts an arabic integer into
// its Roman numeral equivalent
func ToRomanNumeral(d int) (string, error) {
	if d <= 0 {
		return "", errors.New("Integer must be positive")
	} else if d > 3000 {
		return "", errors.New("Integer must be less than or equal to 3000")
	}
	romanNum := []string{}

	quotStr, rem := GetQuotientRemainder(d, 1000, "M")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 500, "L")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 100, "C")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 50, "L")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 10, "X")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 5, "V")
	romanNum = append(romanNum, quotStr)

	quotStr, rem = GetQuotientRemainder(rem, 1, "I")
	romanNum = append(romanNum, quotStr)

	return strings.Join(romanNum, ""), nil
}