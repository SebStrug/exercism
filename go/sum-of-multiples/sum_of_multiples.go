package summultiples

func SumMultiples(limit int, multiples ...int) int {
	// Use a map to make sure we don't get duplicate divisors
	// e.g. 15 for multiples [3,5]
	divisors := make(map[int]bool)
	for i := 1; i <= limit-1; i++ {
		for _, num := range multiples {
			if num == 0 {
				continue
			}
			if i % num == 0 {
				ok := divisors[i]
				if !ok {
					divisors[i] = true
				}
			}
		}
	}
	
	sum := 0
	for div := range divisors {
		sum += div
	}
	return sum
}