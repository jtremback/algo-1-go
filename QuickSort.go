package main

import "fmt"

func Partition(a []int) []int {
	pivot := a[0]
	i := 1
	j := len(a) - 1
	for {
		checkNotGreaterThan(a[1:i], pivot)
		checkNotLessThan(a[j:], pivot)
		if a[i] > pivot {
			item := a[i]
			a[i] = a[j]
			a[j] = item
			j--
		} else {
			i++
		}
	}
}

func checkNotGreaterThan(a []int, pivot int) {
	for i := 0; i < len(a); i++ {
		if a[i] > pivot {
			panic(fmt.Sprint(a[i], " at index ", i, " is greater than ", pivot))
		}
	}
}

func checkNotLessThan(a []int, pivot int) {
	for i := 0; i < len(a); i++ {
		if a[i] < pivot {
			panic(fmt.Sprint(a[i], " at index ", i, " is less than ", pivot))
		}
	}
}
