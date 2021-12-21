package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type SecretAgent struct {
	Person
	ltk bool
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		ltk: true,
	}

	p2 := Person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.firstName, sa1.lastName, sa1.age, sa1.ltk)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
}
