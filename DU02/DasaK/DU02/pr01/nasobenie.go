// du02 project nasobenie.go
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
//prve zadane cislo si ukladam ako exponent 3ky
func sucinA() {
	if a > 0 {
		a--
		sucinA()
		a *= 3
	} else {
		input("zadaj prirodzene cislo:")
		sucinB()
	}
}

//druhe zadane cislo si ukladam ako exponent 5ky
func sucinB() {
	if a > 0 {
		a--
		sucinB()
		a *= 5
	} else {
		a = 1
	}
}

/////////////////
//ziskavam vysledok sucinu
func delSuc() {
	if a%2 == 0 {
		a /= 2
		delSuc()
		a++
	} else {
		a = 0
	}
}

/////////////////

func delB() {
	if a%3 == 0 {
		a /= 3
		a *= 2
		delB()
		a *= 3
	}
}

//cisla su vynasobene, aby som sa dostala k vysledku
//nasobenia musim odstranit exponent 3ky
func delA() {
	if a%3 == 0 {
		a /= 3
		delA()
	} else {
		delSuc()
	}
}

//cislo v premenej "a" delim 5kou
//pri kazdou vydeleni 5kou vynasobim premenu
//"a" 2kou tolkokrat kolko sa da delit 3kou (func delB())
//-> v exponete 2ky bude sucin vstupnych cisiel
func del() {
	if a%5 == 0 {
		a /= 5
		delB()
		del()
	} else {
		delA()
	}
}

/////

func sucin() {
	//2^(a*b) 3^a 5^b
	input("zadaj prirodzene cislo:")
	sucinA()
	del()
}

func main() {
	sucin()
	fmt.Printf("ich sucin je: %d\n", a)
}
