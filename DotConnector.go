package main

type Dot struct {
	X float64
	Y float64
}

func isLessThan(d1 Dot, d2 Dot, vertex Dot) bool {
	d1Slope := slopeTo(vertex, d1)
	d2Slope := slopeTo(vertex, d2)

	return d1Slope < d2Slope
}

func DotConnector(dots []Dot) [][]Dot {
	collection := [][]Dot{}
	for i := 0; i < len(dots); i++ {
		vertex := dots[i]
		mergeSort(dots, vertex)
		buffer := []Dot{dots[0]}
		// bufferIndex := 0
		for j := 1; j < len(dots); j++ {
			currentDot := dots[j]

			// if len(buffer) == 0 {
			// 	// If we have an empty buffer, start it
			// 	buffer = append(buffer, dots[j])
			// } else

			if slopeTo(vertex, buffer[len(buffer)-1]) == slopeTo(vertex, currentDot) {
				// If the current dot has the same slope as the last one in the buffer, add it in
				buffer = append(buffer, dots[j])
			} else if len(buffer) >= 4 {
				// If the current dot does not have the same slope and the last sequence is big enough to keep,
				// put the last sequence in the collection and reset the buffer
				collection = addToCollection(collection, buffer)
				buffer = []Dot{dots[j]}
			} else {
				// If the current dot does not have the same slope and the last sequence is big enough to keep, reset the buffer
				buffer = []Dot{dots[j]}
			}
		}
		if len(buffer) >= 4 {
			// If we have looked at all items, and there is a valid sequence in the buffer
			// put the last sequence in the collection
			collection = addToCollection(collection, buffer)
		}
	}
	return collection
}

func addToCollection(collection [][]Dot, buffer []Dot) [][]Dot {
	appendBuffer := make([]Dot, len(buffer))
	copy(appendBuffer, buffer)
	return append(collection, appendBuffer)
}

func slopeTo(a Dot, b Dot) float64 {
	// (y1 − y0) / (x1 − x0)
	return (b.Y - a.Y) / (b.X - a.Y)
}

func mergeSort(arr []Dot, vertex Dot) int {
	ops := 0
	l := len(arr)
	aux := make([]Dot, l)
	interval := 1
	for interval < l {
		ops++
		lo := 0
		for lo < l {
			ops++
			mid := lo + interval
			hi := mid + interval

			if l < hi {
				hi = l
			}

			ops += merge(arr, aux, vertex, lo, mid, hi)
			lo = hi
		}
		interval = interval * 2
	}

	return ops
}

func merge(arr []Dot, aux []Dot, vertex Dot, lo int, mid int, hi int) int {
	ops := 0
	i := lo
	j := mid
	l := hi - lo
	// Iterate while there are still items to move to aux
	for k := 0; k < l; k++ {
		ops++
		if i >= mid {
			// If there are no items remaining on the left, but there are some on the right
			aux[k] = arr[j]
			j++
		} else if j >= hi {
			// If there are no items remaining on the right, but there are some on the left
			aux[k] = arr[i]
			i++
		} else if isLessThan(arr[j], arr[i], vertex) {
			// If there are items remaining on the right and the left, and the item on the right is smaller
			aux[k] = arr[j]
			j++
		} else {
			// If there are items remaining on the right and the left, and the item on the left is smaller or they are equal
			aux[k] = arr[i]
			i++
		}
	}
	// copy items in aux array back into the main array
	for i := 0; i < l; i++ {
		ops++
		arr[lo+i] = aux[i]
	}
	return ops
}
