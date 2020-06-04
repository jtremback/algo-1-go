package main

import (
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	q := NewTwoStackQueue()
	q.enqueue("to")
	q.enqueue("be")
	q.enqueue("or")

	println("DEQUEUE: ", q.dequeue())
	println("DEQUEUE: ", q.dequeue())
	println("DEQUEUE: ", q.dequeue())
}

// func TestRarrayStackPerf(t *testing.T) {
// 	r := NewRarrayStack()
// 	for i := 0; i <= 100000; i++ {
// 		r.push("string")
// 	}

// 	for i := 0; i <= 99999; i++ {
// 		r.pop()
// 	}

// 	println("Ops", r.Ops)
// }
