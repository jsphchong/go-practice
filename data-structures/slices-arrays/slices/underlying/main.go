package main

import "fmt"

func main(){
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...)

	fmt.Println(x)
	fmt.Println("len", len(x))
	fmt.Println("cap", cap(x))

	fmt.Println(y)
	fmt.Println("len", len(y))
	fmt.Println("cap", cap(y))
}