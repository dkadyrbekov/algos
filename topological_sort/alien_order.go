package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func alienOrder(words []string) string {

	if len(words) == 0 {
		return ""
	}

	graph := make(map[rune][]rune)

	for _, r := range words[0] {
		if _, ok := graph[r]; !ok {
			graph[r] = make([]rune, 0)
		}
	}

	for i := 1; i < len(words); i++ {
		if addToGraph([]rune(words[i-1]), []rune(words[i]), graph) != nil {
			return ""
		}
	}

	inDegree := make(map[rune]int)

	for k := range graph {
		inDegree[k] = 0
	}

	for _, edges := range graph {
		for _, val := range edges {
			inDegree[val]++
		}
	}

	stack := NewStack[rune]()

	for k, v := range inDegree {
		if v == 0 {
			stack.Push(k)
		}
	}

	var result strings.Builder
	for !stack.IsEmpty() {
		v := stack.Pop()
		result.WriteRune(v)
		for _, val := range graph[v] {
			inDegree[val]--
			if inDegree[val] == 0 {
				stack.Push(val)
			}
		}
	}

	//printGraph(graph, words)

	res := result.String()
	if utf8.RuneCountInString(res) != len(graph) {
		return ""
	}

	return res
}

func addToGraph(w1, w2 []rune, graph map[rune][]rune) error {
	for _, r := range w2 {
		if _, ok := graph[r]; !ok {
			graph[r] = make([]rune, 0)
		}
	}

	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] == w2[i] {
			if i == len(w2)-1 && len(w1) > len(w2) {
				return fmt.Errorf("error")
			}
			continue
		}

		graph[w1[i]] = append(graph[w1[i]], w2[i])
		break
	}

	return nil
}

func printGraph(graph map[rune][]rune, words []string) {
	fmt.Printf("words = %v\n", words)

	fmt.Printf("graph===\n")
	for k, v := range graph {
		fmt.Printf("[%c] -> %c\n", k, v)
	}

	fmt.Printf("===============\n\n")
}

func main() {
	examples := [][]string{
		{"m", "mx", "mxe", "mxer", "mxerl", "mxerlo", "mxerlos", "mxerlost", "mxerlostr", "mxerlostrpq", "mxerlostrp"},
		{"xro", "xma", "per", "prt", "oxh", "olv"},
		{"acb", "acd", "bd", "zce", "zcdr"},
		{"acb", "acd", "bd", "zce", "acdr"},
		{"ac", "ab", "zc", "zb"},
	}

	for _, e := range examples {
		alienOrder(e)
	}
}
