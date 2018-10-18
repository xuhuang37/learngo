package main

import (
	"learngo/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	isEm := q.IsEmpty()
	fmt.Println(isEm)
	q.Push("aaa")
}
