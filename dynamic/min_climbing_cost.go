package dynamic

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}

	climbingCost := make([]int, len(cost)+1)
	climbingCost[0], climbingCost[1] = 0, 0

	for i := 2; i <= len(cost); i++ {
		climbingCost[i] = minInt2(climbingCost[i-2]+cost[i-2], climbingCost[i-1]+cost[i-1])
	}

	return climbingCost[len(cost)]
}

func minInt2(a, b int) int {
	if a < b {
		return a
	}

	return b
}
