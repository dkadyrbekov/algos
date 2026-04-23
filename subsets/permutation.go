package subsets

type runeSet struct {
	hashMap map[rune]bool
}

func newRuneSet() *runeSet {
	s := new(runeSet)
	s.hashMap = make(map[rune]bool)
	return s
}

func (s *runeSet) Add(value rune) {
	s.hashMap[value] = true
}

func (s *runeSet) Delete(value rune) {
	delete(s.hashMap, value)
}

func (s *runeSet) Exists(value rune) bool {
	_, ok := s.hashMap[value]
	return ok
}

func (s *runeSet) Keys() []rune {
	keys := make([]rune, 0, len(s.hashMap))
	for r := range s.hashMap {
		keys = append(keys, r)
	}
	return keys
}

func permuteWord(word string) []string {
	rs := newRuneSet()
	for _, r := range word {
		rs.Add(r)
	}

	permutations := make([]string, 0)

	for _, r := range rs.Keys() {
		rs.Delete(r)
		permutations = append(permutations, getPermutations(string(r), rs)...)
		rs.Add(r)
	}

	return permutations
}

func getPermutations(prefix string, set *runeSet) []string {
	if len(set.hashMap) == 0 {
		return []string{prefix}
	}

	permutations := make([]string, 0)
	for _, r := range set.Keys() {
		set.Delete(r)
		permutations = append(permutations, getPermutations(prefix+string(r), set)...)
		set.Add(r)
	}

	return permutations
}
