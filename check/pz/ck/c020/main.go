package main

import (
	"fmt"
)

func main() {
	var (
		m, p, q, r float64
	)
	fmt.Scan(&m, &p, &q)
	r = m * ((100 - p) / 100) * ((100 - q) / 100)
	fmt.Println(r)
}