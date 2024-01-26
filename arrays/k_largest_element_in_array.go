package arrays

import (
	"container/heap"
)

// case [3,2,1,5,6,4] // should return 5
func FindKthLargest(nums []int, k int) int {
	minHeap := MinHeap(nums)
	heap.Init(&minHeap)
    for i := 1; i < k; i++ {
		heap.Pop(&minHeap)
	}
	return heap.Pop(&minHeap).(int)
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] >= h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
