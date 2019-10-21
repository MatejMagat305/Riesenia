package main

//presne podla navodu v tutoriale
func Sqrt(x float64) float64 {
	z := 1.0
	delta := 0.00001
	for ; z*z-x > delta || z*z-x < -delta; z -= (z*z - x) / (2 * z) {
	}
	return z
}

/*func main() {
	fmt.Println(Sqrt(81))

}*/
