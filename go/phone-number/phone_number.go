package phonenumber

import (
	"errors"
	"regexp"
)

// Number returns the numbers only from a valid telephone number
func Number(n string) (string, error) {
	NumericOnly := regexp.MustCompile("[^0-9]+")
	n = NumericOnly.ReplaceAllString(n, "")

	CountryCode := regexp.MustCompile(`^[0,1]+`)
	n = CountryCode.ReplaceAllString(n, "")

	if len(n) != 10 {
		return "", errors.New("Phone number must have 10 numbers")
	} else if n[3]-48 == 0 || n[3]-48 == 1 {
		return "", errors.New("Exchange code cannot begin with 0")
	}
	return n, nil
}

// AreaCode returns the first three numbers of a valid telephone number
func AreaCode(n string) (string, error) {
	num, err := Number(n)
	if err != nil {
		return "", errors.New("Invalid phone number")
	}
	return num[:3], nil
}

// Format returns a valid telephone number in format `(XXX)-NXX-XXXX`
func Format(n string) (string, error) {
	num, err := Number(n)
	if err != nil {
		return "", errors.New("Invalid phone number")
	}
	FormattedNum := "(" + num[:3] + ") " + num[3:6] + "-" + num[6:]
	return FormattedNum, nil
}
