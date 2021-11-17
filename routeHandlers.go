package main

//routeHandlers.go handles the API functions used to modify information in the PostgreSQL database

import(
	"github.com/gorilla/mux"
)

func AddRoutes (route *mux.Router) {
	route.HandleFunc("/users", getUsers).Methods("GET")
	route.HandleFunc("/users/{id}", getUsers).Methods("GET")
	route.HandleFunc("/create/user",createUser).Methods("POST")
	route.HandleFunc("/delete/user/{id}", deleteUser).Methods("DELETE")
	route.HandleFunc("/update/user/{id}", updateUser).Methods("POST")
}