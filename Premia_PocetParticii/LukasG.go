package main

import (
	"log"
	"math/big"
	"time"
)

// https://en.wikipedia.org/wiki/Pentagonal_number_theorem

// Ked budem mat cas:
// http://www.math.clemson.edu/~kevja/PAPERS/ComputingPartitions-MathComp.pdf
// https://math.berkeley.edu/~mhaiman/math172-spring10/partitions.pdf

var memoryN map[int]*big.Int = make(map[int]*big.Int)

func partition(n int) *big.Int {
	value, present := memoryN[n]
	if present == true {
		return value
	}
	res := big.NewInt(0)
	var g1, g2 int
	var rek *big.Int
	for k := 1; k <= n; k++ {
		g1 = (k * (3 * k - 1) )/ 2
		g2 = ((-k) * (3 * (-k) - 1)) / 2

		if n - g1 >= 0 {
			rek = partition(n - g1)
			if k % 2 == 0 {
				res.Sub(res, rek)
			} else {
				res.Add(res, rek)
			}
		}
		if n - g2 >= 0 {
			rek = partition(n - g2)
			if k % 2 == 0 {
				res.Sub(res, rek)
			} else {
				res.Add(res, rek)
			}
		}
	}
	memoryN[n] = res
	return res
}

func main() {
	memoryN[0] = big.NewInt(1)
	start := time.Now()
	partition(100000)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}