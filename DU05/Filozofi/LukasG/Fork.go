package philosophers

import "sync"

type Fork struct {
	id        int
	mutex     sync.Mutex
	available bool
}

func newFork(id int) *Fork {
	return &Fork{id, sync.Mutex{}, true}
}

func (fork *Fork) take() {
	fork.available = false
	fork.mutex.Lock()
}

func (fork *Fork) put() {
	fork.available = true
	fork.mutex.Unlock()
}