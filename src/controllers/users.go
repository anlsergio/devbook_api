package controllers

import "net/http"

// CreateUser adds a new user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a user"))
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