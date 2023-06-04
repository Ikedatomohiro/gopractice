package main

import (
	"fmt"
)

func by2(num int) string {
	if num % 2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(20)

	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	// fmt.Println(result2) // if文の中でしか使えない
	// num := 9
	// if num % 2 == 0 {
	// 	fmt.Println("Even")
	// } else if num % 3 == 0 {
	// 	fmt.Println("by 3")
	// } else {
	// 	fmt.Println("Odd")
	// }

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}