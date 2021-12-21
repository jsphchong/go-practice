package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	p2 := Person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
}
