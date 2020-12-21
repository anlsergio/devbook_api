package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser adds a new user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("The user has been added. User ID: %d", userID)))
}

// GetUsers fetchs a set of all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching all users"))
}

// GetUser fetch information about an specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching a user"))
}

// UpdateUser updates information about an specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser completely removes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}