package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"api/src/db"
	"api/src/models"
	repositories "api/src/repository"
	"api/src/responses"
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
	if err = user.Prepare("register"); err != nil {
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
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.Find(nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, users)
}

// FindUserbyId -> Find specific user by id
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, err := repository.FindByID(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser -> Update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

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
	if err = user.Prepare("modify"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.UpdateUser(userID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// DeleteUser -> Delete specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.DeleteUser(userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
