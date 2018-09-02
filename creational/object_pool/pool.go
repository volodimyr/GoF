package main

import (
	"fmt"
)

/*
	The object pool pattern is a software creational design pattern that uses a set of initialized objects kept ready to use – a "pool"–
	rather than allocating and destroying them on demand.
	A client of the pool will request an object from the pool and perform operations on the returned object.
	When the client has finished, it returns the object to the pool rather than destroying it; this can be done manually or automatically.
*/

type Task struct {
}

func (t Task) Do() {
	fmt.Println("Doing...")
}

type Pool chan *Task

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Task)
	}

	return &p
}

func main() {
	p := New(3)

	for {
		select {
		case task := <-*p:
			task.Do()

		default:
			return
		}
	}
}
