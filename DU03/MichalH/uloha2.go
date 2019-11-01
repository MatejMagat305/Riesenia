package main

import (
    "fmt"
)

type Matrix struct {
   N      int
   Values [][]int
}

// transponuje maticu
func transpose(m Matrix) Matrix {
	r := make([][]int, len(m.Values[0]))
	for i := range m.Values[0] {
		r[i] = make([]int, len(m.Values))
	}
	for i := range m.Values {
		for j := range m.Values[i] {
			r[j][i] = m.Values[i][j]
		}
	}
	return Matrix{m.N, r} 
}

// odstrani i-ty riadok a j-ty stlpec
func removeij(m Matrix, i, j int) Matrix {
    val1 := append([][]int(nil), m.Values[:i]...)
    val1 = append(val1, m.Values[i+1:]...)
    r := Matrix{
        m.N-1,
        val1,
    }
    r = transpose(r)
    
    val2 := append([][]int(nil), r.Values[:j]...)
    val2 = append(val2, r.Values[j+1:]...)
    r = Matrix{
        r.N,
        val2,
    }
    r = transpose(r)

    return r
}

func determinant(m Matrix) float64 {
    if m.N == 1 {
        return float64(m.Values[0][0])
    }
    det := .0
    for i := 0; i < m.N; i++ {
        det1 := float64(m.Values[0][i]) * determinant(removeij(m,0,i))
        if i % 2 == 0 {
            det += det1
        } else {
            det -= det1
        }
    }
    return det 
}

func main() {
    m := Matrix{
        4,
        [][]int{{1,2,3,4},{1,0,2,0},{0,1,2,3},{2,3,0,0}},
    }
    fmt.Println(m)
    fmt.Println(determinant(m))
}
