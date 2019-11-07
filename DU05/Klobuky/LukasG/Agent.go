package hats

import (
	"fmt"
	"time"
)

type Agent struct {
	id         int
	channel    chan Message
	opinion    Hat
	simulation *Simulation
}

func newAgent(id int, simulation *Simulation) *Agent {
	return &Agent{id, make(chan Message), Hat{false}, simulation}
}

func (agent *Agent) look() (int, int) {
	return agent.simulation.tell(agent)
}

func secondsSince(t time.Time) int {
	return int(time.Duration.Seconds(time.Since(t)))
}

func (agent *Agent) broadcast() {
	agent.simulation.mutex.Done()
	m := Message{agent.id, agent.opinion}
	for _, otherAgent := range agent.simulation.agents {
		if otherAgent.id == agent.id {
			continue
		}
		go func() {
			otherAgent.channel <- m
		}()
	}
}

func (agent *Agent) run() {
	go func() {
		start := time.Now()
		maxBlacks := agent.simulation.N - 1
		for {
			blacks, whites := agent.look()
			fmt.Printf("[%ds] id: %d - I see %d BLACK hats and %d WHITE hats \n",
				secondsSince(start), agent.id, blacks, whites)

			if blacks == maxBlacks {
				agent.opinion = Hat{false}
				fmt.Printf("[%ds] id: %d - I have WHITE hat! \n", secondsSince(start), agent.id)
				agent.broadcast()
				return
			}
			time.Sleep(time.Second)

			wait := time.After(time.Second * 4)
			select {
				case m := <- agent.channel:
					fmt.Printf("[%ds] id: %d - Message from %d, he has %s hat \n",
						secondsSince(start), agent.id, m.id, m.opinion.toString())
					agent.opinion = Hat{true}
					fmt.Printf("[%ds] id: %d - I have BLACK hat! \n", secondsSince(start), agent.id)
					agent.broadcast()
					return
				case <- wait:
					fmt.Printf("[%ds] id: %d - No message, wait 5 seconds \n", secondsSince(start), agent.id)
					maxBlacks -= 1
					time.Sleep(time.Second * 5)
			}
		}
	}()
}
