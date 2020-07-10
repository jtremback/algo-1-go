package main

import "fmt"

// If a[i] is not greater than the pivot, move i up one.
// If a[i] is greater than the pivot, switch it with a[j], and move j down one
// If a[j] is smaller than the pivot, switch it with a[i], and move i up one ?????

func Partition(a []int) []int {
	pivot := a[0]
	i := 1
	j := len(a) - 1
	printArray(a, i, j)
	for {
		if i == j {
			break
		}
		if a[i] > pivot {
			item := a[i]
			a[i] = a[j]
			a[j] = item
			j--
		} else {
			i++
		}
		printArray(a, i, j)
	}

	checkNotGreaterThan(a[1:i], pivot)
	checkNotLessThan(a[j+1:], pivot)

	// Move pivot item to middle by swapping with last item on left
	// Is this even neccessary??
	item := a[i-1]
	a[i-1] = a[0]
	a[0] = item

	return a
}

func printArray(a []int, i int, j int) {
	fmt.Printf("%+v i: %v, j: %v\n", a, i, j)
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
