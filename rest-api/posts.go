package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var counter int

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var p Post 
	json.Unmarshal(body, &p)

	if p.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	counter++

	id := strconv.Itoa(counter)

}