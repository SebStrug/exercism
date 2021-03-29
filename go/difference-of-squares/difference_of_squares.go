package diffsquares

// Sum calculate the sum of the integers from 0 to n
func Sum(n int) int {
	return (n * (n + 1)) / 2
}

// SquareOfSum calculates the square of the sum of integers from 0 to n
func SquareOfSum(n int) int {
	sum := Sum(n)
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of integers from 0 to n
func SumOfSquares(n int) int {
	return Sum(n) * ((2 * n) + 1) / 3
}

// Difference calculates the absolute difference between the sum
// of the squares and the square of the sum for integers from 0 to n
func Difference(n int) int {
	res := SumOfSquares(n) - SquareOfSum(n)
	if res < 0 {
		return -res
	}
	return res
}
