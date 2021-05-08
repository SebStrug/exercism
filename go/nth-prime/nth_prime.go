package prime

import (
	"math"
)

// RemoveFromArray removes all multiples of a number n from array a
func RemoveFromArray(a []int, n int) []int {
	newArray := []int{}
	for _, num := range a {
		if num%n != 0 {
			newArray = append(newArray, num)
		}
	}
	return newArray
}

// SieveOfErastothenes finds the prime numbers up to a given limit
func SieveOfErastothenes(n int) []int {
	primes := make([]int, n)
	primes[0] = 2

	// Make an array of numbers from 2 to n**2
	arrayN := make([]int, int(math.Pow(float64(n), 2)))
	for i := 2; i <= int(math.Pow(float64(n), 2)); i++ {
		arrayN[n] = i
	}

	ind := 1
	for {
		if ind == n {
			break
		}
		arrayN = RemoveFromArray(arrayN, arrayN[0])
		primes[ind] = arrayN[0]
		ind++
	}

	return arrayN
}

// Nth gets the nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := SieveOfErastothenes(n)
	return primes[n], true
}
