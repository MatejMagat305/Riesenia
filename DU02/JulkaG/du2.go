package main

import (
  "fmt"
  "strconv"
  "time"
  "math/big"
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





// ULOHA 1

// Koduje cisla tak, ze prve najprv rekurzivne odcita,
// potom zisti druhe a prenasobi ho konstantou c (v tomto
// pripade 10.000) a pri vynarani z rekurzie pripocita
// prve cislo naspat. 
// Vysledok je potom: (a%c) * (a/c).
// Funguje len ak je prve zadavane cislo mensie ako zvolena
// konstanta c.

func times(){
  input("zadaj prirodzene cislo:")
  times1()
  fmt.Printf("ich sucin je: %d", ((a%10000) * (a/10000)))
}

func times1(){
  if a > 0{
    a--
    times1()
    a++
  } else{
    input("zadaj druhe prirodzene cislo:")
    a *= 10000
  }
}





// ULOHA 2

// Po prvom inpute spusti rekurziu,
// ktora najprv vynuluje pri vnarani
// premennu "a", a potom si vyziada 
// dalsie cislo. 
// Ak zadane cislo nie je nula, rekurzivne
// zavola samu seba. 
// Ak zadane cislo bolo nulove, zacne
// sa vynarat z rekurzie a pripocitavat
// naspat, co predtym odcitala. 

func sumNInt(){
  input("zadaj prirodzene cislo:")
  sumNInt1()
  fmt.Printf("ich sucet je: %d", a)
}

func sumNInt1(){
  if a > 0 {
    a--
    sumNInt1()
    a++
  } else if a < 0{
    a++
    sumNInt1()
    a--
  } else {
    input("zadaj prirodzene cislo:")
    if a != 0{
      sumNInt1()
    }
  }
}





// ULOHA 3

// Zadanie bolo natolko podrobne, ze nasledujuce
// funckie su vysledkom jeho presneho nasledovania. :)


// ODMOCNINA

func Sqrt(x float64) float64 {
	z := 1.0
	z_prev := z
	it := 1
	//for i := 0; i < 10; i++ {
	//	z -= (z*z - x) / (2*z)
	//	println(z)
	//}
	z -= (z*z - x) / (2*z)
	for z - z_prev > 0.0000001 || z - z_prev < -0.0000001{
		z_prev = z
		z -= (z*z - x) / (2*z)
    it++
	}
	//println("pocet iteracii: ", it)
	return z
}

// VYREZY

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}
	for x := range s {
		for y := range s[x] {
			s[x][y] = uint8(x^y)
		}
	}
	return s
}





// ULOHA 4

// Pospajala som vsetky kody z cviceni a minulych rokov,
// ktore boli k dispozicii. 
// V metode comb1 zo suboru comb.go som:
//		- vymenila pouzitu metodu "sucin" za metodu 
//			"sucinDivideGranulaAux", ktora bola podla 
//			suboru factorial.go efektivnejsia;
// 		- vymenila metodu "factBig" za metodu
//			"FaktorialMulRange", ktora bola podla
//			suboru factorial.go efektivnejsia
// Vo vysledku su priemerne casy pre rozne "n" cca taketo:
// 		n = 50.000  	:	 150ms
//		n = 100.000		:	 300ms
//		n = 1.000.000	:	 23s

func comb(n int64, k int64) *big.Int {
	citatel := sucinDivideGranulaAux(n-k+1, n)
	citatel.Div(citatel, FaktorialMulRange(k))
	return citatel
}

func sucinDivideGranulaAux(a int64, b int64) *big.Int {
	if b-a < 6 {
		res := big.NewInt(1)
		for a <= b { // pascal-like while
			aa := big.NewInt(a)
			res.Mul(res, aa) // res = res*aa
			res.BitLen()
			a++
		}
		return res
	} else {
		c := (a + b) / 2
		x := sucinDivideGranulaAux(a, c)
		y := sucinDivideGranulaAux(c+1, b)
		return x.Mul(x, y)
	}
}

func FaktorialMulRange(n int64) *big.Int {
  var x = big.NewInt(0)
	x.MulRange(2, n)
	return x
}



func main() {
  //times()
  //sumNInt()
  //fmt.Println(Sqrt(1))
  //fmt.Println(Sqrt(2))
  //fmt.Println(Sqrt(3))
  //var n int64 = 50000
  var n int64 = 100000
  //var n int64 = 1000000
  start := time.Now()
  fmt.Printf("comb(%d, %d) ma %d cifier, time=%v\n", n, n/2, len(comb(n, n/2).String()), time.Since(start))
  //println(comb(n, n/2).String()[:10])
}