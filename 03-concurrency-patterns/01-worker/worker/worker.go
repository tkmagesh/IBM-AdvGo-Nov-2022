package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	workQueue chan Work
	wg        sync.WaitGroup
}

func New(workersCount int) *Worker {
	w := &Worker{
		workQueue: make(chan Work),
	}
	w.wg.Add(workersCount)
	for idx := 0; idx < workersCount; idx++ {
		go func(id int) {
			fmt.Printf("Worker %d starts...\n", id)
			for wk := range w.workQueue {
				wk.Task()
			}
			w.wg.Done()
		}(idx)
	}
	return w
}

func (w *Worker) Run(work Work) {
	w.workQueue <- work
}

func (w *Worker) Shutdown() {
	close(w.workQueue)
	w.wg.Wait()
	fmt.Println("Worker shutdown completed")
}
