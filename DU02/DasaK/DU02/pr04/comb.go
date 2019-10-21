//combination
package main

import (
	"fmt"
	"math/big"
	"time"
)

//vynasobi cisla od {from} po {to}
func factorial(from int, to int, ch chan *big.Int) *big.Int {
	out := big.NewInt(1)
	out.MulRange(int64(from), int64(to))
	ch <- out
	return out
}

func comb(n int) *big.Int {
	//(n n/2) = (n!/(n-k)!k!)
	//= n! / ((n/2)!)^2
	//= faktorial(n/2 + 1, n) / faktorial(1, n/2)
	a := make(chan *big.Int)
	b := make(chan *big.Int)

	go factorial(n/2+1, n, a)
	go factorial(1, n/2, b)

	aa := <-a
	bb := <-b
	return aa.Div(aa, bb)

}

func main() {
	for _, i := range []int{50000, 100000, 1000000, 10000000} {
		start := time.Now()
		a := comb(i)
		end := time.Since(start)
		fmt.Printf("kombinacne cislo pre %d je velkost: %d ", i, len(a.String()))
		fmt.Printf("a cas: %s\n", end)
	}
	// kombinacne cislo pre 50000 je velkost: 15050 a cas: 14.9958ms
	// kombinacne cislo pre 100000 je velkost: 30101 a cas: 43.9327ms
	// kombinacne cislo pre 1000000 je velkost: 301027 a cas: 3.5037226s
	// kombinacne cislo pre 10000000 je velkost: 3010297 a cas: 14m12.9399542s
}
