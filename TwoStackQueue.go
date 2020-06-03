package main

type TwoStackQueue struct {
	Dequeued int
	Forward  RarrayStack
	Reverse  RarrayStack
}

func (q *TwoStackQueue) enqueue(item string) {
	q.Forward.push(item)
}

func (q *TwoStackQueue) dequeue() string {
	item := q.Reverse.pop()

	// If the reverse stack is less than half the size of the forward stack
	// Load the whole forward stack onto the reverse stack
	if q.Forward.N > q.Reverse.N*2 {
		for i := 0; i < q.Forward.N; i++ {
			q.Reverse = NewRarrayStack()
			q.Reverse.push(q.Forward.pop())
		}

		// This is the length of the forward stack minus the number of items that
		// have been popped.
		remaining := q.Forward.N - q.Dequeued
		// Doing this cuts the popped items off the front of the
		// forward queue.
		for i := 0; i < remaining; i++ {
			q.Forward = NewRarrayStack()
			q.Forward.push(q.Reverse.pop())
		}
		q.Dequeued = 0
	}

	q.Dequeued++
	return item
}
