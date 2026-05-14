package dynamic

func countPalindromicSubstrings(s string) int {
	counter := 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				counter++
			}
		}
	}

	return counter
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
