package prime

import "fmt"

// SieveOfErastothenes finds the prime numbers up to a given limit
func SieveOfErastothenes(n int) []int {
	arrayN := make([]int, n-1)
	for i := range arrayN {
		arrayN[n] = i+2
	}
	fmt.Printf("Array: %v", arrayN)
	return arrayN
}

// Nth gets the nth prime number
func Nth(n int) (int, bool) {
	foo := SieveOfErastothenes(20)
	return foo[0], true
}