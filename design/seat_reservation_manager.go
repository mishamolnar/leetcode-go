package design

import "container/heap"

type SeatManager struct {
	queue *IntHeap
}

func Constructor(n int) SeatManager {
	p := new(SeatManager)
	p.queue = &IntHeap{}
	for i := 0; i < n; i++ {
		p.queue.Push(i)
	}
	heap.Init(p.queue)
	return *p
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.queue).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.queue, seatNumber)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
