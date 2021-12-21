package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type ByLast []Person

func (bl ByLast) Len() int {return len(bl)}
func (bl ByLast) Swap(i, j int) {bl[i], bl[j] = bl[j], bl[i]}
func (bl ByLast) Less(i, j int) bool {return bl[i].Last < bl[j].Last}

func main() {
	p1 := Person{"Joseph", "Chong", 32}
	p2 := Person{"Joseph", "Natvig", 32}
	p3 := Person{"Dennis", "Eambovornchai", 32}
	p4 := Person{"Kevin", "Tran", 32}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByLast(people))
	fmt.Println(people)
}
