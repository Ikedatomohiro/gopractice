package main

import (
	"fmt"
)

func main() {
	var (
		n int
		r int
		d int
		h int
		w int
		dia int
	)
	fmt.Scan(&n, &r)
	for i := 0; i < n; i++ {
		fmt.Scan(&h, &w, &d)
		dia = r * 2
		if h >= dia && w >= dia && d >= dia {
			fmt.Println(i + 1)
		}
	}
}
