package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		N int
		d int
		p float64
		day int = 0
		point int = 0
	)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&d, &p)
		str_d := strconv.Itoa(d)
		day = 0
		for _, char := range str_d {
			if string(char) == "3" {
				day = 3
			} else if string(char) == "5" {
				day = 5
			}
		}
		if day == 3 {
			point += int(p * 0.03)
		} else if day == 5 {
			point += int(p * 0.05)
		} else {
			point += int(p * 0.01)
		}
	}
	fmt.Println(point)
}