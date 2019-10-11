// trpaslik
package main
 
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
 
//postupnost 0,1,-1,2,-2,3,-3...
func clen_postupnosti(n int) int {
	return n / 2 * int(math.Pow(-1.0, float64(n)))
}
 
func kantor_reverse(n int) (r int, s int) {
	w := int(math.Floor((math.Sqrt(float64(8*n+1)) - 1) / 2))
	t := int((w*w + w) / 2)
	y := n - t
	return y, w - y
}
 
func kam_strielat(t int) int {
	zac, krok := kantor_reverse(t)
	return clen_postupnosti(zac+1) + t*clen_postupnosti(krok+1)
}
 
var shots []int = []int{}
var tt int = 0
 
/*efektivnejsie funkcia na pocet vystrelov
snazi sa neskusat moznosti, ktore uz boli.
pamatovo narocne a vypoctovo znacne pomalsie riesenie ako
prve*/
func kam_strielat2(t int) int {
	var zac, krok int
	nasiel := false
	for !nasiel {
		nasiel = true
		zac, krok = kantor_reverse(tt)
		for i, p := range shots {
			if zac+i*krok == p {
				tt++
				nasiel = false
				break
			}
		}
	}
	kam := clen_postupnosti(zac+1) + t*clen_postupnosti(krok+1)
	tt++
	shots = append(shots, kam)
	return kam
}
 
func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 30; i++ {
		trpaslik_zac := rand.Intn(100) - 50
		trpaslik_krok := rand.Intn(100) - 50
		a := true
		b := true
		t := 0
		fmt.Println(trpaslik_zac)
		fmt.Println(trpaslik_krok)
		shots = []int{}
		tt = 0
		for a || b {
			if trpaslik_zac+trpaslik_krok*t == kam_strielat(t) && a {
				fmt.Print("kam_strielat trafil v kroku:")
				fmt.Println(t)
				a = false
			}
			if trpaslik_zac+trpaslik_krok*t == kam_strielat2(t) && b {
				fmt.Print("kam_strielat2 trafil v kroku:")
				fmt.Println(t)
				b = false
			}
			t++
		}
		fmt.Println("----------")
	}
}