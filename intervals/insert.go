package intervals

func insertInterval(existingIntervals [][]int, newInterval []int) [][]int {
	const (
		start = 0
		end   = 1
	)

	result := make([][]int, 0, len(existingIntervals)+1)

	for i := 0; i < len(existingIntervals); i++ {
		if existingIntervals[i][start] > newInterval[end] {
			result = append(result, newInterval)
			result = append(result, existingIntervals[i:]...)

			return result
		}

		if intervalsIntersects(existingIntervals[i][start], existingIntervals[i][end], newInterval[start], newInterval[end]) {
			newInterval[start] = minInt(existingIntervals[i][start], newInterval[start])
			newInterval[end] = maxInt(existingIntervals[i][end], newInterval[end])
			continue
		}

		result = append(result, existingIntervals[i])
	}

	result = append(result, newInterval)

	return result
}

func intervalsIntersects(start1, end1, start2, end2 int) bool {
	return minInt(end1, end2)-maxInt(start1, start2) >= 0
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
