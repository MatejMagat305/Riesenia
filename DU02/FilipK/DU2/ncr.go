// ncr
package main

import (
	"fmt"
	"math"
	"math/big"
)

//nevedel som, ze mulrange existuje v bigintoch
//tak mam vlastny :D
func vynasob_range(a, b int, kam chan *big.Int) {
	if a == b {
		kam <- big.NewInt(int64(1))
		return
	}
	if b-a == 1 {
		kam <- big.NewInt(int64(a))
		return
	}
	c := make(chan *big.Int)
	go vynasob_range(a, (a+b)/2, c)
	go vynasob_range((a+b)/2, b, c)
	x := <-c
	y := <-c
	kam <- (x.Mul(x, y))
}

//predpokladam, ze ocakavany pristup divide and conquer
//pouziva viac routine pre male zlepsenie
//musim mat 2 kanaly aby som vedel, ktore cislo je ktore
func ncr(n, k int) *big.Int {
	//n!/(k!(n-k)!)
	c1 := make(chan *big.Int)
	c2 := make(chan *big.Int)
	go vynasob_range(k+1, n+1, c1)
	go vynasob_range(1, (n-k)+1, c2)
	a := <-c1
	b := <-c2
	fmt.Println("-----------")
	return a.Div(a, b)
}

//////////////////////////////////////////////
//riesenie nizsie ma len porazit vypis :)
//////////////////////////////////////////////
//riesenie je optimalizovane pre prostredne hodnoty
//teda neratam co je mensie ci k alebo n-k
//robim velky preprocessing aby som sa vyhol bigintom
//operacie na intoch su jedna strojova instrukcia preto je riesenie
//rychlejsie, algoritmicky je riesenie horsie a pre male vstupy
//bude pomalsie
func ncr2(n, k int) *big.Int {
	rozklady := make(map[int]map[int]int)
	//s malou pomocou dynamiky a budovania zdola je aj prvociselny rozklad
	//milionu cisel vpohode
	// vo for cykle rozkladam vlastne menovatel zlomku
	for i := 2; i <= n-k; i++ {
		akt := make(map[int]int)
		n := i
		for j := 2; j <= int(math.Ceil(math.Sqrt(float64(i)))) && n > 1; j++ {
			if p, ok := rozklady[n]; ok {
				for kluc, hodnota := range p {
					if _, ok1 := akt[kluc]; ok1 {
						akt[kluc] += hodnota
					} else {
						akt[kluc] = hodnota
					}
				}
				break
			} else {
				for n%j == 0 {
					n /= j
					if _, ok1 := akt[j]; ok1 {
						akt[j]++
					} else {
						akt[j] = 1
					}
				}
			}
		}
		if n == i {
			akt[n] = 1
		}
		rozklady[i] = akt
	}
	// tu si urobim jednu mapu prvociselneho rozkladu pre menovatel
	men := make(map[int]int)
	for _, hodnota := range rozklady {
		for kluc, hod := range hodnota {
			if _, ok := men[kluc]; ok {
				men[kluc] += hod
			} else {
				men[kluc] = hod
			}
		}
	}
	//mapa pre citatel
	cit := make(map[int]int)
	for i := k + 1; i <= n; i++ {
		akt := make(map[int]int)
		n := i
		for j := 2; j <= int(math.Ceil(math.Sqrt(float64(i)))) && n > 1; j++ {
			if p, ok := rozklady[n]; ok {
				for kluc, hodnota := range p {
					if _, ok1 := akt[kluc]; ok1 {
						akt[kluc] += hodnota
					} else {
						akt[kluc] = hodnota
					}
				}
				break
			} else {
				for n%j == 0 {
					n /= j
					if _, ok1 := akt[j]; ok1 {
						akt[j]++
					} else {
						akt[j] = 1
					}
				}
			}
		}
		if n == i {
			akt[n] = 1
		}
		for kluc, hodnota := range akt {
			if _, ok := cit[kluc]; ok {
				cit[kluc] += hodnota
			} else {
				cit[kluc] = hodnota
			}
		}
	}
	// delenie big intov je straaaaasne pomale, horsie nez nasobenie tak ho nerobim vobec ;p
	ret := big.NewInt(int64(1)) //nakoniec ale aj tak potrebujem bigint
	for zaklad, mocnina := range cit {
		if _, ok := men[zaklad]; ok {
			mocnina -= men[zaklad]
		}
		z := big.NewInt(int64(zaklad))
		ret = ret.Mul(ret, z.Exp(z, big.NewInt(int64(mocnina)), nil))
		//nasobim uz len tym co zostane
	}
	return ret
}
