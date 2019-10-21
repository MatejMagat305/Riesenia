package main

import (
	"./comb"
	"./rek"
	"./tour"
)

func main ()  {
	rek.HackImport()
	tour.HackImport()
	comb.HackImport()

	//rek.Mul2NatA() // uloha1, sposob A
	//rek.Mul2NatB() // uloha1, sposob B
	//rek.SumNInt() // uloha 2
	//tour.Test(tour.Sqrt1) // uloha 3A (Loops and Functions) cast 1
	//tour.Test(tour.Sqrt2) // uloha 3A (Loops and Functions) cast 2
	//uloha 3B (Slices) pozri tour/GoTour2.go

	comb.Eval(comb.Comb1) // uloha 4, sposob A
	//comb.Eval(comb.Comb3) // uloha 4, sposob B
}

