package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// return error
		return 0, ErrNegativeSqrt(x)
	}
	// initial guess
	z := 1.0
	var prev = z
	count := 0
	for {
		prev = z
		z -= (z*z - x) / (2*z)
		count ++
		fmt.Println(prev)
		fmt.Println(z)
		if (math.Abs(prev - z) < 0.0000001) {
			return z, nil
		}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
