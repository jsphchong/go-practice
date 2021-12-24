package main

import "net/http"

func RESTHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == "/posts" {
			RetrieveAllPosts(w, req)
		} else {
			RetrieveSinglePost(w, req)
		}
	case http.MethodPost:
		CreatePost(w, req)
	case http.MethodPut:
		UpdatePost(w, req)
	case http.MethodDelete:
		DeletePost(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(http.StatusOK)
}