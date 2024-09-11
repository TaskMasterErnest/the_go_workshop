/*
**
We will calculate a sum using several workers.
A worker is essentialluy a function and we will organize these workers in a single struct
**
*/

package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in        chan int
	out       chan int
	subWorker int
	mtx       *sync.Mutex
}

// create methods
func (w *Worker) readThem() {
	// increment subworker number
	w.subWorker++

	go func() {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial

		w.mtx.Lock()
		w.subWorker--
		if w.subWorker == 0 {
			close(w.out)
		}
		w.mtx.Unlock()
	}()
}

func (w *Worker) getResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := range w.out {
			total += i
		}
		wg.Done()
	}()

	wg.Wait()
	return total
}

func main() {
	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	workerNum := 10
	out := make(chan int)
	worker := Worker{in: in, out: out, mtx: mtx}

	for i := 1; i <= workerNum; i++ {
		worker.readThem()
	}

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)

	result := worker.getResult()
	fmt.Println(result)
}
