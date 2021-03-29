package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
)

// LargestSeriesProduct calculate the largest product of contiguous
// digits given a span
func LargestSeriesProduct(digits string, span int) (int, error) {
	fmt.Printf("Testing num %v for span %v\n", digits, span)
	fmt.Printf("Len digits: %v\n", len(digits))
	if span > len(digits) || span < 0 {
		return -1, errors.New("Incorrect value for span")
	} else if digits == "" {
		return 1, nil
	}

	var prod int
	for i := 0; i < len(digits); i++ {
		newProd := 1
		for j := i; j < i+span; j++ {
			if j > len(digits)-1 {
				return prod, nil
			}

			num, err := strconv.Atoi(string(digits[j]))
			if err != nil {
				return 0, err
			}
			newProd *= num
		}
		if newProd > prod {
			prod = newProd
		}
	}
	return prod, nil
}
