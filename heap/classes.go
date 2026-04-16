package heap

import "container/heap"

type PassRatioHeap [][]int

func (h PassRatioHeap) Len() int { return len(h) }
func (h PassRatioHeap) Less(i, j int) bool {
	passRi := float64((h[i][0]+1))/float64((h[i][1]+1)) - float64(h[i][0])/float64(h[i][1])
	passRj := float64((h[j][0]+1))/float64((h[j][1]+1)) - float64(h[j][0])/float64(h[j][1])
	return passRi > passRj
}
func (h PassRatioHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PassRatioHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PassRatioHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {

	h := (*PassRatioHeap)(&classes)
	heap.Init(h)

	for i := 0; i < extraStudents; i++ {
		maxR := heap.Pop(h).([]int)

		maxR[0] += 1
		maxR[1] += 1

		heap.Push(h, maxR)
	}

	avgRation := 0.0
	for i := 0; i < len(classes); i++ {
		avgRation += float64(classes[i][0]) / float64(classes[i][1])
	}

	return avgRation / float64(len(classes))
}
