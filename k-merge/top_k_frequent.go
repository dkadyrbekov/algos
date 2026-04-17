package k_merge

import "container/heap"

type numFreq struct {
	num  int
	freq int
}
type MinNumFreqHeap []numFreq

func (h MinNumFreqHeap) Len() int           { return len(h) }
func (h MinNumFreqHeap) Less(i, j int) bool { return h[i].freq < h[j].freq } // min-heap
func (h MinNumFreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinNumFreqHeap) Push(x interface{}) {
	*h = append(*h, x.(numFreq))
}

func (h *MinNumFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	counter := 0
	h := &MinNumFreqHeap{}

	for num, freq := range m {
		if counter < k {
			heap.Push(h, numFreq{num: num, freq: freq})
			counter++
			continue
		}

		if freq <= (*h)[0].freq {
			continue
		}
		heap.Pop(h)
		heap.Push(h, numFreq{num: num, freq: freq})
	}

	result := make([]int, 0, h.Len())

	for h.Len() > 0 {
		top := heap.Pop(h).(numFreq)

		result = append(result, top.num)
	}

	return result
}
