package grains

import (
	"errors"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n<1 || n > 64 {
		return 0, errors.New("Invalid square number")
	}

	return uint64(1) << uint(n-1), nil
}

func Total() uint64 {
	return ^uint64(0)
}