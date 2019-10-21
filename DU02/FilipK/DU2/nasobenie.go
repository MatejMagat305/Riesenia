// nasobenie.go
package main

import (
	"fmt"
)

//nacitava druhe cislo ako exponent pri 3
func nacitaj_pom() {
	if a != 0 {
		a--
		nacitaj_pom()
		a *= 3
	} else {
		a = 1
	}
}

//nacitava prve cislo do exponenta pri 5
func nasobenie_nacitaj() {
	if a == 0 {
		input("zadaj druhe prirodzene cislo")
		nacitaj_pom()
	} else {
		a--
		nasobenie_nacitaj()
		a *= 5
	}
}

//nasobo 2 pokym sa da delit 3 pri vynarani nasobi spat aj3
func vynasob_pom() {
	if a%3 == 0 {
		a /= 3
		vynasob_pom()
		a *= 2
		a *= 3
	}
}

//kym sa da delit 5(teda exponent pri 5>0) tak pripocita
//exponent pri 3 k tomu pri 2 to tak, ze cislo nasobi 2
func vynasob() {
	for a%5 == 0 {
		a /= 5
		vynasob_pom()
	}
}

// zisti exponent pri 2 a na tuto hodnotu nastavi a
func zapis() {
	if a%2 == 0 {
		a /= 2
		zapis()
		a++
	} else {
		a = 0
	}
}
func nasobenie() {
	input("zadaj prirodzene cislo:")
	nasobenie_nacitaj()
	vynasob()
	zapis()
	fmt.Printf("ich sucin je: %d", a)
	fmt.Println()

}
