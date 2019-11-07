package philosophers

import (
	"fmt"
	"time"
)

type Philosopher interface {
	run()
}

type Simulation struct {
	philosophers []Philosopher
	forks []*Fork
}

func timeSince(t time.Time) int {
	return int(time.Duration.Milliseconds(time.Since(t)))
}

func newSimulation(N int, fixed bool) *Simulation {
	philosophers := make([]Philosopher, N)
	forks := make([]*Fork, N)
	simulation := &Simulation{philosophers, forks}

	for i := 0; i < N; i++ {
		forks[i] = newFork(i)
	}

	for i := 0; i < N; i++ {
		forkId := (i-1 + len(philosophers)) % len(philosophers)
		if fixed {
			philosopher := newFixedPhilosopher(i, forks[i], forks[forkId], simulation)
			philosophers[i] = philosopher
		} else {
			philosopher := newNaivePhilosopher(i, forks[i], forks[forkId], simulation)
			philosophers[i] = philosopher
		}
	}

	simulation.forks = forks
	simulation.philosophers = philosophers
	return simulation
}

func (simulation *Simulation) availableForks() int {
	count := 0
	for _, fork := range simulation.forks {
		if fork.available {
			count += 1
		}
	}
	return count
}

func (simulation *Simulation) observeForks() {
	start := time.Now()
	for {
		fmt.Printf("########\n[%dms] Forks available: %d \n########\n", timeSince(start), simulation.availableForks())
		time.Sleep(time.Millisecond * time.Duration(5 * DELAY))
	}
}

func (simulation *Simulation) runSimulation() {
	fmt.Println("########")
	fmt.Println("DINNER SIMULATION:")
	for _, philosopher := range simulation.philosophers {
		philosopher.run()
	}
	go simulation.observeForks()
	var input string
	_, _ = fmt.Scanln(&input)
	fmt.Println("SIMULATION STOP")
	fmt.Println("########")
}
