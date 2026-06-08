package dynamic

func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	if len(weights) != len(values) {
		return 0
	}

	combs := findAllComb(capacity, weights, 0)

	maxProfit := 0

	for _, comb := range combs {
		combSum := 0
		for i := range comb {
			combSum += values[comb[i]]
		}

		if maxProfit < combSum {
			maxProfit = combSum
		}
	}

	return maxProfit
}

func findAllComb(capacity int, weights []int, step int) [][]int {
	combs := make([][]int, 0)

	for i := step; i < len(weights); i++ {
		if capacity-weights[i] < 0 {
			continue
		}

		subCombs := findAllComb(capacity-weights[i], weights, i+1)

		for j := range subCombs {
			subCombs[j] = append(subCombs[j], i)
		}
		combs = append(combs, subCombs...)
	}

	return combs
}

//var maxSubArrProfit map[int]int
//
//func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
//	if len(weights) != len(values) {
//		return 0
//	}
//
//	maxSubArrProfit = make(map[int]int)
//
//	return findMaxKnapsackProfit2(capacity, weights, values, 0)
//}
//
//func findMaxKnapsackProfit2(capacity int, weights []int, values []int, step int) int {
//	maxProfit := 0
//
//	for i := step; i < len(weights); i++ {
//		if capacity-weights[i] < 0 {
//			continue
//		}
//
//		var tmp int
//		if profit, ok := maxSubArrProfit[i]; ok {
//			tmp = profit
//		} else {
//			tmp = findMaxKnapsackProfit2(capacity-weights[i], weights, values, i+1) + values[i]
//			maxSubArrProfit[i] = tmp
//		}
//
//		if tmp > maxProfit {
//			maxProfit = tmp
//		}
//	}
//
//	return maxProfit
//}

//newWeights := make([]int, 0, len(weights)-1)
//newValues := make([]int, 0, len(weights)-1)
//if i < len(weights)-1 {
//	newWeights = append(newWeights, weights[i+1:]...)
//	newValues = append(newValues, values[i+1:]...)
//}
