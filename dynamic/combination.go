package dynamic

import "sort"

func combinationSum(nums []int, target int) [][]int {
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}}

	sort.Ints(nums)

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(nums) && i >= nums[j]; j++ {
			for z := range dp[i-nums[j]] {
				cArr := append(make([]int, 0), dp[i-nums[j]][z]...)
				cArr = append(cArr, nums[j])
				sort.Ints(cArr)

				if !contains(dp[i], cArr) {
					dp[i] = append(dp[i], cArr)
				}
			}
		}
	}

	return dp[target]
}

func contains(results [][]int, newResult []int) bool {
	for _, res := range results {
		if len(res) != len(newResult) {
			continue
		}

		identical := true
		for i := 0; i < len(res); i++ {
			if res[i] != newResult[i] {
				identical = false
				break
			}
		}

		if identical {
			return true
		}
	}

	return false
}
