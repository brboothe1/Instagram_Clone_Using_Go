package main

// modifyDatabase.go houses the logic used for getting, creating, deleting, and changing information in the database

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	db.First(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.Create(&user)
	err = createdUser.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	db.First(&user, params["id"])
	db.Delete(&user)

	json.NewEncoder(w).Encode(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	db.First(&user, params["id"])
	db.Update(&user)

	json.NewEncoder(w).Encode(&user)
}