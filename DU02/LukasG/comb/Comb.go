package comb

import (
	"fmt"
	"math/big"
	"time"
)

/*
###############################################################
# POMOCOU FAKTORIALU Z HODINY
###############################################################
*/

func productDivideGranAux(a int64, b int64) *big.Int {
	if b - a < 6 {
		res := big.NewInt(1)
		for ; a <= b ; a++ {
			res.Mul(res, big.NewInt(a))
		}
		return res
	} else {
		c := (a + b) / 2
		x := productDivideGranAux(a, c)
		y := productDivideGranAux(c + 1, b)
		return x.Mul(x, y)
	}
}
func factorial(b int64) *big.Int {
	return productDivideGranAux(2, b)
}

//comb(5000, 2500) ma 4994 cifier, cas vypoctu: 1.9999ms, prvych 10 cifier: 1593718685, cas formatovania: 0s
//comb(10000, 5000) ma 9994 cifier, cas vypoctu: 4.0099ms, prvych 10 cifier: 1591790263, cas formatovania: 0s
//comb(50000, 25000) ma 49992 cifier, cas vypoctu: 49.0108ms, prvych 10 cifier: 1127810378, cas formatovania: 1.0003ms
//comb(100000, 50000) ma 99992 cifier, cas vypoctu: 164.026ms, prvych 10 cifier: 2520608368, cas formatovania: 2.0004ms
//comb(500000, 250000) ma 499991 cifier, cas vypoctu: 2.9696665s, prvych 10 cifier: 1122759743, cas formatovania: 41ms
//comb(1000000, 500000) ma 999990 cifier, cas vypoctu: 10.7733795s, prvych 10 cifier: 7899578772, cas formatovania: 163.0351ms
func Comb1(n int64 , k int64) *big.Int { //pomocou faktorialu z cvicenia
	res := factorial(n)
	den := factorial(n - k)
	den.Mul(den, factorial(k))
	res.Div(res, den)
	return res
}

/*
###############################################################
# POMOCOU DYNAMICKEHO PROGRAMOVANIA (Pomale...)
###############################################################
*/

func Comb2(n int64 , k int64) *big.Int {
	last := []*big.Int {big.NewInt(1), big.NewInt(1)}
	for i := int64(2); i <= n; i++ {
		next := make([]*big.Int, i + 1)
		for j := int64(0); j <= i; j++ {
			if j == 0 || i == j {
				next[j] = big.NewInt(1)
			} else {
				next[j] = big.NewInt(0)
				next[j].Add(next[j], last[j])
				next[j].Add(next[j], last[j - 1])
			}
		}
		last = next
	}
	return last[k]
}

func Comb3(n int64 , k int64) *big.Int {
	if k > n/2 { return big.NewInt(0) } // budujeme iba lavu stranu trojuholnika
	last := []*big.Int {big.NewInt(1), big.NewInt(1)}
	for i := int64(2); i <= n; i++ {
		half := i / 2
		next := make([]*big.Int, half + 1)
		if i % 2 == 0 {
			half--
		}
		for j := int64(0); j <= half; j++ {
			if j == 0 {
				next[j] = big.NewInt(1)
			} else {
				next[j] = big.NewInt(0)
				next[j].Add(next[j], last[j])
				next[j].Add(next[j], last[j - 1])
			}
		}
		if i % 2 == 0 {
			next[half + 1] = big.NewInt(0)
			next[half + 1].Add(next[half + 1], last[half])
			next[half + 1].Add(next[half + 1], last[half])
		}
		last = next
	}
	return last[k]
}

/*
###############################################################
# FUNKCIA NA EVALUACIU / VYPIS
###############################################################
*/

func Eval(comb func(n int64 , k int64) *big.Int) {
	values := []int64{5000, 10000, 50000, 100000, 500000, 1000000}
	for _, n := range values {
		start := time.Now()
		res := comb(n, n/2)
		fmt.Printf("comb(%4d, %4d) ma %d cifier, cas vypoctu: %v", n, n/2, res.BitLen(), time.Since(start))
		start = time.Now()
		str := res.String()
		if len(str) > 10 {
			fmt.Printf(", prvych 10 cifier: %s, cas formatovania: %v", str[0:10], time.Since(start))
		}
		fmt.Println()
	}
}
