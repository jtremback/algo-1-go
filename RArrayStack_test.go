package main

import (
	"testing"
)

func TestRarrayStack(t *testing.T) {
	r := NewRarrayStack()
	r.push("to")
	r.push("be")
	r.push("or")

	println(r.pop())
	println(r.pop())
	println(r.pop())
	println(r.pop())

	r.push("to")
	r.push("be")
	r.push("or")

	println(r.pop())
	println(r.pop())
	println(r.pop())
	println(r.pop())

	println("Ops", r.Ops)
}

func TestRarrayStackPerf(t *testing.T) {
	r := NewRarrayStack()
	for i := 0; i <= 100000; i++ {
		r.push("string")
	}

	for i := 0; i <= 99999; i++ {
		r.pop()
	}

	println("Ops", r.Ops)
}
