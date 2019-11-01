package main

import "fmt"

type Matrix struct {
	N      int
	Values [][]int
}

func getCofactor(mat[][] int, temp[][] int, p int, q int, n int) {
	i, j := 0, 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j] = mat[row][col]
				j++
				if j == n - 1 {
					j = 0
					i++
				}
			}
		}
	}
}

func determinantOfMatrix(mat[][] int, n int) int {
	result := 0
	if n == 1 {
		return mat[0][0]
	}
	temp := make([][]int, n)
	for i := range temp {
		temp[i] = make([]int, n)
	}
	sign := 1
	for f := 0; f < n; f++ {
		getCofactor(mat, temp, 0, f, n)
		result += sign * mat[0][f] * determinantOfMatrix(temp, n - 1)
		sign = -sign
	}
	return result
}


func main() {
	m := Matrix{4, [][] int {{1, 0, 2, -1}, {3, 0, 0, 5}, {2, 1, 4, -3},	{1, 0, 5, 0},}}
	//m := Matrix{5, [][] int {{6, 4, 5, 3, 3}, {8, 3, 5, 1, 9}, {9, 5, 3, 6, 8}, {2, 1, 1, 6, 3}, {7, 9, 1, 1, 6},}}
	//m := Matrix{6, [][] int {{-6, -2, -9, -8, 7, -5}, {-6, -2, 5, -5, 4, 7}, {5, -6, -7, -7, -9, 2}, {6, -1, 4, -5, -3, -8}, {-6, 2, -3, 3, -2, -6}, {-3, -5, 1, -9, -5, -2},}}
	//m := Matrix{7, [][] int {{-5, -3, -5, 1, -4, -6, 1}, {-2, 4, -2, -4, -7, -3, 7}, {-6, -10, -10, -3, 7, 9, 2}, {-7, 3, -8, 4, -9, 9, 9}, {-1, 0, -6, 8, 8, -5, 5}, {6, 3, -6, -1, -3, 0, -5}, {-4, -8, 1, -8, -1, 3, 1},}}
	fmt.Print(determinantOfMatrix(m.Values, m.N))

}