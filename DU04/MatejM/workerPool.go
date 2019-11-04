package main

import "sync"

type WorkerPool struct {
	workers []*Worker
	tasks   []*Task
	mutex sync.Mutex
}

func NewWorkerPool(Brigada []string) *WorkerPool {
	workers := make([]*Worker, len(Brigada))
	for i, port := range Brigada {
		workers[i] = NewWorker("127.0.0.1", port) // :8000,
		 	}
	tasks := make([]*Task, 0)
	return &WorkerPool{workers, tasks, sync.Mutex{}}
}
func (wp *WorkerPool) addTask(t *Task) {
	wp.tasks = append(wp.tasks, t) // pridaj na koniec fronty
}
func (wp *WorkerPool) mainLoop(chainRes chan Result, N int) {
	// nejaku filozofiu pre spracovanie taskov vo worker poole
	nn :=N*N
	for i:=0;i<nn ;i++  {
		go func(){chainRes<-wp.dajDalsiTask(nn)}()
	}
}
func (wp *WorkerPool)dajDalsiTask(n int) Result {
	for  {
		wp.mutex.Lock()
		if len(wp.tasks)>0 {
			prvok := wp.popTask()
			wp.mutex.Unlock()
			for {
				for i:=0;i<len(wp.workers) ;i++ {
					wp.workers[i].mux.Lock()
					t :=false
					if !wp.workers[i].busy {
						wp.workers[i].busy=true
						t=true
					}
					wp.workers[i].mux.Unlock()
					if t {
						n :=wp.workers[i].doit(prvok)
						println(n.I1, "   ", n.I2,"   ",n.Skalar)
						return n
					}
				}
			}
		}else{wp.mutex.Unlock()}

	}
}

func (wp *WorkerPool) popTask() *Task {
	vysl := wp.tasks[0]
	wp.tasks=wp.tasks[1:]
	return vysl
}


