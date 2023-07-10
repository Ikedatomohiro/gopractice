package main

import (
	"fmt"
)

func main() {
	var (
		N              int
		t              string
		e, m, s, j, g  int
		total, t_total int
		ans            int
	)
	fmt.Scan(&N)
	ans = 0
	for i := 0; i < N; i++ {
		fmt.Scan(&t, &e, &m, &s, &j, &g)
		total = e + m + s + j + g
		if total >= 350 {
			if t == "s" {
				t_total = m + s
			} else if t == "l" {
				t_total = j + g
			}
			if t_total >= 160 {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}
