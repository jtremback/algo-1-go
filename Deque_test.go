package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDeque(t *testing.T) {
	d := Deque{}
	d.addLast("to")
	d.addLast("be")
	d.addLast("or")
	d.addFirst("is")

	i := 0
	d.iterate(func() bool {
		i = i + 1
		return i <= 2
	}, func(item string) {
		print(item, ", ")
	})
}

func TestFuzzDeque(t *testing.T) {
	d := Deque{}
	iterations := 10000
	for i := 0; i < iterations; i++ {
		randn := rand.Intn(3)
		if randn == 0 {
			item := fmt.Sprintf("%f", rand.Float64())
			d.addLast(item)
			println("addLast(%f)", item)
		} else if randn == 1 {
			item := fmt.Sprintf("%f", rand.Float64())
			d.addFirst(item)
			println("addFirst(%f)", item)
		} else if randn == 2 {
			item := d.removeFirst()
			println("removeFirst() %f", item)
		} else if randn == 3 {
			item := d.removeLast()
			println("removeLast() %f", item)
		} else {
			panic("not enough choices")
		}

		println(d.count)
	}
}
