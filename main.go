package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


var users []User

func main() {
	router := mux.NewRouter()
	// Create a new user
	router.HandleFunc("/users", createUser).Methods("POST")
	// Get all users
	router.HandleFunc("/users", getUsers).Methods("GET")
	// Get a single user by ID
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	// Update a user by ID
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	// Delete a user by ID
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

