package subsets

func findAllSubsets(nums []int) []Set {
	uniqueNums := make(map[int]struct{})

	for _, v := range nums {
		uniqueNums[v] = struct{}{}
	}

	sets := []Set{*NewSet()}

	for num := range uniqueNums {
		for _, s := range sets {
			if !s.Exists(num) {
				ns := NewSet()
				for k := range s.hashMap {
					ns.Add(k)
				}

				ns.Add(num)
				sets = append(sets, *ns)
			}
		}
	}

	return sets
}
