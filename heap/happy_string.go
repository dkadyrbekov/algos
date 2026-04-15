package heap

import (
	"container/heap"
	"strings"
)

type charCounter struct {
	c     rune
	count int
}
type MaxCharHeap []charCounter

func (h MaxCharHeap) Len() int           { return len(h) }
func (h MaxCharHeap) Less(i, j int) bool { return h[i].count > h[j].count } // max-heap
func (h MaxCharHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCharHeap) Push(x interface{}) {
	*h = append(*h, x.(charCounter))
}

func (h *MaxCharHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	h := &MaxCharHeap{
		{c: 'a', count: a},
		{c: 'b', count: b},
		{c: 'c', count: c},
	}
	heap.Init(h)

	var s strings.Builder
	maxC := heap.Pop(h).(charCounter)

	for maxC.count > 0 {
		s.WriteRune(maxC.c)
		maxC.count--
		if maxC.count > 0 && maxC.count >= (*h)[0].count {
			s.WriteRune(maxC.c)
			maxC.count--
		}

		nextMaxC := heap.Pop(h).(charCounter)
		heap.Push(h, maxC)
		maxC = nextMaxC
	}

	return s.String()
}
