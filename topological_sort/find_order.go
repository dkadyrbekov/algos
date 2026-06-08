package main

import (
	"fmt"
)

type DAG struct {
	edges map[rune][]rune
}

func NewDAG(dependencies [][]rune) DAG {
	dag := DAG{
		edges: make(map[rune][]rune),
	}

	for _, dp := range dependencies {
		dag.AddEdge(dp[0], dp[1])
	}

	return dag
}

func (d *DAG) AddEdge(from, to rune) {
	if _, ok := d.edges[from]; !ok {
		d.edges[from] = make([]rune, 0)
	}
	if _, ok := d.edges[to]; !ok {
		d.edges[to] = make([]rune, 0)
	}

	d.edges[from] = append(d.edges[from], to)
}

type Stack struct {
	q []rune
}

func NewStack() Stack {
	return Stack{}
}

func (q *Stack) Push(r rune) {
	q.q = append(q.q, r)
}

func (q *Stack) Pop() rune {
	res := q.q[len(q.q)-1]

	q.q = q.q[:len(q.q)-1]

	return res
}

func (q *Stack) IsEmpty() bool {
	return len(q.q) == 0
}

func reverseArr(arr []rune) []rune {
	reversedArr := make([]rune, 0, len(arr))

	for j := len(arr) - 1; j >= 0; j-- {
		reversedArr = append(reversedArr, arr[j])
	}

	return reversedArr
}

func findOrder(dependencies [][]rune) []rune {
	dag := NewDAG(dependencies)

	inDegree := make(map[rune]int)

	for _, dep := range dag.edges {
		for _, v := range dep {
			inDegree[v]++
		}
	}

	stack := NewStack()

	for node := range dag.edges {
		if inDegree[node] == 0 {
			stack.Push(node)
		}
	}

	order := make([]rune, 0)

	for !stack.IsEmpty() {
		curr := stack.Pop()
		order = append(order, curr)
		for _, v := range dag.edges[curr] {
			inDegree[v]--
			if inDegree[v] == 0 {
				stack.Push(v)
			}
		}
	}

	if len(order) != len(dag.edges) {
		return make([]rune, 0)
	}

	return reverseArr(order)
}

func main() {
	testCases := [][][]rune{
		{},
		{
			{'C', 'A'},
		},
		{
			{'C', 'A'},
			{'B', 'A'},
			{'D', 'C'},
			{'E', 'B'},
			{'E', 'D'},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("tc := %v\nresult = %v\n", tc, findOrder(tc))
	}
}
