package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main2() {
	// os := getOsName()
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Windows!")
	default:
		fmt.Println("Default!", os)
	}
	// fmt.Println("Default!", os) これはエラーとなる

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	}
}