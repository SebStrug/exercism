package grains

import (
	"errors"
	"math/big"
)

// Square calculates the total number of grains on a square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid square")
	}
	// using `math.Pow` is faster than calculating the power ourselves
	res := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(n-1)), nil).Uint64()
	return res, nil
}

// Total calculates the total number of grains on a square
// and all the squares preceding it
func Total() uint64 {
	// sum of 2^n from 1 to n is 2^(n+1) - 1
	val := new(big.Int).Exp(big.NewInt(2), big.NewInt(64+1), nil).Uint64()
	return val - 1
}
