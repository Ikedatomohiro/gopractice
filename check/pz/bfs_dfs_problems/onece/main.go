// https://paiza.jp/works/mondai/bfs_dfs_problems/bfs_dfs_problems__move_once
package main

import (
	"fmt"
)

func main() {
	var (
		H int
		W int
		y int
		x int
	)
	fmt.Scan(&H, &W, &y, &x)
	area := make([][]string, H)

	for i := range area {
		area[i] = make([]string, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			area[i][j] = "."
		}
	}
	area[y][x] = "*"
	if y - 1 >= 0 {
		area[y-1][x] = "*"
	}
	if y + 1 < H {
		area[y+1][x] = "*"
	}
	if x - 1 >= 0 {
		area[y][x-1] = "*"
	}
	if x + 1 < W {
		area[y][x+1] = "*"
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(area[i][j])
		}
		fmt.Println()
	}
}