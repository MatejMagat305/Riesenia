package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	pole := make([][]uint8, dy)
	for i, _ := range pole {
		pole[i] = make([]uint8, dx)
	}
	for i, _ := range pole {
		for j, _ := range pole[i] {
			//pole[i][j] = uint8((i + j) / 2)
			//pole[i][j] = uint8((i * j) / 8)
			//pole[i][j] = uint8((i + j))
			//pole[i][j] = uint8(math.Abs(float64(i)) + math.Abs(float64(j)))
			//pole[i][j] = uint8(math.Pow(float64(i), 2.5) + math.Pow(float64(j), 2.5)) //kruzky

			pole[i][j] = uint8(4*math.Pow(float64(i), 2) + 4*math.Pow(float64(j), 2))
		}
	}
	return pole
}

func main() {
	pic.Show(Pic)
}
