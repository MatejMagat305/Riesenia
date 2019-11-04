// determinant2
package main

/*"fmt"
"math/rand"
"time"*/

type Matrix struct {
	N      int
	Values [][]int
}

/* wikipedia lahko prezradi zlozitost ktoru ani nie problem vypocitat kedze napr. pre maticu
10*10 treba 10krat spocitat maticu velkosti 9*9 pre 9 potom 9krat 8*8... teda n!
lepsi algoritmus som programoval na minulej domacej ;)
The Laplace expansion is computationally inefficient for high dimension matrices, with a time
complexity in big O notation of {\displaystyle O(n!)}O(n!). Alternatively, using a decomposition
into triangular matrices as in the LU
decomposition can yield determinants with a time complexity of {\displaystyle O(n^{3})}O(n^{3}).[1]
*/

//transpose transponuje maticu
func (m *Matrix) transpose() {
	r := make([][]int, len(m.Values[0]))
	for i := range m.Values[0] {
		r[i] = make([]int, len(m.Values))
	}
	for i := range m.Values {
		for j := range m.Values[0] {
			r[j][i] = m.Values[i][j]
		}
	}
	m.Values = r
}

// odtsrini z matice i ty riadok a j ty stlpec vrati novu, vsetky hodnoty sa kopiruju
func (m Matrix) trim(i int, j int) *Matrix {
	ret := [][]int{}
	for i := range m.Values {
		akt := make([]int, len(m.Values[i]))
		copy(akt, m.Values[i])
		ret = append(ret, akt)
	}
	ret = append(ret[:i], ret[i+1:]...)
	r := new(Matrix)
	r.N = m.N - 1
	r.Values = ret
	r.transpose()
	r.Values = append(r.Values[:j], r.Values[j+1:]...)
	r.transpose()
	return r
}

//determinant cez laplaceho formulu da sa zadat granularit pre paralelny vypocet
func Determinant(m *Matrix, granularity ...int) int {
	det := 0
	ch := make(chan int)
	if m.N == 2 {
		return m.Values[0][0]*m.Values[1][1] - m.Values[1][0]*m.Values[0][1]
	}
	if len(granularity) > 0 && m.N > granularity[0] {
		for i := 0; i < m.N; i++ {
			go func(kam chan int, mat *Matrix, n int) {
				akt := m.Values[0][n] * Determinant(mat, granularity[0])
				if n%2 == 1 {
					akt *= -1
				}
				kam <- akt
			}(ch, m.trim(0, i), i)
		}
		for i := 0; i < m.N; i++ {
			det += <-ch
		}
	} else {
		for i := 0; i < m.N; i++ {
			akt := m.Values[0][i] * Determinant(m.trim(0, i))
			if i%2 == 1 {
				akt *= -1
			}
			det += akt

		}
	}
	return det
}

/*func main() {

	mm := new(Matrix)
	mm.N = 4
	mm.Values = [][]int{{1, 7, 3, 4}, {3, 4, 5, 6}, {7, 8, 25, 10}, {11, 12, 13, 14}} //ma byt 1920
	fmt.Println(Determinant(mm, 2))
	mm.N = 3
	mm.Values = [][]int{{-2, 2, -3}, {-1, 1, 3}, {2, 0, -1}}
	fmt.Println(Determinant(mm, 2))
	mm.N = 10
	mm.Values = make([][]int, mm.N)
	rand.Seed(time.Now().UnixNano())
	for i := range mm.Values {
		mm.Values[i] = make([]int, mm.N)
		for j := range mm.Values[i] {
			mm.Values[i][j] = rand.Intn(1000)
		}
	}
	//casove dostihy pre rozne granularity
	// vysledok velmi zavisi od poctu jadier
	//a nebude mat velky zmysel robit ovela viac
	// routine ako jadier cize best bude okolo velkostmatice-1 alebo 2
	for i := 10; i > 1; i -= 1 {
		start := time.Now()
		fmt.Println("granularita:", i)
		fmt.Println(Determinant(mm, i))
		fmt.Println(time.Since(start))
	}
}*/
