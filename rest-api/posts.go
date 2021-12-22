package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
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

	if _, found := pm[id]; found {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		p.Id = id
		pm[id] = &p
	}

	json.NewEncoder(w).Encode(pm[id])
}

func RetrieveSinglePost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	if v, ok := pm[id]; ok {
		json.NewEncoder(w).Encode(v)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}	
}

func UpdatePost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := pm[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var p Post
	json.Unmarshal(body, &p)

	if p.Title == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		p.Id = id
		pm[id] = &p
		json.NewEncoder(w).Encode(p) 
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := pm[id]; ok {
		delete(pm, id)
	} else { 
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func RetrieveAllPosts(w http.ResponseWriter, r *http.Request){
	pt := []Post{}
	for _, v := range pm {
		pt = append(pt, *v)
	}
	sort.Sort(ById(pt))
	json.NewEncoder(w).Encode(pt)
}