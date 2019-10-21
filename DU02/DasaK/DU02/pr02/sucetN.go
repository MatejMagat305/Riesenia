// pr02 project sucetN.go
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

////////////////////////
//rekurzivne si zapamatavam hodnotu suctu,
//pripocitavam/odcitavam na zaklade hodnoty v "a"
func sucet2() {
	if a > 0 {
		a--
		sucet2()
		a++
	} else if a < 0 {
		a++
		sucet2()
		a--
	} else {
		input("zadaj cele cislo:")
		if a != 0 {
			sucet2()
		}
	}

}

//postupne sa citaju cisla zo vstupu, az
//dovtedy kym sa nenacita 0
func sucet() {
	input("zadaj cele cislo:")
	if a != 0 {
		sucet2()
	}
}

func main() {
	sucet()
	fmt.Printf("ich sucet je: %d\n", a)
}
