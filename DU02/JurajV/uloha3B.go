package main

import "golang.org/x/tour/pic"
// go get golang.org/x/tour/pic

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for x := range res {
		res[x] = make([]uint8, dx)
		for y := range res[x] {
			//res[x][y] = uint8((x + y) / 2)
			//res[x][y] = uint8(x * y)
			res[x][y] = uint8(x ^ y)
		}
	}
	return res
}

func main() {
	// func Show(f func(int, int) [][]uint8)
	pic.Show(Pic) // v GoLande neukazuje obrazky, teda aspon neviem ako
}
