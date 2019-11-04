//determinant - Laplace's formula
//https://sk.wikipedia.org/wiki/Determinant_(matematika)
//https://www.khanacademy.org/math/linear-algebra/matrix-transformations/inverse-of-matrices/v/linear-algebra-nxn-determinant

package main

import (
	"fmt"
	"math"
	"time"
)

type Matrix struct {
	N      int
	Values [][]float64
}

//z matice M sa vytvori novama matica M2 bez i-teho riadku a j-teho stlpca
func minor(m *Matrix, i int, j int) Matrix {
	pole := make([][]float64, m.N-1)
	index := 0
	for ii := 0; ii < m.N; ii += 1 {
		if i != ii {
			cpy := make([]float64, 0)
			cpy = append(cpy, m.Values[ii][:j]...)
			cpy = append(cpy, m.Values[ii][j+1:]...)
			pole[index] = cpy
			index += 1
		}
	}
	return Matrix{m.N - 1, pole}
}

//pocitanie determinantu pomocou Laplaceho formuly
//gran - urcuje stupen granularity, ci sa ma determinant pocitat paralelne
func Determinant(m *Matrix, gran int) int {
	//fmt.Println(m)
	if m.N == 2 {
		return int(m.Values[0][0]*m.Values[1][1] - m.Values[0][1]*m.Values[1][0])
	}
	out := 0
	ch := make(chan int)
	if m.N >= gran {
		for i := 0; i < m.N; i += 1 {
			m2 := minor(m, 0, i)
			go func(stack chan int, mx *Matrix, x int) {
				d := float64(Determinant(&m2, gran))
				stack <- int(math.Pow(-1, float64(x)) * m.Values[0][x] * d)
			}(ch, &m2, i)
		}
		for i := 0; i < m.N; i += 1 {
			out += <-ch
		}
	} else {
		for i := 0; i < m.N; i += 1 {
			m2 := minor(m, 0, i)
			d := float64(Determinant(&m2, gran))
			out += int(math.Pow(-1, float64(i)) * m.Values[0][i] * d)
		}
	}
	return int(out)
}

func main() {
	//m := Matrix{2, [][]float64{{1, 2}, {3, 4}}} //-2
	//m := Matrix{3, [][]float64{{-2, 2, -3}, {-1, 1, 3}, {2, 0, -1}}} //18
	//m := Matrix{4, [][]float64{{1, 1, 4, 1}, {2, 2, 10, 6}, {3, 9, 21, 17}, {5, 11, 29, 23}}} //-48
	//m := Matrix{3, [][]float64{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}}} //-360

	// m := Matrix{10, [][]float64{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}} //0
	//granularita 0 -> 1.3534007s
	//granularita 5 -> 740.5317ms
	//granularita 10 -> 243.3402ms

	// m := Matrix{8, [][]float64{
	// 	{7, 6, 4, 8, 7, 6, 4, 8},
	// 	{4, 3, 6, 2, 4, 3, 6, 2},
	// 	{2, -1, 2, -7, 2, -1, 2, -7},
	// 	{0, 2, 5, 1, 0, 2, 5, 1},
	// 	{7, 6, 4, 8, 7, 6, 4, 8},
	// 	{4, 3, 6, 2, 4, 3, 6, 2},
	// 	{2, -1, 2, -7, 2, -1, 2, -7},
	// 	{0, 2, 5, 1, 0, 2, 5, 1}}}
	//granularita 0 -> 33.8879ms
	//granularita 3 -> 30.9178ms
	//granularita 7 -> 3.9916ms

	start := time.Now()
	fmt.Println(Determinant(&m, 0))
	fmt.Println(time.Since(start))
}

//casova zlozitost Laplaceho formuly: O(n!)

//hladanie determinantu pomocu gauss-jordanovej eliminacie: O(n^3)
//(vytvorenie trojuholnikovej matice)

//Nasobenie matic
//Strassen algorithm - algoritmus na nasobenie matic, pocitanie determinantu tymto
//algoritmom by malo zlozitost ~O(n^2.81)
//Coppersmith–Winograd algorithm - lepsie nasobenie matic oproti Stressen. al. (O(n^2.375))
