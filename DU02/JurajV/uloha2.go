package main

import (
	"fmt"
	"strconv"
)

// ---------------------------- SOURCES ----------------------------

var a int = 0

func input(prompt string) {
	var i string
	fmt.Println(prompt)
	fmt.Scan(&i)
	a, _ = strconv.Atoi(i)
}

// ---------------------------- ULOHA 2 ----------------------------

func pripocitaj() {
	if a > 0 {
		a -= 1
		pripocitaj()
		a += 1
	} else if a < 0 {
		a += 1
		pripocitaj()
		a -= 1
	} else {
		input("Zadaj prirodzene cislo:")
		if a != 0 {
			pripocitaj()
		}
	}
}

func scitajCisla() {
	input("Zadaj prirodzene cislo:")
	pripocitaj()
	fmt.Printf("Ich sucet je: %d", a)
}

func main() {
	scitajCisla()
}
