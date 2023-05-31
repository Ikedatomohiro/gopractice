package main

import (
	"fmt"
)

func main() {
	var (
		S string
		dict = map[string]string{
			"A" : "4",
			"E" : "3",
			"G" : "6",
			"I" : "1",
			"O" : "0",
			"S" : "5",
			"Z" : "2",
		}
		ans string
	)
	fmt.Scan(&S)

	for _, r := range S {
		if dict[string(r)] != "" {
			ans += dict[string(r)]
		} else {
			ans += string(r)
		}
	}
	fmt.Println(ans)
}