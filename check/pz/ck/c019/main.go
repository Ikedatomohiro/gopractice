package main

import (
	"fmt"
)

func main() {
	var (
		Q, N, total int
	)
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		total = 1
		fmt.Scan(&N)
		for j := 2; j <= N / 2; j++ {
			if N % j == 0 {
				total += j
			}
		}
		if total == N {
			fmt.Println("perfect")
		} else if total == N - 1 {
			fmt.Println("nearly")
		} else {
			fmt.Println("neither")
		}

	}
}