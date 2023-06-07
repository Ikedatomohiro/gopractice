// https://paiza.jp/works/mondai/bfs_dfs_problems/bfs_dfs_problems__judge_reach
package main

import (
	"fmt"
)

func main() {
	var (
		H, W, sy, sx, gy, gx int
		S string
	)
	fmt.Scan(&H, &W, &sy, &sx, &gy, &gx)
	area := make([][]string, H)
	for i := range area {
		area[i] = make([]string, W)
	}
	for i := 0; i < H; i++ {
		fmt.Scan(&S)
		for j := 0; j < W; j++ {
			area[i][j] = string(S[j])
		}
	}
	q := [][]int{{sy, sx}}
	move := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		y, x := q[0][0], q[0][1]
		q = q[1:]
		area[y][x] = "*"
		for _, m := range move {
			if y + m[0] >= 0 && y + m[0] < H && x + m[1] >= 0 && x + m[1] < W && area[y+m[0]][x+m[1]] != "*" && area[y+m[0]][x+m[1]] != "#" {
				if y+m[0] == gy && x+m[1] == gx {
					fmt.Println("Yes")
					return
				}
				q = append(q, []int{y+m[0], x+m[1]})
			}
		}
	}
	fmt.Println("No")
}