package tour

import (
	"fmt"
	"math/rand"
	"math"
	"strings"
)

func Sqrt1(x float64) float64 { // s fixnym poctom 10 iteracii
	z := 1.0
	for i := 0; i < 10; i++ {
		e := z*z - x
		z -= e / (2 * z)
		fmt.Printf("current z: %f; error: %f \n", z, e)
	}
	return z
}

func Sqrt2(x float64) float64 { // iterujeme az dokym nie sme dost blizko
	z := 1.0
	threshold := 0.0000001
	e := z*z - x
	count := 0
	for ; math.Abs(e) > threshold; count++ {
		z -= e / (2 * z)
		e = z*z - x
	}
	fmt.Printf("total number of iterations: %d \n", count)
	return z
}

type function func(float64) float64

func Test(MySqrt function) { // otestuj algoritmus na par nahodnych vstupoch
  for i := 0; i < 10; i++ {
    input := rand.Intn(100) 
	fmt.Printf("###### TESTING WITH INPUT: %06d ######\n", input)
	fmt.Println("result:", MySqrt(float64(input)))
	fmt.Println(strings.Repeat("#", 40))
  }
}

