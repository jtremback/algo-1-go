package main

import (
	"fmt"
	"testing"
)

func TestQSPartition(t *testing.T) {
	a := []int{5, 3, 1, 4, 5, 8, 2, 2, 1, 5, 4, 6}
	fmt.Printf("%+v\n", Partition(a))
}

func TestFuzzQSPartition(t *testing.T) {
	trials := []int{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50}
	for i := 0; i < len(trials); i++ {
		arr := generateRandomSlice(trials[i], 10)
		Partition(arr)
		print("\n-------------\n")
		// if !isSorted(arr) {
		// 	panic("not sorted")
		// }
		// fmt.Println("Length: ", trials[i], "Ops: ", ops, "Ratio: ", ops/trials[i])
	}
}
