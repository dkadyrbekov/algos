package heap

import (
	"container/heap"
	"sort"
)

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumMachines(tasks [][]int) int {
	n := len(tasks)
	if n == 0 {
		return 0
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})

	h := &MinIntHeap{tasks[0][1]}
	heap.Init(h)

	maxCPUs, counter := 1, 1
	for i := 1; i < n; i++ {
		if tasks[i][0] < (*h)[0] {
			heap.Push(h, tasks[i][1])
			counter++
			if counter > maxCPUs {
				maxCPUs = counter
			}
		} else {
			heap.Pop(h)
			heap.Push(h, tasks[i][1])
		}
	}

	return maxCPUs
}
