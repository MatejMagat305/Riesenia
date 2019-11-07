package philosophers

import (
	"fmt"
	"math/rand"
	"time"
)

type NaivePhilosopher struct {
	id int
	left *Fork
	right *Fork
	sim *Simulation
}

func (philosopher *NaivePhilosopher) run() {
	go func() {
		start := time.Now()
		for {
			fmt.Printf("[%dms] %d is thinking... \n", timeSince(start), philosopher.id)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10) *DELAY))
			fmt.Printf("[%dms] %d trying to pick RIGHT fork... \n", timeSince(start), philosopher.id)
			philosopher.right.take()
			fmt.Printf("[%dms] %d picked RIGHT fork... \n", timeSince(start), philosopher.id)
			time.Sleep(time.Millisecond * time.Duration(3 *DELAY))
			fmt.Printf("[%dms] %d trying to pick LEFT fork... \n", timeSince(start), philosopher.id)
			philosopher.left.take()
			fmt.Printf("[%dms] %d picked LEFT fork and is eating... \n", timeSince(start), philosopher.id)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10) *DELAY))
			philosopher.right.put()
			philosopher.left.put()
			fmt.Printf("[%dms] %d put down both forks... \n", timeSince(start), philosopher.id)
		}
	}()
}

func newNaivePhilosopher(id int, left *Fork, right *Fork, sim *Simulation) *NaivePhilosopher {
	return &NaivePhilosopher{id, left, right, sim}
}