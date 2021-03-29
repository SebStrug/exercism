package collatzconjecture

import "errors"

// CollatzConjecture determines the number of steps required to reach 1
// implementing the collatz conjecture
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Number must be positive")
	}
	steps := 0
	for {
		// if n is 1, no steps taken
		if n == 1 {
			return steps, nil
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		steps++
	}
}
