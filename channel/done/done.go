package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d recieved %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [10] worker

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	//for _,worker := range workers {
	//	<-worker.done
	//}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
	// wait for all working
	//for _,worker := range workers {
	//	<-worker.done
	//}

}

func main() {
	fmt.Println("Channel is the first-class citizen")
	chanDemo()
	fmt.Println("Buffer channel")
	//buffedChannel()
	fmt.Println("Channel close and range")
	//channelClose()
}
