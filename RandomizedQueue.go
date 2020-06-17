package main

import "math/rand"

type RandomizedQueue struct {
	n     int
	count int
	array []*string
	ops   int
}

// construct an empty randomized queue
func NewRandomizedQueue() RandomizedQueue {
	return RandomizedQueue{
		array: make([]*string, 2),
	}
}

// is the randomized queue empty?
func (r *RandomizedQueue) IsEmpty() bool {
	return r.count == 0
}

// return the number of items on the randomized queue
func (r *RandomizedQueue) Size() int {
	return r.count
}

// add the item
func (r *RandomizedQueue) Enqueue(item string) {
	r.ops++
	// defrag if we've run out of space for new items
	if r.n == len(r.array) {
		newArraySize := r.count * 2
		// Make sure that the size is not zero because that causes a bug
		if newArraySize == 0 {
			newArraySize = 2
		}
		r.array = r.defrag(newArraySize)
		r.n = r.count
	}

	r.array[r.n] = &item
	r.n++
	r.count++
}

// remove and return a random item
func (r *RandomizedQueue) Dequeue() string {
	r.ops++
	if r.count == 0 {
		return ""
	}
	item, index := r.getItem()
	r.array[index] = nil
	r.count--

	if (r.count != 0) && ((r.n / r.count) >= 4) {
		// defrag if underutilization ratio is over 4
		r.array = r.defrag(r.count * 2)
		r.n = r.count
	}

	return item
}

// return a random item (but do not remove it)
func (r *RandomizedQueue) Sample() string {
	if r.count == 0 {
		return ""
	}
	item, _ := r.getItem()
	return item
}

func (r *RandomizedQueue) Iterate(cond func() bool, cb func(string)) {
	r.ops++
	// iterator copies the entire data structure and pops all the items off it
	// defrag copies items so we use it. the copied queue is at capacity instead of double
	// because we know we will only ever pop off it
	arr2 := r.defrag(r.count)
	q2 := RandomizedQueue{n: r.count, count: r.count, array: arr2}
	// current := d.first
	for q2.count != 0 && cond() {
		cb(q2.Dequeue())
	}
}

// remove and return a random item
func (r *RandomizedQueue) getItem() (string, int) {
	for {
		r.ops++
		// Get random item from between 0 and N
		randn := rand.Intn(r.n)
		item := r.array[randn]
		// Only leave loop once a non-nil item has been found
		if item != nil {
			return *item, randn
		}
	}
}

// How to defrag:
//   - Make an array of size count * n
//   - Iterate over the old array from 0-N, skipping nils
//   - Copy each non-nil item to the new array
func (r *RandomizedQueue) defrag(capacity int) []*string {
	newArray := make([]*string, capacity)
	j := 0

	for i := 0; i < r.n; i++ {
		r.ops++
		item := r.array[i]
		if item != nil {
			newArray[j] = item
			j++
		}
	}

	return newArray
}
