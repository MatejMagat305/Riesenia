package main

import (
	"math/rand"
)

//https://gist.github.com/n1try/c5082f0f1db7f4abcb6d995dc275fe7f
//https://www.geeksforgeeks.org/determinant-of-a-matrix/
//https://stackoverflow.com/questions/44116183/randomly-generating-a-matrix-in-golang
func generate(nb int) Matrix{
	randMatrix :=make([][]int,nb)
	for i:=0;i<nb ; i++ {
		g:=make([]int,nb)
		randMatrix[i]=g
	}
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			randMatrix[i][j] = rand.Intn(100)
			//looping over each element of array and assigning it a random variable
		}
	}
	return Matrix{nb,randMatrix}
}

type Matrix struct {
	N      int
	Values [][]int
}
//staré riešenie
func getCofactor(mat [][]int, temp [][]int, p int, q int, n int) {
	i := 0
	j := 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			// Copying into temporary matrix only those element
			// which are not in given row and column
			if row != p && col != q {
				temp[i][j] = mat[row][col]
				j++

				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func dajTemp(mat [][]int, n int) [][]int{
	temp := make([][]int, n)
	for i:=0;i<n;i++{
		b:=make([]int, n)
		temp[i]=b
	}
	return temp
}

func determinantOfMatrix(mat [][]int, n int) int {
	D := 0;
	if n == 1 {
		return mat[0][0];
	}
	temp := dajTemp(mat, n)

	sign := 1;
	for f := 0; f < n; f++ {
		getCofactor(mat, temp, 0, f, n);
		c:=sign * mat[0][f] * determinantOfMatrix(temp, n-1);
		D += c
		sign = -sign;
	}

	return D;
}

func Determinant_s(m *Matrix) int{
	return determinantOfMatrix(m.Values,m.N)
}

//koniec starého
/*
func main() {
	X :=Matrix{4, [][]int{{1,3,-4,6},{-2,-5,0,2},{0,-1,2,0},{3,2,8,1}}}
	gr:=3 //granularita
	println(Determinant_s(&X))
	println(Determinant_p_pred(&X,gr))
	println("------------------------")
	var m Matrix= generate(10)
	for ; gr<10;gr++  {
		start := time.Now()
		Determinant_s(&m)
		t:=time.Now()
		println(t.Sub(start))
		start = time.Now()
		Determinant_p_pred(&m,gr)
		t=time.Now()
		println(t.Sub(start))
	}
}
*/
// nove riešenie para

func Determinant_p_pred(matrix *Matrix, i int) int {
	return Determinant_p(matrix.Values,matrix.N,i)
}

func Determinant_p(mat [][]int, n int,gr int) int {
	if n<=gr {
		return determinantOfMatrix(mat,n)
	}
	D := 0;
	if n == 1 {
		return mat[0][0];
	}
	sign := 1;
	ch := make(chan int,n)
	for f := 0; f < n; f++ {
		go func(mat [][]int, n int,gr int, f int, ng int) {
			//print("\n",n,"        ","     " , f,"\n")
			temp:= dajTemp(mat,n)
			getCofactor(mat, temp, 0, f, n)
			vysl :=ng * mat[0][f] * Determinant_p(temp,n-1, gr)
			//print(vysl )
			ch<-vysl}(mat, n, gr,f, sign)
		sign = -sign;
	}
	for  f := 0; f < n; f++{//bez f mi to vyhlosovalo deadlock, neviem prečo.........
		a,b :=<-ch
		if b {
			D+=a
		}
	}

	return D;
}