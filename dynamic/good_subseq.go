package dynamic

func countGoodSubsequences(s string) int {
	dp := make([]int, len(s))

	return 0
}

//func countGoodSubsequences(s string) int {
//	dp := make([]int, len(s))
//
//	for i := 0; i < len(s); i++ {
//		counter := 0
//		for j := 0; j <= i; j++ {
//			counter += countGoodSubSeqWithLen(s[j:i], s[j], j+1)
//		}
//		dp[i] = counter
//	}
//
//	count := 0
//	for _, c := range dp {
//		count += c
//	}
//
//	return count
//}
//
//func countGoodSubSeqWithLen(str string, r byte, subseqLen int) int {
//	return 0
//}
//
//func isGoodSubsequence(s string) bool {
//	if len(s) == 0 {
//		return true
//	}
//
//	charCountMap := make(map[rune]int)
//	for _, r := range s {
//		charCountMap[r]++
//	}
//
//	count := charCountMap[rune(s[0])]
//
//	for _, r := range s {
//		if count != charCountMap[r] {
//			return false
//		}
//	}
//
//	return true
//}
