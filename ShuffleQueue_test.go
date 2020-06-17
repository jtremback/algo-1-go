package main

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestShuffleQueue(t *testing.T) {
	q := NewShuffleQueue()
	q.Enqueue("to")
	q.Enqueue("be")
	q.Enqueue("or")
	q.Enqueue("not")
	q.Enqueue("to")
	q.Enqueue("be")

	q.Iterate(func() bool {
		return true
	}, func(item string) {
		print(item, ", ")
	})
}

func TestFuzzShuffleQueue(t *testing.T) {
	r := NewShuffleQueue()
	iterations := 10000000
	start := time.Now()
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
	}
	elapsed := time.Since(start)
	log.Printf("ShuffleQueue took %s", elapsed)
	println(elapsed.Microseconds())
}
