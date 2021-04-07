package armstrong

import (
	"fmt"
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	numString := fmt.Sprint(n)
	numDigits := float64(len(numString))
	// fmt.Printf("Number: %v, Number as string: %v, number of digits: %v\n", n, numString, numDigits)

	var sum float64
	for _, digitStr := range fmt.Sprint(n) {
		digitNum, _ := strconv.ParseFloat(string(digitStr), 64)
		sum += math.Pow(float64(digitNum), numDigits)
		// fmt.Printf("Digit: %v, Sum: %v\n", digitNum, sum)
	}
	return int(sum) == n
}
