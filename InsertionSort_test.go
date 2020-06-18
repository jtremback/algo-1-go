package main

import (
	"testing"
)

type Array []int

func (a Array) Swap(i int, j int) {
	b := a[i]
	c := a[j]
	a[i] = c
	a[j] = b
}

func (a Array) Compare(i int, j int) int {
	if a[i] > a[j] {
		return -1
	}
	if a[i] < a[j] {
		return 1
	}
	return 0
}

func (a Array) Len() int {
	return len(a)
}

func TestInsertionSort(t *testing.T) {
	a := Array{9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(a)
	println(a)
}
