package main

import (
	"fmt"
)

func main() {
	var (
		n          int
		s          string
		i, a       int
		ans1, ans2 int
	)
	fmt.Scan(&n)
	for j := 0; j < n; j++ {
		fmt.Scan(&s)
		if s == "SET" {
			fmt.Scan(&i, &a)
			if i == 1 {
				ans1 = a
			} else {
				ans2 = a
			}
		} else if s == "ADD" {
			fmt.Scan(&a)
			ans2 = ans1 + a
		} else if s == "SUB" {
			fmt.Scan(&a)
			ans2 = ans1 - a
		}
	}
	fmt.Println(ans1, ans2)
}
