package main

import (
	"fmt"
)

func main() {
	var (
		n, s, e, h, l              int
		ans_s, ans_e, ans_h, ans_l int
	)
	ans_h = 0
	ans_l = 1000000
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &e, &h, &l)
		if i == 0 {
			ans_s = s
		} else if i == n-1 {
			ans_e = e
		}
		if h > ans_h {
			ans_h = h
		}
		if l < ans_l {
			ans_l = l
		}
	}
	fmt.Println(ans_s, ans_e, ans_h, ans_l)
}
