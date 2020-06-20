package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{11, 10, 1, 9, 8, 7, 4, 6, 5, 2, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Printf("%#v\n", arr)
}
