// scitanie_n
package main

import (
	"fmt"
)

//postup ako na cviku akurat sa dokola vola to iste
//rekurzivne sa vola scitaj ak sa zada 0 tak konci rekurzia
// pri vynarani sa pripocitava 1
func scitaj() {
	if a == 0 {
		input("zadaj cislo")
		if a == 0 {
			return
		}
	}

	if a < 0 {
		a++
		scitaj()
		a--
	} else {
		a--
		scitaj()
		a++
	}
}
func scitaj_run() {
	scitaj()
	fmt.Println(a)
}
