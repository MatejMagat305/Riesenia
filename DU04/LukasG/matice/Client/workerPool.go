package main

import (
	"time"
)

type WorkerPool struct {
	workers []*Worker
	tasks   []*Task
}

func Extract(address string ) (host string, port string) {
	flag := false
	for _, c := range address {
		if c == ':' {
			flag = true
		} else if flag {
			port += string(c)
		} else {
			host += string(c)
		}
	}
	return host, port
}

func NewWorkerPool(Brigada []string) *WorkerPool {
	workers := make([]*Worker, len(Brigada))
	for i, address := range Brigada {
		host, port := Extract(address)
		workers[i] = NewWorker(host, port)
	}
	tasks := make([]*Task, 0)
	return &WorkerPool{workers, tasks}
}

func (wp *WorkerPool) addTask(t *Task) {
	wp.tasks = append(wp.tasks, t)
}

func(wp *WorkerPool) getTimes(chanRes chan Result) {
	task := &Task{-1, Vektor{}, -1, Vektor{}}
	for _, worker := range wp.workers {
		worker.busy = true
		worker.doit(task, chanRes)
	}
}

func (wp *WorkerPool) mainLoop(chanRes chan Result, chanTime chan time.Duration) {
	counter := make(chan bool, len(wp.workers))
	workTime := time.Duration(0)
	for _, task := range wp.tasks {
		counter <- true
		start := time.Now()
		for _, worker := range wp.workers {
			if worker.busy == false {
				worker.busy = true
				go func(t *Task, w* Worker) {
					w.doit(t, chanRes)
					_ = <- counter
				}(task, worker)
				break
			}
		}
		workTime += time.Since(start)
	}
	chanTime <- workTime
}



