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



// createUser creates a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	fmt.Println(users)
	json.NewEncoder(w).Encode(user)
}

// getUsers gets all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// getUser gets a single user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode(&User{})
}



// updateUser updates a user by ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users[index].Name = params["name"]
			users[index].Email = params["email"]
			users[index].Password = params["password"]
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	fmt.Println("->>>>>",users)

	json.NewEncoder(w).Encode(users)
}

// deleteUser deletes a user by ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(users)
}



