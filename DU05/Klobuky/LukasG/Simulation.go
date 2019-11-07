package hats

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Simulation struct {
	N             int
	configuration map[int]Hat
	agents        []*Agent
	mutex         *sync.WaitGroup
}

func (simulation *Simulation) tell(agent *Agent) (int, int) {
	blacks := 0
	whites := 0
	for _, otherAgent := range simulation.agents {
		if otherAgent.id == agent.id {
			continue
		}
		if simulation.configuration[otherAgent.id].color == true {
			blacks += 1
		} else {
			whites += 1
		}
	}
	return blacks, whites
}

func (simulation *Simulation) printConfiguration() {
	fmt.Println("CONFIGURATION:")
	for id := 0; id < simulation.N; id++ {
		fmt.Printf("Agent id: %d, hat: %s \n", id, simulation.configuration[id].toString())
	}
	fmt.Println("-----")
}

func (simulation *Simulation) check() {
	faults := 0
	fmt.Println("VERIFICATION:")
	for id := 0; id < simulation.N; id++ {
		truth := simulation.configuration[id]
		opinion := simulation.agents[id].opinion
		fmt.Printf("Agent id: %d, hat: %s, opinion: %s, correct: %v \n",
			id, truth.toString(), opinion.toString(), opinion.color == truth.color)
		if truth != opinion {
			faults += 1
		}
	}
	fmt.Printf("Total faults: %d \n", faults)
	fmt.Println("-----")
}

func (simulation *Simulation) run() {
	fmt.Println("SIMULATION:")
	for _, agent := range simulation.agents {
		agent.run()
	}
	simulation.mutex.Wait()
	time.Sleep(time.Second) //to make sure that all routines ended
	fmt.Println("-----")
}

func generateConfiguration(N int) map[int]Hat {
	rand.Seed(time.Now().Unix())
	m := make(map[int]Hat)
	blacks := N - 1
	for id := 0; id < N; id++ {
		if rand.Intn(2) == 1 && blacks > 0 {
			m[id] = Hat{true}
			blacks -= 1
		} else {
			m[id] = Hat{false}
		}
	}
	return m
}

func newSimulation(N int) *Simulation {
	assignment := generateConfiguration(N)
	mutex := new(sync.WaitGroup)
	mutex.Add(N)
	var agents []*Agent
	simulation := &Simulation{N, assignment, agents, mutex}
	for i := 0; i < N; i++ {
		agent := newAgent(i, simulation)
		agents = append(agents, agent)
	}
	simulation.agents = agents
	return simulation
}


