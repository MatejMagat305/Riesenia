package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	N      int	//zrejme velkost
	Values [][]int
}

func Determinant(m *Matrix) int{
	var det int = 0
	if m.N == 2{
		return m.Values[0][0] * m.Values[1][1] - m.Values[0][1] * m.Values[1][0]
	}
	tmp_mat := new(Matrix)
	tmp_mat.N = m.N - 1
	tmp_mat.Values = make([][]int, m.N - 1)
	for i := 0; i < m.N; i++{
		for j := 0; j < i; j++{
			tmp_mat.Values[j] = make([]int, m.N - 1)
			for k := 1; k < m.N; k++{
				tmp_mat.Values[j][k - 1] = m.Values[j][k]
			}
		}
		for j := i + 1; j < m.N; j++{
			tmp_mat.Values[j - 1] = make([]int, m.N - 1)
			for k := 1; k < m.N; k++{
				tmp_mat.Values[j - 1][k - 1] = m.Values[j][k]
			}
		}
		//fmt.Println("tmp_mat:", tmp_mat)
		det += m.Values[i][0] * Determinant(tmp_mat) * int(math.Pow(-1, float64(i)))
		//fmt.Println("det:", det)
	}
	
	return det
}

func main() {
	mat := new(Matrix)
	mat.N = 3
	var count int = 1
	mat.Values = make([][]int, mat.N)
	for i := 0; i < mat.N; i++{
		mat.Values[i] = make([]int, mat.N)
		for j := 0; j < mat.N; j++{
			mat.Values[i][j] = count
			count++
		}
	}
	mat.Values[0][0] = 2
	
	//fmt.Println("mat:", mat.Values)
	
	res := Determinant(mat)
	fmt.Printf("%d", res)
}
