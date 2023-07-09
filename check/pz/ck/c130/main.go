package main

import (
	"fmt"
)

func main() {
	var (
		n         int
		a, b      string
		ans_list  []int
		ans_count int
	)
	fmt.Scan(&n)
	ans_count = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		if a == "n" || b == "n" {
			ans_list = append(ans_list, i+1)
			ans_count += 1
		}
	}
	fmt.Println(ans_count)
	for _, v := range ans_list {
		fmt.Println(v)
	}
}
