package main

// arr: |...| |  | |  |...|
//          i    j    n
func Merge(arr []int, aux []int, lo int, mid int, hi int) int {
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
		} else if arr[j] < arr[i] {
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

func MergeSort(arr []int) int {
	ops := 0
	l := len(arr)
	aux := make([]int, l)
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

			ops += Merge(arr, aux, lo, mid, hi)
			lo = hi
		}
		interval = interval * 2
	}

	return ops
}
