package main

type Sortable interface {
	// < 0 if left is greater than right, 0 if equal, > 0 if left is lesser than right
	Compare(int, int) int
	Swap(int, int)
	Len() int
}

// For each item, swap it backwards through the array until the item before it is not greater than it
func InsertionSort(a Sortable) {
	n := a.Len()
	i := 1
	for i < n {
		j := i
		// if it is smaller than the element before it
		for j > 0 && a.Compare(j, j-1) > 0 {
			// Swap the elements
			a.Swap(j, j-1)
			// and move backwards
			j--
		}
		i++
	}
}
