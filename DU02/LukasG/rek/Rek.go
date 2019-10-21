package rek

import (
	"fmt"
	"strconv"
)

var a int = 0

func input(prompt string) {
	var i string
	fmt.Println(prompt)
	fmt.Scan(&i)
	a, _ = strconv.Atoi(i)
}

/*
###############################################################
# PR1 - SPOSOB A s kodovanim cez prvocisla
# Funguje len na male vstupy, kedze si vysledok sucinu A*B
# pamata ako 2^(A*B) a napr. pri 2^(9*8) = 2^72 int preteka :)
###############################################################
*/

func triNaA() {
	if a > 0 {
		a--
		triNaA()
		a *= 3
	} else {
		input("zadaj druhe prirodzene cislo (B):")
		patNaB()
	}
}

func patNaB() {
	if a > 0 {
		a--
		patNaB()
		a *= 5
	} else {
		a = 1
	}
}

func Mul2NatA() {
	input("zadaj prve prirodzene cislo (A):")
	triNaA() // v a budem mat zakodovane 3^A * 5^B
	mul2NatA1()
	fmt.Printf("ich sucin je: %d\n", a)
}

func mul2NatA1() {
	if a % 3 == 0 {
		a /= 3
		mul2NatA2() // zavolaj za kazde vydelenie 3jkou
		mul2NatA1()
	} else if a % 2 == 0 { // za kazde vydelenie 3jkou (A) sme B krat (pocet vydeleni 5kou) nasobili 2
		a /= 2
		mul2NatA1()
		a += 1 // takze nakonci ma cislo tvar 5^B * 2^(A * B), takze pocet moznych deleni 2 == A * B!
	} else {
		a = 0
	}
}

func mul2NatA2() {
	if a % 5 == 0 { // del 5kou dokym sa da (B krat)
		a /= 5
		mul2NatA2()
		a *= 5
		a *= 2 // pri vynarani nasob 5kou (aby sme mali stale zakodovane B) a zaroven 2kou.
	}
}

/*
###############################################################
# PR1 - SPOSOB B s rozdelenim a na dve 16 bitove cisla
# Spoliehame sa, ze int by mal byt (podla dokumentacie)
# minimalne 32 bitovy na kazdom stroji...
###############################################################
*/

func Mul2NatB() { // predpokladame, ze a je 32 bitovy, takze 16 bitov na M, 16 bitov na N
	input("zadaj prve prirodzene cislo (A):")
	mul2NatB1()
	a = ((a & 0xFFFF0000) / 0x10000) * (a & 0xFFFF)
	fmt.Printf("ich sucin je: %d\n", a)
}

func mul2NatB1() {
	if a > 0 {
		a -= 1
		mul2NatB1()
		a += 1
	} else {
		input("zadaj druhe prirodzene cislo (B):")
		a *= 0x10000 //posun B do vrchnych 16 bitov
	}
}

/*
###############################################################
# PR2
###############################################################
*/

func SumNInt() {
	sumNInt1()
	fmt.Printf("ich sucet je: %d\n", a)
}

func sumNInt1(){
	if a > 0 {
		a--
		sumNInt1()
		a++
	} else if a < 0 {
		a++
		sumNInt1()
		a--
	} else {
		input("zadaj prirodzene cislo: ")
		if a != 0 {
			sumNInt1()
		}
	}
}



