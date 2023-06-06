package main

import (
	"fmt"
)

type Human interface {
	Say() string
	ChangeNameMessage() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println("Hello, I am " + p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func (p Person) ChangeNameMessage() string {
	message := "Name is changed"
	return message
}

func main() {
	var mike Human = &Person{"Mike"}
	// var x Human = &Person{"X"}
	DriveCar(mike)
	mike.Say()
	// DriveCar(x)
	message := mike.ChangeNameMessage()
	// mike.Name = mike.ChangeNameMessage()
	// fmt.Println(mike.Name) // これではエラーとなる。
	if person, ok := mike.(*Person); ok {
		fmt.Println(person.Name)
	}
	fmt.Println(message)
}