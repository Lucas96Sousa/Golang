package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	repositories "api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userFindDB, err := repository.FindByEmail(user.Email)
	if err != nil {

		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePassword(userFindDB.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(userFindDB.ID)
	w.Write([]byte(token))
}
