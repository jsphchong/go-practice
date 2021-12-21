package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := Person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
	}

	p2 := Person{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       27,
	}

	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
