package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	// foo()
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	file, err := os.Open("./lesson.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // あとでファイルの閉じ忘れが内容にセットで記述できる
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
	fmt.Println("success")
}
