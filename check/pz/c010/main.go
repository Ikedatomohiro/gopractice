package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64
		b float64
		R float64
		N int
		x float64
		y float64
	)
	fmt.Scan(&a, &b, &R)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x, &y)
		if math.Pow((x-a), 2) + math.Pow((y-b), 2) >= math.Pow(R, 2) {
			fmt.Println("silent")
		} else {
			fmt.Println("noisy")
		}
	}
}