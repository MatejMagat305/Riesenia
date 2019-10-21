package main

import (
	"fmt"
	"math/big"
	"time"
)

func comb(n, k int64) *big.Int {
	res := big.NewInt(1)
	var j int64 = 1
	for i := n; i > k; i-- {
		ii := big.NewInt(i)
		jj := big.NewInt(j)
		res = res.Mul(res, ii).Div(res, jj)
		j += 1
	}
	return res
}

func testuj(n int64, k int64) {
	start := time.Now()
	res := comb(n, k)
	finalTime := time.Since(start)
	fmt.Printf("comb(%d, %d):\n", n, k)
	fmt.Printf("- computing time: %v\n", finalTime)
	str := res.String()
	length := len(str)
	fmt.Printf("- first 10 digits: %s\n", str[:10])
	fmt.Printf("- total digit count: %d\n\n", length)
	//fmt.Println("- result:", str)

}

func main() {
	numbers := []int64{50, 100, 1000, 10000, 50000, 100000, 200000, 500000, 1000000}
	for i := range numbers {
		n := numbers[i]
		k := numbers[i] / 2
		testuj(n, k)
	}
}