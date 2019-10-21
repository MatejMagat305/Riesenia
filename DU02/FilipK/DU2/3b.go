package main

//import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i, _ := range ret {
		ret[i] = make([]uint8, dx)
	}
	for i, _ := range ret {
		for j, _ := range ret[i] {
			x := float64(float64(j)/float64(dy)*3.4 - 1.7)
			y := -float64(float64(i)/float64(dx)*4 - 2)
			if math.Pow((math.Pow(x, 2)+math.Pow(y, 2)-1), 3)-math.Pow(x, 2)*math.Pow(y, 3) < 0.0001 {
				ret[i][j] = 42
			} else {
				ret[i][j] = 200
			}
		}
	}
	return ret
}

/*func main() {
	pic.Show(Pic)
}*/
