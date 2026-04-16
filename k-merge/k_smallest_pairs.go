package k_merge

import "container/heap"

type pair struct {
	v1 int
	v2 int
}
type MinPairHeap []pair

func (h MinPairHeap) Len() int           { return len(h) }
func (h MinPairHeap) Less(i, j int) bool { return h[i].v1+h[i].v2 < h[j].v1+h[j].v2 } // min-heap
func (h MinPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinPairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinPairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(list1 []int, list2 []int, k int) [][]int {
	result := make([][]int, 0)

	h := &MinPairHeap{}

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			heap.Push(h, pair{v1: list1[i], v2: list2[j]})
		}
	}

	for i := 0; i < k && h.Len() > 0; i++ {
		p := heap.Pop(h).(pair)
		result = append(result, []int{p.v1, p.v2})
	}

	return result
}
