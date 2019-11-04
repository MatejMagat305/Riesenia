package main

import "strings"
import "fmt"

type WorkerPool struct {
	workers []*Worker
	tasks   []*Task
}

func NewWorkerPool(Brigada []string) *WorkerPool {
	workers := make([]*Worker, len(Brigada))
	for i, port := range Brigada {
		h := strings.Split(port, ":")
		p := port
		host := "localhost"
		if len(h) > 1 {
			host = h[0]
			p = h[1]
		} else {
			p = h[0]
		}
		workers[i] = NewWorker(host, p) // :8000,
	}
	tasks := make([]*Task, 0)
	return &WorkerPool{workers, tasks}
}

func (wp *WorkerPool) addTask(t *Task) {
	wp.tasks = append(wp.tasks, t) // pridaj na koniec fronty
}

func (wp *WorkerPool) mainLoop(chainRes chan Result) {
	// nejaku filozofiu pre spracovanie taskov vo worker poole
	moj := make(chan Result, len(wp.workers))
	treba := len(wp.tasks)
	ako_vyplatit := make([]int, len(wp.workers))
	for i := range ako_vyplatit {
		ako_vyplatit[i] = 0
	}
	for i, v := range wp.workers {
		if len(wp.tasks) == 0 {
			break
		}
		akt := wp.tasks[0]
		wp.tasks = wp.tasks[1:]
		akt.Wid = i
		go v.doit(akt, moj)
	}
	for treba > 0 {
		akt := <-moj
		ako_vyplatit[akt.Wid]++
		go func() {
			chainRes <- akt
		}()
		if len(wp.tasks) > 0 {
			a := wp.tasks[0]
			wp.tasks = wp.tasks[1:]
			a.Wid = akt.Wid
			go wp.workers[a.Wid].doit(a, moj)
		}
		treba--
	}
	fmt.Println("robotnikov treba vyplatit takto podla prace")
	fmt.Println(ako_vyplatit)
}
