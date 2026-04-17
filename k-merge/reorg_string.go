package k_merge

import (
	"container/heap"
	"strings"
)

type charFreq struct {
	char rune
	freq int
}
type MaxCharFreqHeap []charFreq

func (h MaxCharFreqHeap) Len() int           { return len(h) }
func (h MaxCharFreqHeap) Less(i, j int) bool { return h[i].freq > h[j].freq } // min-heap
func (h MaxCharFreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCharFreqHeap) Push(x interface{}) {
	*h = append(*h, x.(charFreq))
}

func (h *MaxCharFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reorganizeString(str string) string {
	m := make(map[rune]int)

	for _, ch := range str {
		m[ch]++
	}

	h := &MaxCharFreqHeap{}

	for char, freq := range m {
		heap.Push(h, charFreq{char: char, freq: freq})
	}

	if h.Len() == 0 {
		return ""
	}

	var result strings.Builder
	prevCharFreq := heap.Pop(h).(charFreq)
	result.WriteRune(prevCharFreq.char)
	prevCharFreq.freq--
	for h.Len() > 0 {
		currCharFreq := heap.Pop(h).(charFreq)
		result.WriteRune(currCharFreq.char)
		currCharFreq.freq--

		if prevCharFreq.freq > 0 {
			heap.Push(h, prevCharFreq)
		}

		prevCharFreq = currCharFreq
	}

	if prevCharFreq.freq != 0 {
		return ""
	}

	return result.String()
}
