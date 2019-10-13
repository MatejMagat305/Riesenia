// mincove_cislo
package main

import (
	"fmt"
)

var zname map[int]int = make(map[int]int)

var sada []int = []int{50, 20, 10, 5, 2, 1}

//var sada []int = []int{200, 100, 50, 20, 10, 5, 2, 1}
//var sada []int = []int{100, 50, 25, 10, 5, 1}

//greede rozdelenie cisla na mince
//systemy su kanonicke tak bude fungovat
//dokaz na efektivnych algoritmoch :) [EAZ]
func rozdel(n int) int {
	if ret, ok := zname[n]; ok {
		return ret
	}
	pocet := 0
	akt := n
	for _, i := range sada {
		if h, ok := zname[akt]; ok {
			pocet += h
			break
		}
		for i <= akt {
			pocet++
			akt -= i
		}
	}
	zname[n] = pocet
	return pocet
}

func mincove(n int) int {
	roz := rozdel(n)
	best := roz
	for i := 0; i <= sada[0]; i++ {
		akt := rozdel(n+i) + rozdel(i)
		if akt < best {
			best = akt
		}
	}
	return best
}

func najmensie_cislo_pre_n(n int) int {
	i := 1
	for {
		if mincove(i) == n {
			return i
		}
		i++
	}
}

func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(najmensie_cislo_pre_n(i))
	}
	fmt.Println(najmensie_cislo_pre_n(100000)) //zbehne za 1min

}
