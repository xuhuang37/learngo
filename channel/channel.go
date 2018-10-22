package main

import (
		"time"
	"fmt"
)
func worker(id int, c chan int) {
	for n:= range c {
		fmt.Printf("Worker %d recieved %c\n",id,n)
	}
}
func createWorker(id int)chan<-int {
	ch:=make(chan int)
	go worker(id,ch)
	return ch
}

func chanDemo()  {
	var channels [10]chan<-int
	for i:=0;i<10 ;i++  {
		channels[i] = createWorker(i)
	}
	for i:=0;i<10 ;i++  {
		channels[i]<- 'a'+ i
	}
	for i:=0;i<10 ;i++  {
		channels[i]<- 'A'+ i
	}
	time.Sleep(time.Millisecond)
}

func buffedChannel()  {
	c:=make(chan int, 3)
	go worker(0,c)
	c<-'a'
	c<-'b'
	c<-'c'
	time.Sleep(time.Millisecond)
}
func channelClose()  {
	c:=make(chan int, 3)
	go worker(0,c)
	c<-'a'
	c<-'b'
	c<-'c'
	c<-'d'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//fmt.Println("Channel is the first-class citizen")
	//chanDemo()
	//fmt.Println("Buffer channel")
	//buffedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
