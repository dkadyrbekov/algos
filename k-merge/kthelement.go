package k_merge

import "container/heap"

type elem struct {
	val       int
	listIndex int
	index     int
}
type MinValHeap []elem

func (h MinValHeap) Len() int           { return len(h) }
func (h MinValHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinValHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinValHeap) Push(x interface{}) {
	*h = append(*h, x.(elem))
}

func (h *MinValHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestNumber(lists [][]int, k int) int {
	n := len(lists)
	if n == 0 {
		return 0
	}

	vals := &MinValHeap{}
	for i := 0; i < len(lists); i++ {
		if len(lists[i]) > 0 {
			heap.Push(vals, elem{
				val:       lists[i][0],
				listIndex: i,
				index:     0,
			})
		}
	}

	minKthElem := 0

	for i := 0; i < k && vals.Len() > 0; i++ {
		e := heap.Pop(vals).(elem)
		minKthElem = e.val

		if e.index+1 < len(lists[e.listIndex]) {
			heap.Push(vals, elem{
				val:       lists[e.listIndex][e.index+1],
				listIndex: e.listIndex,
				index:     e.index + 1,
			})
		}
	}

	return minKthElem
}
