package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{11, 10, 1, 9, 8, 7, 4, 6, 5, 2, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Printf("%#v\n", arr)
}

func TestFuzzMergeSort(t *testing.T) {
	trials := []int{10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000}
	for i := 0; i < len(trials); i++ {
		arr := generateRandomSlice(trials[i], 99999999)
		ops := MergeSort(arr)
		if !isSorted(arr) {
			panic("not sorted")
		}
		fmt.Println("Length: ", trials[i], "Ops: ", ops, "Ratio: ", ops/trials[i])
	}
}

func generateRandomSlice(size int, rnge int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(rnge) - rand.Intn(rnge)
	}
	return slice
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		// If the current item is lower than the previous one
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
