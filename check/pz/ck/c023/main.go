package main

import (
	"fmt"
)

func main() {
	var atari = make([]int, 6)
	for j := 0; j < 6; j++ {
		fmt.Scan(&atari[j])
	}
	var n, result int
	fmt.Scan(&n)
	result = 0
	for i := 0; i < n; i++ {
		var kuji = make([]int, 6)
		for j := 0; j < 6; j++ {
			fmt.Scan(&kuji[j])
		}
		for _, v := range kuji {
			for _, w := range atari {
				if v == w {
					result++
				}
			}
		}
		fmt.Println(result)
		result = 0
	}
}
