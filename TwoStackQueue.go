package main

import "fmt"

type TwoStackQueue struct {
	Dequeued int
	Forward  RarrayStack
	Reverse  RarrayStack
	Ops      int
}

func NewTwoStackQueue() TwoStackQueue {
	return TwoStackQueue{
		Dequeued: 0,
		Forward:  NewRarrayStack(),
		Reverse:  NewRarrayStack(),
		Ops:      0,
	}
}

func (q *TwoStackQueue) printStats(location string) {
	println(location)
	println("Dequeued: ", q.Dequeued)
	println("Ops: ", q.Ops)
	fmt.Printf("Forward: %v", q.Forward.S)
	println(" len: ", q.Forward.N)
	fmt.Printf("Reverse: %v", q.Reverse.S)
	println(" len: ", q.Reverse.N)
	println("-------------")
}

func (q *TwoStackQueue) enqueue(item string) {
	q.printStats("enqueue start")
	q.Ops++
	q.Forward.push(item)
}

func (q *TwoStackQueue) dequeue() string {
	q.printStats("dequeue start")
	q.Ops++

	// If the reverse stack is less than half the size of the forward stack
	// Load the whole forward stack onto the reverse stack
	if q.Forward.N > q.Reverse.N*2 {
		// This is the length of the forward stack minus the number of items that
		// have been popped.
		remaining := q.Forward.N - q.Dequeued
		println("remaining", remaining)
		q.Reverse = NewRarrayStack()
		tempReverse := NewRarrayStack()
		for i := 0; i < remaining; i++ {
			q.Ops++
			item := q.Forward.pop()
			q.Reverse.push(item)
			// tempReverse is used to repopulate the forward stack
			tempReverse.push(item)
		}

		// Doing this cuts the popped items off the front of the
		// forward queue.
		q.Forward = NewRarrayStack()
		for i := 0; i < remaining; i++ {
			q.Ops++
			item := tempReverse.pop()
			q.Forward.push(item)
		}

		// Reset the counter
		q.Dequeued = 0
	}

	item := q.Reverse.pop()
	q.Dequeued++
	return item
}
