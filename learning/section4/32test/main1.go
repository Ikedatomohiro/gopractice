package main

import (
	"fmt"
)

func main() {
	l := []int{100, 300, 23, 11, 23, 4, 6, 4}
	min := 0
	for _, v := range l {
		if min == 0 || min > v {
			min = v
		}
	}
	fmt.Println(min)
}