package main

import (
	"fmt"
)

func main() {
	var (
		s string
		t string
	)
	fmt.Scan(&s, &t)
	str := fmt.Sprintf("%s@%s", s, t)
	fmt.Println(str)
}