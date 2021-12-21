package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 7}
	for index, element := range x {
		fmt.Println(index, "+", element)
	}
}
