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

	if len(list1) == 0 || len(list2) == 0 {
		return result
	}

	h := &MinPairHeap{}

	for i := 0; i < minInt(len(list1), k); i++ {
		heap.Push(h, pair{v1: list1[i], v2: list2[0]})
	}

	i, j, jk := 0, 0, 1
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		result = append(result, []int{p.v1, p.v2})

		if len(result) == k {
			break
		}

		j++
		if j == len(list2) {
			j = jk
			jk++
			i++
		}

		if i > len(list1)-1 || j > len(list2)-1 {
			continue
		}

		heap.Push(h, pair{v1: list1[i], v2: list2[j]})
	}

	return result
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
