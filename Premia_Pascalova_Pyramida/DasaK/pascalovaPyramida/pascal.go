//Pascalova pyramída
//https://en.wikipedia.org/wiki/Pascal%27s_pyramid

package main

import (
	"fmt"
	"time"
)

type Agent007 struct {
	a, b, c int      // Agent's coordinates in the grid
	x, y, z chan int // inputs for Agent007
}

type Grid [][][]*Agent007

func main() {
	N := 10
	fmt.Println(N)
	grid := make(Grid, N)
	for i := range grid {
		grid[i] = make([][]*Agent007, N)
		for j := range grid[i] {
			grid[i][j] = make([]*Agent007, N)
			for k := range grid[i][j] {
				grid[i][j][k] = new(Agent007)
				grid[i][j][k].a = i
				grid[i][j][k].b = j
				grid[i][j][k].c = k
				grid[i][j][k].x = make(chan int)
				grid[i][j][k].y = make(chan int)
				grid[i][j][k].z = make(chan int)
			}
		}
	}
	/////vnutorne
	for i := 1; i < N; i++ {
		for j := 1; j < N; j++ {
			for k := 1; k < N; k++ {
				go func(i, j, k int) {
					val1 := <-grid[i][j][k].x
					val2 := <-grid[i][j][k].y
					val3 := <-grid[i][j][k].z
					val := val1 + val2 + val3
					//fmt.Println(i+1+j+1, j+1, val)
					fmt.Println(i, j, k, val)
					//(a,b,c) = (a-1,b,c)+(a,b-1,c)+(a,b,c-1)
					if i+1 < N {
						go func() { grid[i+1][j][k].x <- val }()
					}
					if j+1 < N {
						go func() { grid[i][j+1][k].y <- val }()
					}
					if k+1 < N {
						go func() { grid[i][j][k+1].z <- val }()
					}
				}(i, j, k)
			}
		}
	}
	/////strana
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			go func(i, j int) {
				val1 := <-grid[i][j][0].x
				val2 := <-grid[i][j][0].y
				val := val1 + val2
				fmt.Println(i, j, 0, val)
				if i+1 < N {
					go func() { grid[i+1][j][0].x <- val }()
				}
				if j+1 < N {
					go func() { grid[i][j+1][0].y <- val }()
				}
				go func() { grid[i][j][0+1].z <- val }()
			}(i, j)
			go func(i, j int) {
				val1 := <-grid[i][0][j].x
				val2 := <-grid[i][0][j].z
				val := val1 + val2
				fmt.Println(i, 0, j, val)
				if i+1 < N {
					go func() { grid[i+1][0][j].x <- val }()
				}
				go func() { grid[i][0+1][j].y <- val }()
				if j+1 < N {
					go func() { grid[i][0][j+1].z <- val }()
				}
			}(i, j)
			go func(i, j int) {
				val1 := <-grid[0][i][j].y
				val2 := <-grid[0][i][j].z
				val := val1 + val2
				fmt.Println(0, i, j, val)
				go func() { grid[0+1][i][j].x <- val }()
				if i+1 < N {
					go func() { grid[0][i+1][j].y <- val }()
				}
				if j+1 < N {
					go func() { grid[0][i][j+1].z <- val }()
				}
			}(i, j)
		}
	}
	/////hrana
	for i := 0; i < N; i++ {
		go func(i int) {
			grid[i][0][0].x <- 1
			grid[i][0][0].y <- 0
			grid[i][0][0].z <- 0
		}(i)
	}
	for j := 1; j < N; j++ {
		go func(j int) {
			grid[0][j][0].x <- 0
			grid[0][j][0].y <- 1
			grid[0][j][0].z <- 0
		}(j)
	}
	for k := 1; k < N; k++ {
		go func(k int) {
			grid[0][0][k].x <- 0
			grid[0][0][k].y <- 0
			grid[0][0][k].z <- 1
		}(k)
	}
	time.Sleep(10000)
}
