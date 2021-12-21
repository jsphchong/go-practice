package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func prepareRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// Routes 
	router.Path("/posts").HandlerFunc(RESTHandler).Methods(http.MethodPost)
	router.Path("/posts/{id}").HandlerFunc(RESTHandler).Methods(http.MethodGet)
	router.Path("/posts/{id}").HandlerFunc(RESTHandler).Methods(http.MethodPut)
	router.Path("/posts/{id}").HandlerFunc(RESTHandler).Methods(http.MethodDelete)
	router.Path("/posts").HandlerFunc(RESTHandler).Methods(http.MethodGet)

	return router
}