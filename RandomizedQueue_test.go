package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomizedQueue(t *testing.T) {
	r := NewRandomizedQueue()
	r.Enqueue("to")
	r.Enqueue("be")
	r.Enqueue("or")
	r.Enqueue("not")
	r.Enqueue("to")
	r.Enqueue("be")

	r.Iterate(func() bool {
		return true
	}, func(item string) {
		print(item, ", ")
	})
}

func TestFuzzRandomizedQueue(t *testing.T) {
	r := NewRandomizedQueue()
	iterations := 100
	for i := 0; i < iterations; i++ {
		randn := rand.Intn(2)
		if randn == 0 {
			item := fmt.Sprintf("%f", rand.Float64())
			r.Enqueue(item)
			// println("Enqueue(%f)", item)
		} else if randn == 1 {
			// item :=
			r.Dequeue()
			// println("Dequeue() %f", item)
		} else {
			panic("not enough choices")
		}

		println(r.count)
	}
	println("ops", r.ops)
}
