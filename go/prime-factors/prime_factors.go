package prime

func Factors(n int64) []int64 {
	divisors := []int64{} // I think the capacity of the array should be the sqrt of the number

	var it int64 = 2
	for {
		// If n is 1 we have all prime factors
		// If the iterator is greater than the number, we have to reset it
		if n <= 1 || it > n {
			return divisors
		}

		if n % it == 0 {
			divisors = append(divisors, it)
			n /= it
		} else {
			it++
		}

	}

}