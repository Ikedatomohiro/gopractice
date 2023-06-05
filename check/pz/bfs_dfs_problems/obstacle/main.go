// https://paiza.jp/works/mondai/bfs_dfs_problems/bfs_dfs_problems__move_obstacle
package main

import (
	"fmt"
)

func main() {
	var (
		H, W, N int
		y, x int
		d int = 0
		S string
		q [][]int
		tmp_q [][]int
	)
	fmt.Scan(&H, &W, &N, &y, &x)
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
	q = append(q, []int{y, x})
	move := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		y, x = q[0][0], q[0][1]
		q = q[1:]
		area[y][x] = "*"
		for _, m := range move {
			if y + m[0] >= 0 && y + m[0] < H && x + m[1] >= 0 && x + m[1] < W && area[y+m[0]][x+m[1]] != "*" && area[y+m[0]][x+m[1]] != "#" {
				tmp_q = append(tmp_q, []int{y+m[0], x+m[1]})
			}
		}
		if len(q) == 0 {
			q = tmp_q
			tmp_q = make([][]int, 0)
			d++
		}
		if d > N {
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