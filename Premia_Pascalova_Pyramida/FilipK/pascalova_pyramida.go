package main

//program prerobeny z Agentov007 :))
//princip je rovnaky akurat v kocke
// indexy a,b,c su od 0
//nedopocitavam dalsiu cast kocky ako to robili 007
//a vzorec je (a+b+c)!/(a!*b!*c!)
// hrany a strany sa riesia same
// akurat je treba if na prvy pripad 0 0 0
import (
	"fmt"
	"time"
)

type Agent008 struct {
	i, j, k int      // Agent's coordinates in the cube
	a, b, c chan int // inputs for Agent008
}

type Cube [][][]*Agent008

func main() {
	N := 10
	fmt.Println(N)
	cube := make(Cube, N)
	for i := range cube {
		cube[i] = make([][]*Agent008, N)
		for j := range cube[i] {
			cube[i][j] = make([]*Agent008, N)
			for k := range cube[i][j] {
				akt := new(Agent008)
				akt.i = i
				akt.j = j
				akt.k = k
				akt.a = make(chan int)
				akt.b = make(chan int)
				akt.c = make(chan int)
				cube[i][j][k] = akt
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				go func(i, j, k int) {
					a := 0
					b := 0
					c := 0
					if i > 0 {
						a = <-cube[i][j][k].a
					}
					if j > 0 {
						b = <-cube[i][j][k].b
					}
					if k > 0 {
						c = <-cube[i][j][k].c
					}
					val := a + b + c
					if i == 0 && j == 0 && k == 0 {
						val = 1
					}
					fmt.Println(i, j, k, val)
					if i+1 < N {
						go func() { cube[i+1][j][k].a <- val }()
					}
					if j+1 < N {
						go func() { cube[i][j+1][k].b <- val }()
					}
					if k+1 < N {
						go func() { cube[i][j][k+1].c <- val }()
					}
				}(i, j, k)
			}
		}
	}
	time.Sleep(10000000000000)
}
