package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	close(c) // CLOSE THE CHANNEL

	fmt.Println("About to exit")
}

// send channel
func send(c chan<- int){
	c <- 42
}

// receive channel
func receive(c <-chan int){
	fmt.Println("the value receive from the channel:", <-c)
}