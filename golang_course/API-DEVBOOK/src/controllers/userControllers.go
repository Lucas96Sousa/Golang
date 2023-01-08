package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser -> Insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}

	db, err := db.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

// FindUsers -> Find all user in database
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

// FindUserbyId -> Find specific user by id
func FindUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user by id"))
}

// UpdateUser -> Update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upadate user data"))
}

// DeleteUser -> Delete specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user data"))
}
