// Večerajúci filozófovia

//5ti filozofovia, 5 vidliciek, na jedenie potrebuju 2
// filozof - 3 stavy: eat, think, hungry
//https://en.wikipedia.org/wiki/Dining_philosophers_problem

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	N       = 5            //pocet filozofov
	filozof []*Philosopher // pole vsetych filozofov
	forks   []*fork
)

type fork struct {
	mux sync.Mutex
}

type Philosopher struct { // filozof definovany ako objekt
	id                  int // id filozofa
	leftFork, rightFork *fork
}

func newPhilosopher(id int) *Philosopher {
	f := &Philosopher{id, forks[id], forks[(id+1)%N]}
	filozof = append(filozof, f)
	return f
}

func (f Philosopher) write(action string) {
	fmt.Printf("Filozof %d %s\n", f.id, action)
}

// zivot filozofa

//1.DEADLOCK

//Zivot filozofa:
//-premysla
//-zoberie lavu vidlicku
//-premysla
//-zoberie pravu vidlicku
//-je
//-odlozi pravu a lavu vidlicku a opakuje svoj zivotny cyklus
func (f Philosopher) run() {
	go func() {
		for {
			f.write("premysla")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			//f.write("je hladny")
			f.write("je hladny - ziada lavu vidlicku")
			f.leftFork.mux.Lock()

			f.write("premysla")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			f.write("je hladny - ziada pravu vidlicku")
			f.rightFork.mux.Lock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			f.write("je")

			f.write("polozil pravu vidlicku")
			f.rightFork.mux.Unlock()
			f.write("polozil lavu vidlicku")
			f.leftFork.mux.Unlock()

		}
	}()
}

//2.do nekonečna

//Zivot filozofa:
//-premysla
//-zoberie vidlicku s nizsim cislom
//-zoberie vidlicku s vyssim cislom
//-je
//-odlozi lavu a pravu vidlicku a opakuje svoj zivotny cyklus
func (f Philosopher) run2() {
	go func() {
		for {
			f.write("premysla")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			f.write("je hladny")
			if f.id != N-1 { //zoberiem prvu vidlicku s nizsim cislom (je to lava okrem filozofa s id=N-1)
				f.write("ziada lavu vidlicku")
				f.leftFork.mux.Lock()
				f.write("ziada pravu vidlicku")
				f.rightFork.mux.Lock()
			} else {
				f.write("ziada pravu vidlicku")
				f.rightFork.mux.Lock()
				f.write("ziada lavu vidlicku")
				f.leftFork.mux.Lock()
			}
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			f.write("je")

			f.write("polozil lavu vidlicku")
			f.leftFork.mux.Unlock()
			f.write("polozil pravu vidlicku")
			f.rightFork.mux.Unlock()
		}
	}()
}

//---------------------------------------------------------------
func main() {
	rand.Seed(time.Now().UnixNano())
	forks = make([]*fork, N)
	for i := 0; i < N; i++ {
		forks[i] = new(fork)
	}
	/////
	filozof = make([]*Philosopher, N)
	for i := 0; i < N; i++ {
		filozof[i] = newPhilosopher(i)
	}
	fmt.Println("-----")
	for i := 0; i < N; i++ {
		//filozof[i].run() //deadlock
		filozof[i].run2() //runs forever
	}
	finito := make(chan int, 3)
	<-finito
	fmt.Println("finito")
}
