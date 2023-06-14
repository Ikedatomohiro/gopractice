package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}


	x := <-ch // queueのような挙動になる
	fmt.Println(x)
	fmt.Println(len(ch))
	y := <-ch
	fmt.Println(y)
	fmt.Println(len(ch))
	ch <- 300
	fmt.Println(len(ch))


}