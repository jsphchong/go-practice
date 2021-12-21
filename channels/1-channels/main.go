package main

import (
	"fmt"
)

func main(){
	ca := make(chan int)
	go func(){
		ca <- 42
	}()
	fmt.Print(<-ca)
}