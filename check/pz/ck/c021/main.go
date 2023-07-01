package main

import (
	"fmt"
)

func main() {
	var (
		xc, yc, r_1, r_2, n int
		x, y                int
	)
	fmt.Scan(&xc, &yc, &r_1, &r_2, &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		if (x-xc)*(x-xc)+(y-yc)*(y-yc) >= r_1*r_1 && (x-xc)*(x-xc)+(y-yc)*(y-yc) <= r_2*r_2 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
