package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, pr := range prerequisites {
		if _, ok := graph[pr[0]]; !ok {
			graph[pr[0]] = make([]int, 0)
		}
		if _, ok := graph[pr[1]]; !ok {
			graph[pr[1]] = make([]int, 0)
		}

		if _, ok := inDegree[pr[0]]; !ok {
			inDegree[pr[0]] = 0
		}
		if _, ok := inDegree[pr[1]]; !ok {
			inDegree[pr[1]] = 0
		}

		graph[pr[1]] = append(graph[pr[1]], pr[0])
		inDegree[pr[0]]++
	}

	var orderCounter int
	stack := NewStack[int]()

	for k, v := range inDegree {
		if v == 0 {
			stack.Push(k)
		}
	}

	for !stack.IsEmpty() {
		v := stack.Pop()
		orderCounter++
		for _, c := range graph[v] {
			inDegree[c]--
			if inDegree[c] == 0 {
				stack.Push(c)
			}
		}
	}

	return orderCounter == len(graph)
}
