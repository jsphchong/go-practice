package main

import (
	"log"
	"net/http"
)

func main() {
	myRouter := prepareRouter()
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}