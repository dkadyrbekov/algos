package k_merge

import "container/heap"

type point struct {
	x int
	y int
}
type MaxDistancePointHeap []point

func (h MaxDistancePointHeap) Len() int { return len(h) }
func (h MaxDistancePointHeap) Less(i, j int) bool {
	return h[i].x*h[i].x+h[i].y*h[i].y > h[j].x*h[j].x+h[j].y*h[j].y
}                                            // max-heap
func (h MaxDistancePointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxDistancePointHeap) Push(x interface{}) {
	*h = append(*h, x.(point))
}

func (h *MaxDistancePointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {

	h := &MaxDistancePointHeap{}

	for i, p := range points {
		if i < k {
			heap.Push(h, point{x: p[0], y: p[1]})
			continue
		}

		if p[0]*p[0]+p[1]*p[1] < (*h)[0].x*(*h)[0].x+(*h)[0].y*(*h)[0].y {
			heap.Pop(h)
			heap.Push(h, point{x: p[0], y: p[1]})
		}
	}

	result := make([][]int, 0, h.Len())

	for h.Len() > 0 {
		p := heap.Pop(h).(point)

		result = append(result, []int{p.x, p.y})
	}

	return result
}
