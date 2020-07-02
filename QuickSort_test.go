package main

import (
	"testing"
)

func TestQSPartition(t *testing.T) {
	a := []int{1, 3, 1, 4, 5, 8, 2, 2, 1, 5, 4, 6}
	println(Partition(a))
}
