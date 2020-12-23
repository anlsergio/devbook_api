package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateUser adds a new user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers fetchs a set of all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.Get(nameOrUsername)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	
	responses.JSON(w, http.StatusOK, users)
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