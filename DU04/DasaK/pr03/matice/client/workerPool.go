package main

import (
	"strings"
)

type WorkerPool struct {
	workers []*Worker
	tasks   []*Task
}

func NewWorkerPool(Brigada []string) *WorkerPool {
	workers := make([]*Worker, len(Brigada))
	for i, port := range Brigada {
		//workers[i] = NewWorker("localhost", port) // :8000,
		x := strings.Split(port, ":")
		if len(x) == 1 {
			workers[i] = NewWorker("localhost", x[0])
		} else {
			workers[i] = NewWorker(x[0], x[1])
		}
	}
	tasks := make([]*Task, 0)
	return &WorkerPool{workers, tasks}
}

func (wp *WorkerPool) addTask(t *Task) {
	wp.tasks = append(wp.tasks, t) // pridaj na koniec fronty
}

func (wp *WorkerPool) mainLoop(chainRes chan Result) {
	// nejaku filozofiu pre spracovanie taskov vo worker poole
	for len(wp.tasks) > 0 {
		for i, _ := range wp.workers {
			task := wp.tasks[0]
			wp.tasks = wp.tasks[1:]
			wp.workers[i].doit(task, chainRes)
		}
	}

}
