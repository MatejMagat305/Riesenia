// umocnovanie project main.go
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
//ulozia sa vstupne hodnoty ako 5^prve * 7^druhe
func umocni2() {
	if a > 0 {
		a--
		umocni2()
		a *= 5
	} else {
		input("zadaj prirodzene cislo (exponent):")
		umocni3()
	}
}

func umocni3() {
	if a > 0 {
		a--
		umocni3()
		a *= 7
	} else {
		a = 1
	}
}

//////////////////
func delSuc() {
	if a%2 == 0 {
		a /= 2
		delSuc()
		a++
	} else {
		a = 0
	}
}

func delA1() {
	if a%3 == 0 {
		a /= 3
		delA1()
	} else {
		delA2()
	}
}

func delA2() {
	if a%5 == 0 {
		a /= 5
		delA2()
	} else {
		delSuc()
	}
}

///////////
func vydel3() {
	if a%3 == 0 {
		a /= 3
		vydel3()
	}
}

func uloz() {
	if a%2 == 0 {
		a /= 2
		a *= 3
		uloz()
	}
}

///////////
func nasobX() {
	if a%5 == 0 {
		a /= 5
		nasobY()
		nasobX()
		a *= 5
	}
}

func nasobY() {
	if a%3 == 0 {
		a /= 3
		a *= 2
		nasobY()
		a *= 3
	}
}

func del() {
	if a%7 == 0 {
		if a%2 == 0 {
			vydel3()
			uloz() //medzivysledok z dvojky do exponentu 3ky
		}
		a /= 7
		nasobX()
		del()
	} else {
		delA1()
	}
}

func umocni() {
	//prve^druhe
	input("zadaj prirodzene cislo (zaklad):")
	umocni2()
	//2^(a^b) 3^a 5^a 7^b
	//exponent 2ky je vysledok, 3ky medzivysledok
	//a 5ky a 7ky su vstupne 2 prirodzene cisla
	//do exponetu 3ky ulozim 1ku - prve nasobenie bude 1*a
	//postupne budem b-krat nasobit exponenty 3ky a 5ky
	//vysledok si budem ukladat do exponentu 2ky
	//kym nebude b==0, exponent 2ky si dam ako exponent 3ky
	//a potup nasobenia opakujem
	if a%7 == 0 {
		a *= 3
		del()
	} else if a%5 == 0 {
		//ak by bol zaklad 0 -> 0^b = 0
		a = 0
	} else {
		//ak by bol exponent 0 -> a^0 = 1
		a = 1
	}
}

func main() {
	umocni()
	fmt.Printf("ich mocnina je: %d \n", a)
}
