package greedy

func gasStationJourney(gas []int, cost []int) int {
	if len(gas) != len(cost) || len(gas) == 0 {
		return -1
	}

	n := len(cost)

	check := 0
	for check < n {
		sum := 0
		lastChecked := check
		for i := 0; i < n; i++ {
			currI := (check + i) % n
			sum += gas[currI] - cost[currI]

			if sum < 0 {
				check = maxInt(check+1, lastChecked)
				break
			}

			lastChecked = currI

			if i == n-1 {
				return check
			}
		}
	}

	return -1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
