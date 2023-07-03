// https://paiza.jp/works/mondai/graph_dfs_problems/graph_dfs__path_one_step3
package main

import (
	"fmt"
)

func main() {
	var (
		n, s, k int
		graph   = make([]int, n)
		// s_index int
	)
	fmt.Scan(&n, &s, &k)
	for i := 1; i <= n; i++ {
		if i != s {
			graph = append(graph, i)
		}
	}
	fmt.Print(s)
	for i := 0; i < k; i++ {
		fmt.Print(" ", graph[i])
	}
	fmt.Println()
}
