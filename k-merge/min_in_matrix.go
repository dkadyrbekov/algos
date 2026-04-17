package k_merge

import "container/heap"

type rowVal struct {
	val int
	row int
	col int
}
type MinRowValHeap []rowVal

func (h MinRowValHeap) Len() int           { return len(h) }
func (h MinRowValHeap) Less(i, j int) bool { return h[i].val < h[j].val } // min-heap
func (h MinRowValHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinRowValHeap) Push(x interface{}) {
	*h = append(*h, x.(rowVal))
}

func (h *MinRowValHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallestElement(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return -1
	}

	//matrix = matrix[:minInt(k, len(matrix))]

	h := &MinRowValHeap{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, rowVal{row: i, col: 0, val: matrix[i][0]})
	}

	kthElem := -1
	for i := 0; i < k; i++ {
		e := heap.Pop(h).(rowVal)
		row, col := e.row, e.col
		kthElem = e.val

		if col < len(matrix[row])-1 {
			heap.Push(h, rowVal{
				row: row,
				col: col + 1,
				val: matrix[row][col+1],
			})
		}
	}

	return kthElem
}
