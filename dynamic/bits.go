package dynamic

func countingBits(n int) []int {
	result := make([]int, 0, n+1)

	for i := 0; i <= n; i++ {
		result = append(result, countBits(i))
	}
	return result
}

func countBits(x int) int {
	count := 0

	for x > 0 {
		x &= (x - 1)
		count++
	}

	return count
}
