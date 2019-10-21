// DU2 project main.go
package main

import (
	"fmt"
	"strconv"
	"time"
)

// toto je ta globalna premenna
var a int = 0

// toto nemeniť, len použiť
func input(prompt string) {
	var i string
	fmt.Println(prompt)
	fmt.Scan(&i)
	a, _ = strconv.Atoi(i)
}

func main() {
	//scitaj_run()
	//nasobenie()

	//postup 1 divide and conquer ale na posledny zo vstupov by potreboval do 10min
	zac := time.Now()
	x := ncr(1000000, 500000)
	fmt.Println(time.Since(zac))
	fmt.Println(len(x.String()))
	/////
	///toto zbehne do 2min
	//postup pre rychlost
	zac = time.Now()
	x = ncr2(10000000, 5000000)
	fmt.Println(time.Since(zac))
	fmt.Println(len(x.String()))
}
