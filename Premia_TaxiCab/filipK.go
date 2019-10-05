/*
pocet stvorcov=1, cislo=2
pocet stvorcov=2, cislo=50
pocet stvorcov=3, cislo=325
pocet stvorcov=4, cislo=1105
pocet stvorcov=5, cislo=8125
pocet stvorcov=6, cislo=5525
pocet stvorcov=7, cislo=105625
pocet stvorcov=8, cislo=27625
pocet stvorcov=9, cislo=71825
pocet stvorcov=10, cislo=138125
pocet stvorcov=11, cislo=5281250
pocet stvorcov=12, cislo=160225
pocet stvorcov=13, cislo=1221025
pocet stvorcov=14, cislo=2442050
pocet stvorcov=15, cislo=1795625
pocet stvorcov=16, cislo=801125
pocet stvorcov=17, cislo=446265625
pocet stvorcov=18, cislo=2082925
pocet stvorcov=19, cislo=41259765625
pocet stvorcov=20, cislo=4005625
pocet stvorcov=21, cislo=44890625
pocet stvorcov=22, cislo=30525625
pocet stvorcov=23, cislo=61051250
pocet stvorcov=24, cislo=5928325
pocet stvorcov=25, cislo=303460625
pocet stvorcov=26, cislo=53955078125
pocet stvorcov=27, cislo=35409725
pocet stvorcov=28, cislo=100140625
pocet stvorcov=29, cislo=1289367675781250
pocet stvorcov=30, cislo=52073125
pocet stvorcov=31, cislo=763140625
pocet stvorcov=32, cislo=29641625
pocet stvorcov=33, cislo=28056640625
pocet stvorcov=34, cislo=33721923828125
pocet stvorcov=35, cislo=7586515625
pocet stvorcov=36, cislo=77068225
pocet stvorcov=37, cislo=5158830625
pocet stvorcov=38, cislo=10317661250
pocet stvorcov=39, cislo=701416015625
pocet stvorcov=40, cislo=148208125
pocet stvorcov=41, cislo=2053764050
pocet stvorcov=42, cislo=1301828125
pocet stvorcov=43, cislo=8716125488281250
pocet stvorcov=44, cislo=62587890625
pocet stvorcov=45, cislo=885243125
pocet stvorcov=46, cislo=2356840332031250
pocet stvorcov=47, cislo=108951568603515625
pocet stvorcov=48, cislo=243061325
pocet stvorcov=49, cislo=476962890625
pocet stvorcov=50, cislo=8800358125
pocet stvorcov=51, cislo=438385009765625
pocet stvorcov=52, cislo=128970765625
pocet stvorcov=53, cislo=257941531250
pocet stvorcov=54, cislo=1310159825
pocet stvorcov=55, cislo=4741572265625
pocet stvorcov=56, cislo=3705203125
pocet stvorcov=57, cislo=10959625244140625
pocet stvorcov=58, cislo=11924072265625
pocet stvorcov=59, cislo=23848144531250
pocet stvorcov=60, cislo=1926705625
pocet stvorcov=61, cislo=2692548668925781250
pocet stvorcov=62, cislo=1490902050625
pocet stvorcov=63, cislo=22131078125
pocet stvorcov=64, cislo=1215306625
pocet stvorcov=65, cislo=118539306640625
pocet stvorcov=66, cislo=813642578125
pocet stvorcov=67, cislo=25672050625
pocet stvorcov=68, cislo=51344101250
pocet stvorcov=69, cislo=6849765777587890625
pocet stvorcov=70, cislo=220008953125
pocet stvorcov=71, cislo=33656858361572265625
pocet stvorcov=72, cislo=3159797225
pocet stvorcov=73, cislo=21796059390625
pocet stvorcov=74, cislo=43592118781250
pocet stvorcov=75, cislo=149606088125
pocet stvorcov=76, cislo=7452545166015625
pocet stvorcov=77, cislo=801325712890625
pocet stvorcov=78, cislo=20341064453125
pocet stvorcov=79, cislo=375255382323302910663187503814697265625
pocet stvorcov=80, cislo=6076533125
pocet stvorcov=81, cislo=37994634925
pocet stvorcov=82, cislo=80606728515625
pocet stvorcov=83, cislo=161213457031250
pocet stvorcov=84, cislo=48167640625
pocet stvorcov=85, cislo=74087066650390625
pocet stvorcov=86, cislo=372627258300781250
pocet stvorcov=87, cislo=37272551265625
pocet stvorcov=88, cislo=2315751953125
pocet stvorcov=89, cislo=11726730697603215958224609494209289550781250
pocet stvorcov=90, cislo=32753995625
pocet stvorcov=91, cislo=20033142822265625
pocet stvorcov=92, cislo=15280246734619140625
pocet stvorcov=93, cislo=21035536475982666015625
pocet stvorcov=94, cislo=641801265625
pocet stvorcov=95, cislo=1283602531250
pocet stvorcov=96, cislo=12882250225
pocet stvorcov=97, cislo=2015168212890625
pocet stvorcov=98, cislo=37181513078125
pocet stvorcov=99, cislo=13831923828125
pocet stvorcov=10000, cislo=377533286937122430335333125
pocet stvorcov=100000, cislo=165090530203990597249775003263135625
pocet stvorcov=1000000, cislo=125262652593971298607146638685883361650570625




*/





    // taxi_cab
    package main
     
    import (
    	"fmt"
    	"math"
    	"math/big"
    	"sort"
    )
     
    ///////////////matika je vysvetlena tu http://mathworld.wolfram.com/SumofSquaresFunction.html
    /////////////trochu paralelnosti :)
    //////////opravene pretecenia :D
    func prime_fact(x int) map[int]int {
    	n := x
    	ret := make(map[int]int)
    	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))) && n > 1; i++ {
    		for n%i == 0 {
    			n = n / i
    			if _, ok := ret[i]; ok {
    				ret[i]++
    			} else {
    				ret[i] = 1
    			}
    		}
    	}
    	if n != 1 {
    		if _, ok := ret[n]; ok {
    			ret[n]++
    		} else {
    			ret[n] = 1
    		}
    	}
    	return ret
    }
     
    func nice_primes(kolko int, kam chan int) {
    	counter := 0
    	akt := 2
    	boli := []int{2}
    	for {
    		akt++
    		for _, i := range boli {
    			if float64(i) > math.Sqrt(float64(akt)) {
    				boli = append(boli, akt)
    				if akt%4 == 1 {
    					kam <- akt
    					counter++
    					if counter == kolko {
    						close(kam)
    						return
    					}
    				}
    				break
    			}
    			if akt%i == 0 {
    				break
    			}
    		}
    	}
    }
     
    func factored(x int, kanaly ...chan []int) [][]int {
    	h := prime_fact(x)
    	vys := [][]int{{x}}
    	for _, kanal := range kanaly {
    		kanal <- vys[0]
    	}
    	for k, _ := range h {
    		if x/k == 1 {
    			continue
    		}
    		for _, v := range factored(x / k) {
    			akt := []int{k}
    			for _, i := range v {
    				akt = append(akt, i)
    			}
    			sort.Slice(akt, func(i, j int) bool { return akt[i] > akt[j] })
    			vys = append(vys, akt)
    			for _, kanal := range kanaly {
    				kanal <- akt
    			}
    		}
    	}
    	for _, v := range kanaly {
    		close(v)
    	}
    	return vys
    }
    func vrat(co []int, kam chan []int) {
    	kam <- co
    }
    func najmensie_n_pre_pocet_stvorcov(n int) *big.Int {
    	primes := []int{}
    	primes_chan := make(chan int, 10)
    	b_parne := 2 * n
    	b_parne_chan := make(chan []int, 10)
    	b_neparne_e := 2*n + 1
    	b_neparne_e_chan := make(chan []int, 10)
    	b_neparne_o := 2*n - 1
    	b_neparne_o_chan := make(chan []int, 10)
    	best := new(big.Int)
    	best.SetInt64(int64(-1))
    	go nice_primes(b_neparne_e, primes_chan)
    	for {
    		h, ok := <-primes_chan
    		if !ok {
    			primes_chan = nil
    			break
    		} else {
    			primes = append(primes, h)
    		}
    	}
    	go factored(b_parne, b_parne_chan)
    	go factored(b_neparne_e, b_neparne_e_chan)
    	go factored(b_neparne_o, b_neparne_o_chan)
    	for {
    		select {
    		case pom, ok := <-b_parne_chan:
    			if !ok {
    				b_parne_chan = nil
    			} else {
    				akt := new(big.Int)
    				akt.SetInt64(1)
    				for i, v := range pom {
    					tmp := new(big.Int)
    					tmp.SetInt64(int64(primes[i]))
    					tmp2 := new(big.Int)
    					tmp2.SetInt64(int64(v - 1))
    					tmp.Exp(tmp, tmp2, nil)
    					akt.Mul(akt, tmp)
    				}
    				if akt.Cmp(best) < 0 || best.Cmp(new(big.Int).SetInt64(-1)) == 0 {
    					best = akt
    					//fmt.Println(best)
    				}
     
    			}
     
    		case pom, ok := <-b_neparne_e_chan:
    			if !ok {
    				b_neparne_e_chan = nil
    			} else {
    				akt := new(big.Int)
    				akt.SetInt64(int64(1))
    				for i, v := range pom {
    					tmp := new(big.Int)
    					tmp.SetInt64(int64(primes[i]))
    					tmp2 := new(big.Int)
    					tmp2.SetInt64(int64(v - 1))
    					tmp.Exp(tmp, tmp2, nil)
    					akt.Mul(akt, tmp)
    				}
    				if akt.Cmp(best) < 0 || best.Cmp(new(big.Int).SetInt64(-1)) == 0 {
    					best = akt
    					//fmt.Println(best)
    				}
     
    			}
    		case pom, ok := <-b_neparne_o_chan:
    			if !ok {
    				b_neparne_o_chan = nil
    			} else {
    				akt := new(big.Int)
    				akt.SetInt64(int64(2))
    				for i, v := range pom {
    					tmp := new(big.Int)
    					tmp.SetInt64(int64(primes[i]))
    					tmp2 := new(big.Int)
    					tmp2.SetInt64(int64(v - 1))
    					tmp.Exp(tmp, tmp2, nil)
    					akt.Mul(akt, tmp)
    				}
    				if akt.Cmp(best) < 0 || best.Cmp(new(big.Int).SetInt64(-1)) == 0 {
    					best = akt
    					//fmt.Println(best)
    				}
    			}
     
    		}
    		if b_neparne_e_chan == nil && b_neparne_o_chan == nil && b_parne_chan == nil {
    			return best
    		}
    	}
    	return best
    }
     
    func main() {
    	i := 1
    	for i < 100 {
    		fmt.Printf("pocet stvorcov=%d, cislo=%d", i, najmensie_n_pre_pocet_stvorcov(i))
    		fmt.Println()
    		i++
    	}
    	i = 10000
    	fmt.Printf("pocet stvorcov=%d, cislo=%d", i, najmensie_n_pre_pocet_stvorcov(i))
    	fmt.Println()
    	i = 100000
    	fmt.Printf("pocet stvorcov=%d, cislo=%d", i, najmensie_n_pre_pocet_stvorcov(i))
    	fmt.Println()
    	i = 1000000
    	fmt.Printf("pocet stvorcov=%d, cislo=%d", i, najmensie_n_pre_pocet_stvorcov(i))
    	fmt.Println()
    	i = 10000000
    	fmt.Printf("pocet stvorcov=%d, cislo=%d", i, najmensie_n_pre_pocet_stvorcov(i))
    }