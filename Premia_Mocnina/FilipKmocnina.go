// mocnina
package main

import (
	"fmt"
	"strconv"
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

//vnara sa a nacita exponent do 5^exponent
func nacitaj_pom() {
	if a != 0 {
		a--
		nacitaj_pom()
		a *= 5
	} else {
		a = 1
	}
}

//vnara sa podla zakladu ak je 0 zavola nacitat exponent
func nacitaj() {
	if a == 0 {
		input("zadaj exponent")
		nacitaj_pom()
	} else {
		a--
		nacitaj()
		a *= 3
	}
}

//pouzivam sucin umocnenych prvocisel toto umocni 7 na rovnaky exponent ako 2
func skopiruj2do7() {
	for a%7 == 0 {
		a /= 7
	}
	if a%2 == 0 {
		a /= 2
		skopiruj2do7()
		a *= 2
		a *= 7
	}
}

//k mocnine 2 pripocita mocninu 3
func pripocitaj3() {
	if a%3 == 0 {
		a /= 3
		pripocitaj3()
		a *= 3
		a *= 2
	}

}

//vynasobi mocninu pri 2 mocninou pri 3
func vynasob2_3() {
	skopiruj2do7()
	for a%2 == 0 {
		a /= 2
	}
	for a%7 == 0 {
		a /= 7
		pripocitaj3()
	}
}

//zabalene umocnenie
func umocni() {
	a *= 2
	for a%5 == 0 {
		a /= 5
		vynasob2_3()
	}
}

//zapise vysledok
func zapis() {
	if a%2 == 0 {
		a /= 2
		zapis()
		a++
	} else {
		a = 0
	}
}
func mocnina() {
	input("zadaj zaklad")
	nacitaj()
	umocni()
	zapis()
	fmt.Printf("mocnina je: %d", a)
	fmt.Println()

}
func main() {
	mocnina()
}
