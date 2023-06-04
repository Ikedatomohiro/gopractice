package main

import (
    "fmt"
)

func main() {
    var i int = 100
    var j int = 200
    var p1 *int
    var p2 *int
    p1 = &i
	fmt.Println(p1, *p1)
    p2 = &j
    i = *p1 + *p2
	fmt.Println(i, *p1, *p2)
    p2 = p1
	fmt.Println(p2, *p2)
    j = *p2 + i
    fmt.Println(j)

	var p3 *int = &i
	*p3++
	fmt.Println(*p3)


}