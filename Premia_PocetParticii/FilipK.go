// particie
package main

import (
	"fmt"
	"math/big"
	"time"
)

/* http://mathworld.wolfram.com/PartitionFunctionP.html
je tam jediny vzorec, ktory je pekny na pohlad tak som si vybral ten
*/
func post(n int) int {
	n = n + 1
	if n&1 == 0 {
		return n / 2
	} else {
		return -n / 2
	}
}
func pentagonal(n int) int {
	n = post(n)
	return n * (3*n - 1) / 2
}

func partitions(n int) *big.Int {
	build := make([]*big.Int, n+1)
	build[0] = big.NewInt(1)
	build[1] = big.NewInt(1)
	for i := 2; i <= n; i++ {
		build[i] = big.NewInt(0)
	}
	for i := 2; i <= n; i++ {
		for j := 1; i >= pentagonal(j); j++ {
			k := pentagonal(j)
			if ((j+1)/2)&1 != 0 {
				build[i].Add(build[i], build[i-k])
			} else {
				build[i].Sub(build[i], build[i-k])
			}
		}
	}
	return build[n]
}

func main() {

	i := 100000
	//1 000 000 trva 2:20
	//10 000 000 eventualne dobehne tipujem tak 14giga ram a do 3hodin
	//asi mi unikol trik ako na to alebo som zvolil zly vzorec ;p
	//tipujem, ze pri takych cislach je aj aproximacia celkom presna
	zac := time.Now()
	fmt.Println(partitions(i))
	fmt.Println(time.Since(zac))

}
