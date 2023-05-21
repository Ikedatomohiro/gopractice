package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n int
		m int
		r int
		print bool
		exist bool = false
	)
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&r)
		str_r := strconv.Itoa(r)
		print = true
		for _, char := range str_r {
			num, _ := strconv.Atoi(string(char))
			if num == n {
				print = false
				break
			}
		}
		if print {
			fmt.Println(r)
			exist = true
		}
	}
	if !exist {
		fmt.Println("none")
	}
}
