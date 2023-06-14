package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User no found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	e1 := &UserNotFound{Username: "mike"}
	e2 := &UserNotFound{Username: "mike"}
	if err := myFunc(); err != nil {
		fmt.Println(err)
		if err == e1 {
			fmt.Println("e1")
		} else if err == e2 {
			fmt.Println("e2")
		}
	}
}