// https://paiza.jp/works/mondai/graph_dfs_problems/graph_dfs__path_one_step2
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		n, s, k, now int
		list, path   []int
	)
	fmt.Scan(&n, &s, &k)
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	now = s
	path = append(path, now)
	for i := 0; i < k; i++ {
		for {
			rand.Seed(time.Now().UnixNano())
			pick := list[rand.Intn(len(list))]
			if pick != now {
				now = pick
				path = append(path, now)
				break
			}
		}

	}
	for _, val := range path {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
