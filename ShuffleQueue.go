package main

// TODO: Decide if enqueue or dequeue should be random
import "math/rand"

type ShuffleQueue struct {
	n     int
	array []*string
}

func NewShuffleQueue() ShuffleQueue {
	return ShuffleQueue{n: 0, array: make([]*string, 2)}
}

func (q *ShuffleQueue) resize(capacity int) {
	copy := make([]*string, capacity)
	for i := 0; i < q.n; i++ {
		copy[i] = q.array[i]
	}
	q.array = copy
}

// This enqueues randomly
func (q *ShuffleQueue) Enqueue(item string) {
	// If we are running out of room, resize the array
	if q.n == len(q.array) {
		q.resize(q.n * 2)
	}

	// This is a special case since there is only one option
	if q.n == 0 {
		q.array[q.n] = &item
		q.n++
		return
	}

	// Choose random index
	randn := rand.Intn(q.n)
	// Place item from that index at end of queue
	q.array[q.n] = q.array[randn]
	// Place new item in random spot
	q.array[randn] = &item
	q.n++
}

// This dequeues non-randomly
func (q *ShuffleQueue) Dequeue() string {
	if q.n == 0 {
		return ""
	}

	q.n--
	item := q.array[q.n]

	// if the queue is taking up 1/4 of the backing array, cut the
	// backing array in half.
	if len(q.array) > 2 && q.n == len(q.array)/4 {
		q.resize(len(q.array) / 2)
	}

	return *item
}

// This dequeues randomly
func (q *ShuffleQueue) DequeueR() string {
	if q.n == 0 {
		return ""
	}

	// Choose random index
	randn := rand.Intn(q.n)
	// Select item from that position in queue
	item := q.array[randn]
	// Decrement n to the position that has the last item
	q.n--
	// Place last item in random position
	q.array[randn] = q.array[q.n]

	// Clean up last item
	q.array[q.n] = nil

	// if the queue is taking up 1/4 of the backing array, cut the
	// backing array in half.
	if q.n == len(q.array)/4 {
		q.resize(len(q.array) / 2)
	}

	return *item
}

func (q *ShuffleQueue) Iterate(cond func() bool, cb func(string)) {
	copy := make([]*string, q.n)
	for i := 0; i < q.n; i++ {
		copy[i] = q.array[i]
	}

	q2 := ShuffleQueue{
		array: copy,
		n:     q.n,
	}

	for i := 0; i < q.n; i++ {
		cb(*q2.array[i])
	}
}
