package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	last, z, delta, diff := 0.0, 0.1, 1e-6, 1.0
	for diff > delta {
		last = z
		z -= (z*z - x) / (2*z)
		diff = last - z
		if diff < 0 {
			diff = -diff
		}
	}
	return z
}

func test(x float64) bool {
	res1, res2 := Sqrt(x), math.Sqrt(x)
	fmt.Println("x =", int(x))
	fmt.Println("My Sqrt  :", res1)
	fmt.Println("Math Sqrt:", res2)
	fmt.Println("My Sqrt == Math Sqrt:", res1 == res2)
	fmt.Println()
	return res1 == res2
}

func main() {
	matches := 0
	n := 20
	for i := 1; i <= n; i++ {
		if test(float64(i)) {
			matches += 1
		}
	}
	fmt.Println("Total matches with math library: ", matches)
}