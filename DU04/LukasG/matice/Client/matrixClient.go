package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Cislo int64
type Matica [][]Cislo
type Vektor []Cislo


func (m Matica) Print(headline string) {
	fmt.Println(headline)
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%v ", m[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func generujMaticu(n int) Matica {
	m := make(Matica, n)
	for i := range m {
		m[i] = make(Vektor, n)
		for j := range m[i] {
			m[i][j] = Cislo(rand.Float64())
		}
	}
	return m
}

func transponujMaticu(m Matica) Matica {
	r := make(Matica, len(m[0]))
	for i := range m {
		r[i] = make(Vektor, len(m))
	}
	for i := range m {
		for j := range m[i] {
			r[j][i] = m[i][j]
		}
	}
	return r
}


func main() { // matrixClient.exe n port1, port2, ....
	if len(os.Args) < 2 {
		fmt.Println("pouzitie:")
		fmt.Println("client.exe velkost host1:port1, host2:port2, ....\t- spusti skalarny server na host:porte")
	} else {
		N, _ := strconv.Atoi(os.Args[1]) // velkost matice
		m1 := generujMaticu(N)           // matica 1
		mx2 := generujMaticu(N)          // matica 2
		if N < 4 {                       // vacsie uz aj tak nevieme citat
			m1.Print("Matica A:")
			mx2.Print("Matica B:")
		}
		m2 := transponujMaticu(mx2) // riadky za stlpce

		res := make(Matica, len(m1)) // vysledna matica
		for i := range res {
			res[i] = make(Vektor, len(m2))
		}

		start := time.Now()
		chanRes := make(chan Result)
		wp := NewWorkerPool(os.Args[2:])
		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m2); j++ {
				v1 := m1[i]
				v2 := m2[j]
				task := &Task{i, v1, j, v2}
				wp.addTask(task)
			}
		}

		chanTime := make(chan time.Duration)
		go wp.mainLoop(chanRes, chanTime)
		for ww := 0; ww < len(m1)*len(m2); ww++ {
			result := <-chanRes
			fmt.Printf("prisiel vysledok %v\n", result)
			res[result.I1][result.I2] = result.Skalar
		}

		if N < 4 {
			res.Print("\nMatica A*B:")
		}

		totalTime := time.Since(start)
		clientTime := <- chanTime
		totalMs := float64(totalTime.Milliseconds())
		clientMs := float64(clientTime.Milliseconds())
		fmt.Printf("\nZ celkoveho casu %v client pracoval %v \n", totalTime, clientTime)
		fmt.Printf("Takze jeho vytazenost bola %v percent \n", (clientMs / totalMs) * 100)

		fmt.Println("\nZistujem vytazenie serverov...")
		go wp.getTimes(chanRes)
		for w := 0; w < len(wp.workers); w++ {
			result := <-chanRes
			serverMs := float64(result.Skalar)
			fmt.Printf("Server %v:%v pracoval celkovo %vms, takze jeho vytazenost bola: %v percent \n",
				wp.workers[w].host, wp.workers[w].port, result.Skalar, (serverMs/ totalMs) * 100)
		}
	}
}
