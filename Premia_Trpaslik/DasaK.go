//trpaslik
package main
 
import (
	"fmt"
	"math"
)
 
func mapa(x int) int {
	//N -> Z (0,1,2,3,4,5,6,...) -> (0,1,-1,2,-2,3,-3,...)
	if x%2 == 0 {
		return x / (-2)
	}
	return int(x/2) + 1
}
 
//https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
func reverse(c int) (a int, b int) {
	//c = 1/2(a + b)(a + b + 1) + b
	//inverting cantor pairing function to calculate a, b from c
	w := int(math.Floor((math.Sqrt(float64(8*c+1)) - 1) / 2))
	t := int((w*w + w) / 2)
	b = c - t
	a = w - b
	return mapa(b), mapa(a)
}
 
func trpaslik(t int) int {
	a, b := reverse(t)
	//fmt.Println(a, b)
	return a + t*b
}
 
func main() {
	for i := 0; i < 2000; i++ {
		fmt.Printf("Time %d, aim at %d.\n", i, trpaslik(i))
	if trpaslik(i) == 21 {
	  break
	}
	}
}