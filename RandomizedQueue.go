package main

import "math/rand"

type RandomizedQueue struct {
	n     int
	count int
	array []*string
}

// construct an empty randomized queue
func NewRandomizedQueue() RandomizedQueue {}

// How to defrag:
//   - Make an array of size count * n
//   - Iterate over the old array from 0-N, skipping nils
//   - Copy each non-nil item to the new array

func (r *RandomizedQueue) resize(capacity int) []*string {
	newArray := make([]*string, capacity)
	j := 0

	for i := 0; i < r.n; i++ {
		item := r.array[i]
		if item != nil {
			newArray[j] = item
			j++
		}
	}

	return newArray
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
func (r *RandomizedQueue) Enqueue(item *string) {
	// resize if we've run out of space for new items
	if r.n == len(r.array) {
		r.array = r.resize(r.count * 2)
	}
	// resize/defrag if underutilization ratio is over 4
	if (r.n / r.count) >= 4 {
		r.array = r.resize(r.count * 2)
	}

	r.array[r.n] = item
	r.n++
}

// remove and return a random item
func (r *RandomizedQueue) getItem() (*string, int) {
	for {
		// Get random item from between 0 and N-1
		randn := rand.Intn(r.n - 1)
		item := r.array[randn]
		// Only leave loop once a non-nil item has been found
		if item != nil {
			return item, randn
		}
	}
}

// remove and return a random item
func (r *RandomizedQueue) Dequeue() *string {
	item, index := r.getItem()
	r.array[index] = nil
	return item
}

// return a random item (but do not remove it)
func (r *RandomizedQueue) Sample() *string {
	item, _ := r.getItem()
	return item
}

func (r *RandomizedQueue) Iterate(cond func() bool, cb func(*string)) {
	// iterator copies the entire data structure and pops all the items off it
	// resize copies items so we use it. the copied queue is at capacity
	// because we know we will only pop off it
	arr2 := r.resize(r.count)
	q2 := RandomizedQueue{n: r.count, count: r.count, array: arr2}
	// current := d.first
	for i := 0; i < r.count && cond(); i++ {
		cb(q2.Dequeue())
		// for current != nil && cond() {
		// 	cb(current.item)
		// 	current = current.next
	}
}
