package main

import (
	"log"
	"net/http"
)

var pm PostMap

func main() {
	myRouter := prepareRouter()
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}