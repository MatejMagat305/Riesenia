package determinant

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Number int
type Matrix [][]Number
type Vector []Number

func (m Matrix) Sub(t int, k int) Matrix {
	s := make([][]Number, len(m) - 1)
	newRow := 0
	newCol := 0
	for oldRow := 0; oldRow < len(m); oldRow++ {
		if oldRow == t {
			continue
		}
		newCol = 0
		s[newRow] = make([]Number, len(m[0]) - 1)
		for oldCol := 0; oldCol < len(m[0]); oldCol++ {
			if oldCol == k {
				continue
			}
			s[newRow][newCol] = m[oldRow][oldCol]
			newCol++
		}
		newRow++
	}
	return s
}

/*
 Zlozitost Laplacovej formuly na vypocet determinantu som popisal uz v predchadzajucej ulohe, kopia:

 Zjavne O(n!),
 pretoze pre maticu velkosti n vykoname n krat vypocet na matici velkosti n-1
 pre kazdu z podmatic n-1 vykoname n-1 krat operaciu na matici n-2 atd.
 Inak povedane T(n) = n*T(n-1) + O(n)
 Takze zlozitost nic moc, no programuje sa to velmi prijemne :)
*/
func (m Matrix) DeterminantSequence() Number {
	if len(m) == 1 {
		return m[0][0]
	}

	det := Number(0)
	for k := 0; k < len(m); k++ {
		c := Number(1)
		if k % 2 == 1 {
			c = Number(-1)
		}
		s := m.Sub(0, k)
		det += m[0][k] * c * s.DeterminantSequence()
	}
	return det
}

/*
 Rychlost paralelneho algoritmu ma najskor nemilo prekvapila - je podstatne pomalsi ako sekvencny
 (odkomentuj riadky v Time Compare). Po zamysleni to ale dava zmysel a riesili sme to aj na prednaske,
 prilis vela go rutin program spomali, pretoze neustale bojuju o CPU cas. No a v algoritme ktory je O(n!), sa
 ich navytvara naozaj masivne vela... Existuje vsak riesenie, pozri DeterminantLimiter nizsie.
 */
func (m Matrix) DeterminantParallel() Number {
	if len(m) == 1 {
		return m[0][0]
	}

	if len(m) <= 4 {
		m.DeterminantSequence()
	}

	wg := sync.WaitGroup{}
	wg.Add(len(m))
	det := Number(0)
	for k := 0; k < len(m); k++ {
		sub := m.Sub(0, k)
		c := Number(1)
		if k % 2 == 1 { c = Number(-1) }

		go func(d int) {
			det += m[0][d] * c * sub.DeterminantParallel()
			wg.Done()
		}(k)
	}

	wg.Wait()
	return det
}

/*
 Implementacna inspiracia clankom: https://medium.com/@_orcaman/when-too-much-concurrency-slows-you-down-golang-9c144ca305a
 Hranica granularity (4) a limit poctu go rutin (10) su konstanty, ktore po experimentoch dobre bezia na mojej masine...
 Myslienka je, ze ked uz mame pustenych prilis vela go rutin, tak radsej ideme sekvencne, dokym sa neuvolnia...
*/
var limiter = make(chan bool, 10)
func (m Matrix) DeterminantLimiter() Number {
	if len(m) == 1 {
		return m[0][0]
	}
	if len(m) <= 4 {
		m.DeterminantSequence()
	}

	wg := sync.WaitGroup{}
	wg.Add(len(m))
	det := Number(0)
	for k := 0; k < len(m); k++ {
		sub := m.Sub(0, k)
		c := Number(1)
		if k % 2 == 1 { c = Number(-1) }

		select {
			case limiter <- true:
				go func(d int) {
					det += m[0][d] * c * sub.DeterminantLimiter()
					_ = <- limiter
					wg.Done()
				}(k)
			default:
				det += m[0][k] * c * sub.DeterminantSequence()
				wg.Done()
		}
	}
	wg.Wait()
	return det
}

func RandomMatrix(size int) Matrix {
	m := make(Matrix, size)
	for i := 0; i < size; i++ {
		r := make([]Number, size)
		for j:=0; j < size; j++ {
			r[j] = Number(rand.Int31n(5))
		}
		m[i] = r
	}
	return m
}

/*
 n = 10, sequence: 736.1756ms
 n = 10, parallel with limiter: 158.0258ms

 n = 11, sequence: 8.2618556s
 n = 11, parallel with limiter: 1.9304268s

 n = 12, sequence: 1m40.0837812s
 n = 12, parallel with limiter: 23.7683548s !!!

 Myslim si, ze moze byt :) Z O(n!) toho clovek uz asi viac nedostane.
 */
func TimeCompare() {
	sizes := []int{10, 11, 12}
	var start time.Time
	for _, n := range sizes {
		m := RandomMatrix(n)
		start = time.Now()
		_ = m.DeterminantSequence()
		fmt.Printf("n = %d, sequence: %v\n", n, time.Since(start))

		/*
		start = time.Now()
		_ = m.DeterminantParallel()
		fmt.Printf("n = %d, parallel: %v\n" , n, time.Since(start))
		fmt.Println()
		*/

		start = time.Now()
		_ = m.DeterminantLimiter()
		fmt.Printf("n = %d, parallel with limiter: %v\n" , n, time.Since(start))
		fmt.Println()
	}
}

func Test() {
	TimeCompare()
}
