package tour

/* import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
}
*/
//funguje iba v online prostredi (https://tour.golang.org/moretypes/18)

func Pic(dx, dy int) [][]uint8 {
    picture := make([][]uint8, dy) // skusam si predalokaciu a nasledne indexovanie...
	for y := 0; y < dy; y++ {
	    var row []uint8
		for x := 0; x < dx; x++ {
			//value := (x+y)/2
			//value := x*y
			value := x^y
			row = append(row, uint8(value)) // ...aj pridavanie do dynamickeho pola
		}
		picture[y] = row 
	}
	return picture
}



