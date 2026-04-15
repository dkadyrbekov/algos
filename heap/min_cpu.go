package heap

import (
	"container/heap"
	"sort"
)

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
