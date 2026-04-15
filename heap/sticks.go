package heap

import "container/heap"

func connectSticks(sticks []int) int {
	h := (*MinIntHeap)(&sticks)

	heap.Init(h)

	totalCost := 0
	for len(*h) > 1 {
		mergeCost := heap.Pop(h).(int) + heap.Pop(h).(int)
		totalCost += mergeCost

		heap.Push(h, mergeCost)
	}

	return totalCost
}
