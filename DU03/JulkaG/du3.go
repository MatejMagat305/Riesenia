package main

import (
	//"golang.org/x/tour/wc"
	//"strings"
  "fmt"
)

// ULOHA 1

// Maps je tu nakopirovane aj so svojim 
// vlastnym mainom priamo z tour Go.
// Errors sa pusta v dolnom maine. :)

// Maps

/*func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, s := range strings.Fields(s){
		j, _ := m[s]
		m[s] = j + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}*/


// Errors


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return  0, ErrNegativeSqrt(x)
	}
	z := 1.0
	z_prev := z
	it := 1
	z -= (z*z - x) / (2*z)
	for z - z_prev > 0.0000001 || z - z_prev < -0.0000001{
		z_prev = z
		z -= (z*z - x) / (2*z)
    it++
	}
	return z, nil
}



// ULOHA 2

// Funckia Determinant je rekurzivna.
// Trivialny pripad je 2x2 matica, kedy
// rekurzia skonci a da vysledok podla vzorca.
// Inak sa postupne vytvori N podmatic,
// pre kazdu sa vypocita determinant, prisudi
// sa mu znamienko a vynasobi sa to prislusnym
// clenom tak, ako to bolo vo video-navode,
// ktory nam bol poskytnuty v zadani ulohy. :)
// Nakoniec sa vsetko spocita a vrati vysledok.

type Matrix struct {
   N      int
   Values [][]int
}

func (m Matrix) Print() {
	for i := range m.Values {
		for j := range m.Values[i] {
			fmt.Printf("%v ", m.Values[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Matrix)transponujMatrix() Matrix {
	r := Matrix{m.N, make([][]int, len(m.Values[0]))}
	for i := 0; i < len(m.Values[0]); i++ {	
		r.Values[i] = make([]int, len(m.Values))
	}
	for i := 0; i < len(m.Values); i++ {		
		for j := 0; j < len(m.Values[0]); j++ {	
			r.Values[j][i] = m.Values[i][j]	
		}
	}
	return r
}

// podmatica, neobsahuje i ty riadok a jty stlpec 
func (m Matrix)podMatrix(j int) Matrix {
  var r Matrix 
  r.N = m.N-1
  r.Values = m.Values[1:]
  r = r.transponujMatrix()
  r.Values = append(r.Values[:j], r.Values[j+1:]...)
  r = r.transponujMatrix()
  return r
}

 func Determinant(m *Matrix) int{
    if m.N == 2 {
      return m.Values[0][0] * m.Values[1][1] - m.Values[0][1] * m.Values[1][0]
    }
    det := 0
    n := m.N
    for n > 0 {
      n--
      r := m.podMatrix(n)
      x := m.Values[0][n]*Determinant(&r)
      if n%2 == 1{
        x = -x
      }
      det += x
    }  
    return det
 }


// ULOHA 3

/*
1. 
kombinacne cisla

2.
V tabulke je vytvoreny Pascalov trojuholnik tak,
ze jeho spic je v poli [0][0]. Jednotky su vynechane
a miesto toho su prvotnym vstupom do kanalov krajnych
agentov (1. riadok a stlpec). Gorutina s agentom na 
pozicii [0][0] vykona svoju pracu prva a naplni
chybajuci (zatial prazdny) kanal agentov na poziciach 
[0][1] a [1][0]. Postupne sa takto paralelne plni cela
tabulka.

3.
Agent nerobi nic, len uchovava cislo svojho riadku a
stlpca v tabulke. Taktiez si pamata dva kanaly.
Gorutiny, ktore vyuzivaju agentov, si najprv z kanalov
im priradeneho agenta vyberu dve hodnoty, spocitaju ich
a vysledok poslu agentovi, ktori je v tabulke pod nim a 
za nim. Taktiez na konzolu vypisuju kombinacne cislo
v tvare n, k, comb(n, k).

4. 
NxN = N^2 (v kazdom policku gridu je jeden)
V tomto konkretnom pripade je N=10, teda 
10x10 = 100

5.
Pascalov trojuholnik
*/



func main() {
  //fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))

  /*m := Matrix{}

  m.N = 2
  m.Values = [][]int{
    []int{1, 2},
    []int{3, 4},
  }
  println(Determinant(&m)) // -2

  m.N = 3
  m.Values = [][]int{
    []int{1, 2, 3},
    []int{4, 5, 6},
    []int{7, 8, 9},
  }
  println(Determinant(&m)) // 0

  m.N = 3
  m.Values = [][]int{
    []int{4, 2, 3},
    []int{3, 3, 3},
    []int{4, 5, 6},
  }
  println(Determinant(&m)) // 9
  */
}