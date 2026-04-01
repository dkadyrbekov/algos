package sliding_window

func findRepeatedDnaSequences(s string) []string {

	substrMap := make(map[string]int)

	i, j := 0, 10

	for j <= len(s) {
		substrMap[s[i:j]]++
		j++
		i++
	}

	repeatedDNASeq := []string{}

	for k, v := range substrMap {
		if v > 1 {
			repeatedDNASeq = append(repeatedDNASeq, k)
		}
	}

	return repeatedDNASeq
}
