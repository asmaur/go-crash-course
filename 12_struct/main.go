package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) greet() string {
	return "Hi, my name is " + p.firstName + " " + p.lastName + " and i'am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirth() {
	p.age++
}

func main() {

	person1 := Person{"you", "me", 20}
	person2 := Person{"mani", "mooe", 20}

	fmt.Println(person1)
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

	person1.hasBirth()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
