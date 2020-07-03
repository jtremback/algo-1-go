package main

import "fmt"

// If a[i] is not greater than the pivot, move i up one.
// If a[i] is greater than the pivot, switch it with a[j], and move j down one
// If a[j] is smaller than the pivot, switch it with a[i], and move i up one ?????

func Partition(a []int) []int {
	pivot := a[0]
	i := 1
	j := len(a) - 1
	for {
		printArray(a, i, j)
		// checkNotGreaterThan(a[1:i], pivot)
		// checkNotLessThan(a[j:], pivot)
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

func printArray(a []int, i int, j int) {
	fmt.Printf("%+v\n", a)
	pointers := ""
	for k := 0; k < len(a)+2; k++ {
		if k == i {
			pointers += " i"
		} else if k == j {
			pointers += " j"
		} else {
			pointers += "  "
		}

	}
	println(pointers)
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
