package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	Tasks       []Task
	TaskChannel chan Task
	wg          sync.WaitGroup
	workerSize  int
}

// 1
func (w *WorkerPool) worker(name int) {
	for task := range w.TaskChannel {
		fmt.Printf("Worker %d starting task\n", name)
		task.Process()
		w.wg.Done()
	}
}

func (w *WorkerPool) Start() {
	w.TaskChannel = make(chan Task)
	w.wg = sync.WaitGroup{}
	w.wg.Add(len(w.Tasks))

	for i := 0; i < w.workerSize; i++ {
		go w.worker(i)
	}

	for _, task := range w.Tasks {
		w.TaskChannel <- task
	}

}
