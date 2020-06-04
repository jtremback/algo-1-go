package main

type RarrayStack struct {
	N   int
	S   []string
	Ops int
}

func NewRarrayStack() RarrayStack {
	return RarrayStack{N: 0, S: make([]string, 2)}
}

func (r *RarrayStack) isEmpty() bool {
	r.Ops++
	return r.N == 0
}

func (r *RarrayStack) resize(capacity int) {
	r.Ops++
	copy := make([]string, capacity)
	for i := 0; i < r.N; i++ {
		r.Ops++
		copy[i] = r.S[i]
	}
	r.S = copy
}

func (r *RarrayStack) push(item string) {
	r.Ops++
	if r.N == len(r.S) {
		r.resize(r.N * 2)
	}

	r.S[r.N] = item
	r.N = r.N + 1
}

func (r *RarrayStack) pop() string {
	r.Ops++
	if r.isEmpty() {
		return ""
	}

	if r.N > 0 && r.N == (len(r.S)/4) {
		r.resize(len(r.S) / 2)
	}

	r.N = r.N - 1
	item := r.S[r.N]
	return item
}
