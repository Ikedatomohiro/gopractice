package main

import (
	"fmt"
)

func main() {
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	fmt.Printf("%T\n", p)
	// var p2 *int
	// fmt.Println(p2)
	// *p2++
	// fmt.Println(p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)
}