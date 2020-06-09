package main

type Deque struct {
	first *Node
	last  *Node
	count int
}

type Node struct {
	item string
	next *Node
}

// is the deque empty?
func (d *Deque) isEmpty() bool {
	// assert that both first and last must be nil if one is
	if (d.first == nil) != (d.last == nil) {
		panic("data structure irregularity")
	}
	return d.first == nil
}

// return the number of items on the deque
func (d *Deque) size() int {
	return d.count
}

// add the item to the front
func (d *Deque) addFirst(item string) {
	node := Node{
		item: item,
		next: nil,
	}

	if d.isEmpty() {
		d.first = &node
		d.last = &node
	} else {
		node.next = d.first
		d.first = &node
	}

	d.count++
}

// add the item to the back
func (d *Deque) addLast(item string) {
	// println("addLast")
	node := Node{
		item: item,
		next: nil,
	}

	if d.isEmpty() {
		d.first = &node
		d.last = &node
	} else {
		d.last.next = &node
		d.last = &node
	}

	d.count++
}

func (d *Deque) iterate(cond func() bool, cb func(string)) {
	current := d.first
	for current != nil && cond() {
		cb(current.item)
		current = current.next
	}
}

// remove and return the item from the front
func (d *Deque) removeFirst() string {
	// println("removeFirst")
	if d.isEmpty() {
		return ""
	}
	d.count--
	node := d.first
	d.first = node.next
	// If this empties the list, set last to nil as well
	if d.first == nil {
		d.last = nil
	}
	return node.item
}

// remove and return the item from the back
func (d *Deque) removeLast() string {
	if d.isEmpty() {
		return ""
	}
	d.count--
	node := d.last
	d.first = node.next
	return node.item
}
