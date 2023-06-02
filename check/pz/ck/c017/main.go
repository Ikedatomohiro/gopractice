package main

import (
	"fmt"
)

func main() {
	var (
		a, b, n, A, B int
	)
	fmt.Scan(&a, &b, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A, &B)
		if a > A {
			fmt.Println("High")
		} else if a < A {
			fmt.Println("Low")
 		} else {
			if b > B {
				fmt.Println("Low")
			} else {
				fmt.Println("High")
			}
		}
	}
}