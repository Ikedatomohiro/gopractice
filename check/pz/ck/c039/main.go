package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s   string
		ans int
	)
	fmt.Scan(&s)
	for _, r := range strings.Split(s, "") {
		if r == "<" {
			ans += 10
		} else if r == "/" {
			ans += 1
		}
	}
	fmt.Println(ans)
}
