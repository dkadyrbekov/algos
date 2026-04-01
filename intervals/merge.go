package intervals

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	resultIntervals := make([][]int, 0, len(intervals))

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		for j := i + 1; j < len(intervals); j++ {
			if end >= intervals[j][0] {
				end = maxEnd(end, intervals[j][1])
				i = j
				continue
			}
			break
		}

		resultIntervals = append(resultIntervals, []int{start, end})
	}

	return resultIntervals
}

func maxEnd(a, b int) int {
	if a > b {
		return a
	}

	return b
}
