package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second / 15)
		fmt.Printf("Worker %d recieved %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	ch := make(chan int)
	go worker(id, ch)
	return ch
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				100 * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	var values []int
	tm := time.After(5 * time.Second)
	tick := time.Tick(2 * time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tick:
			fmt.Printf("queue lens=%d\n", len(values))
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tm:
			fmt.Println("bye bye~")
			return
		}
	}
}
