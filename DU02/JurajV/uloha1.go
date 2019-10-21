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

// ---------------------------- ULOHA 1 ----------------------------

func zakoduj3() {
	if a > 0 {
		a -= 1
		zakoduj3()
		a *= 3
	} else {
		input("Zadaj prirodzene cislo:")
		zakoduj5()
	}
}

func zakoduj5() {
	if a > 0 {
		a -= 1
		zakoduj5()
		a *= 5
	} else {
		a = 1
	}
}

func zvysExponentDvojke() {
	a *= 2
}

func log5() {
	if a % 5 == 0 {
		a /= 5
		zvysExponentDvojke()
		log5()
		a *= 5
	}

}

func log3() {
	if a % 3 == 0 {
		log5()
		a /= 3
		log3()
	}
}

func log2() {
	if a % 2 == 0 {
		a /= 2
		log2()
		a += 1
	} else {
		a = 0
	}
}

func vynasobCisla() {
	input("Zadaj prirodzene cislo:")
	zakoduj3() // aj 5
	log3() // zakoduje vysledok do mocnin 2ky
	log2() // poskytne vysledok
	fmt.Printf("Ich sucin je: %d\n", a)
}


// ---------------------------- RUN ----------------------------

func main() {
	vynasobCisla()
}
