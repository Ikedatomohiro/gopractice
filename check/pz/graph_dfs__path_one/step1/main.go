// https://paiza.jp/works/mondai/graph_dfs_problems/graph_dfs__path_one_step1
package main

import (
	"fmt"
)

func main() {
	var (
		n, s, a, v, max int
	)
	max = 0
	fmt.Scan(&n, &s)
	g := make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		for j := 0; j < a; j++ {
			fmt.Scan(&v)
			g[i] = append(g[i], v)
		}
	}
	for _, val := range g[s-1] {
		if max < val {
			max = val
		}
	}
	fmt.Println(max)
}
