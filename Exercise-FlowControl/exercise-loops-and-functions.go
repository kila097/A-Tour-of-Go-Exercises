package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0	// initial guess
	var prev = z
	count := 0
	for {
		prev = z
		z -= (z*z - x) / (2*z)
		count ++
		fmt.Println(z)
		// stops after the value almost stops changing
		if (math.Abs(prev - z) < 0.0000001) {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("math.Sqrt's answer: \n")
	fmt.Println(math.Sqrt(2))
	// testing
	fmt.Println(Sqrt(9))
}
