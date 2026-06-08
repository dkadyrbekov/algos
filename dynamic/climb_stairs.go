package dynamic

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := []int{1, 2}

	ways := 0
	for i := 3; i <= n; i++ {
		ways = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = ways
	}

	return ways
}
