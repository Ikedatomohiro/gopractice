// https://paiza.jp/works/mondai/bfs_dfs_problems/bfs_dfs_problems__move_n_times
package main

import (
	"fmt"
)

func main() {
	var (
		H int
		W int
		N int
		y int
		x int
		dist int = 0
	)
	fmt.Scan(&H, &W, &N, &y, &x)
	area := make([][]string, H)
	for i := 0; i < H; i++ {
		area[i] = make([]string, W)
		for j := 0; j < W; j++ {
			area[i][j] = "."
		}
	}
	q := make([][]int, 0)
	q_tmp := make([][]int, 0)
	q = append(q, []int{y, x})
	for len(q) > 0 {
		y, x = q[0][0], q[0][1]
		// スタックを取り出す
		q = q[1:]
		area[y][x] = "*"
		if y - 1 >= 0 && area[y-1][x] != "*" {
			q_tmp = append(q_tmp, []int{y-1, x})
		}
		if y + 1 < H && area[y+1][x] != "*"{
			q_tmp = append(q_tmp, []int{y+1, x})
		}
		if x - 1 >= 0  && area[y][x-1] != "*"{
			q_tmp = append(q_tmp, []int{y, x-1})
		}
		if x + 1 < W && area[y][x+1] != "*" {
			q_tmp = append(q_tmp, []int{y, x+1})
		}
		if len(q) == 0 {
			q = q_tmp
			q_tmp = make([][]int, 0)
			dist++
		}
		if dist > N {
			break
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(area[i][j])
		}
		fmt.Println()
	}
}