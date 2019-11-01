package main
//https://www.geeksforgeeks.org/determinant-of-a-matrix/
import ("fmt")
type Matrix struct {
	N      int
	Values [][]int
}

func display(mat [][]int, row int, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d", mat[i][j]);
		}
		fmt.Println();
	}
}
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

func determinantOfMatrix(mat [][]int, n int) int {
	D := 0;
	if n == 1 {
		return mat[0][0];
	}
	temp := make([][]int, n)
	for i:=0;i<n;i++{
		b:=make([]int, n)
		temp[i]=b
	}

	sign := 1;
	for f := 0; f < n; f++ {
		getCofactor(mat, temp, 0, f, n);
		D += sign * mat[0][f] * determinantOfMatrix(temp, n-1);
		sign = -sign;
	}

	return D;
}

func Determinant(m *Matrix) int{
	return determinantOfMatrix(m.Values,m.N)
}
func main(){
	N:=4
	mat :=Matrix{N,[][]int{{1, 0, 2, -1},{3, 0, 0, 5},	{2, 1, 4, -3},	{1, 0, 5, 0}	}}

	fmt.Printf("Determinant of the matrix is : %d \n",
		Determinant(&mat))

	mat =Matrix{2, [][]int{{1,2},{3,4}}}

	fmt.Printf("Determinant of the matrix is : %d",
		Determinant(&mat))


}