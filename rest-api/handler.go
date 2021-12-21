package main

import "net/http"

func RESTHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == "/posts" {
			
		} else {

		}
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(http.StatusOK)
}