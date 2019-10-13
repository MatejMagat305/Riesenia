package main

import (
	"fmt"
	"math"
	"time"
)

//funkcia greedy zistuje kolko minci
//potrebuje na vyskladanie sumy {sum}
func greedy(sum int, pole []int, length int) int {
	//skladanie sumy
	i := length - 1
	counter := 0
	for sum != 0 {
		if pole[i] <= sum {
			sum2 := sum % pole[i]
			counter += ((sum - sum2) / pole[i])
			sum = sum2
		}
		i -= 1
	}
	return counter
}

//funkcia greedy2 zaokruhluje sumu nahor podla roznych
//minci z pola, vysledny pocet pouzitich minci bude ->
// greedy(zaokruhlena suma) + greedy(zaokruhlena suma - suma)
//pocitam vlastne pocet minci co by clovek "zaplatil" + tie co
//by sa mu "vydali", zapamatavam si najlepsie riesenie
func greedy2(sum int, pole []int, length int) int {
	//vydavanie
	i := length - 1
	best := 1<<63 - 1
	for i > 0 {
		pom := (sum % pole[i])
		if pom == 0 {
			i -= 1
			break
		}
		sum2 := sum + (pole[i] - pom)
		counter := greedy(sum2, pole, length) + greedy((pole[i]-pom), pole, length)
		if counter < best {
			best = counter
		}
		i -= 1
	}
	return best
}

//funkcia biland hlada najmensiu sumu pre vstupne n
func biland(n int, pole []int, length int) int {
	for i := 0; i < (1<<63 - 1); i += 1 {
		a := greedy(i, pole, length)
		b := greedy2(i, pole, length)
		if b < a {
			a = b
		}
		if a == n {
			return i
		}
	}
	return -1
}

//program funguje aj pre eura a dolare, kedze funkcie greedy
//maju ako vstupny parameter pole, predpoklad pola je ze je
//usporiadane podla hodnoty minci
func main() {
	var coins = []int{1, 2, 5, 10, 20, 50}
	//var euro = []int{1, 2, 5, 10, 20, 50, 100, 200}
	//var dollar = []int{1, 5, 10, 25, 50, 100}
	for i := 0; i < 20; i += 1 {
		fmt.Printf("najmensia mozna suma pre n %d je: %d\n", i, biland(i, coins, 6))
		//fmt.Printf("najmensia mozna suma pre n %d je: %d\n", i, biland(i, euro, 8)) //euro
		//fmt.Printf("najmensia mozna suma pre n %d je: %d\n", i, biland(i, dollar, 6))	//dollar
	}
	for i := 1; i < 5; i += 1 {
		j := int(100 * math.Pow(10, float64(i)))
		start := time.Now()
		fmt.Printf("najmensia mozna suma pre n %d je: %d, ", j, biland(j, coins, 6))
		end := time.Since(start)
		fmt.Printf("%d trvala: %s\n", j, end)
	}
	// najmensia mozna suma pre n 1000 je: 49833, 1000 trvala: 23.9566ms
	// najmensia mozna suma pre n 10000 je: 499833, 10000 trvala: 232.5537ms
	// najmensia mozna suma pre n 100000 je: 4999833, 100000 trvala: 2.3818507s
	// najmensia mozna suma pre n 1000000 je: 49999833, 1000000 trvala: 22.8804216s
}
