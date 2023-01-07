package controllers

import (
	"api/src/db"
	"api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser -> Insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    bodyRequest, err := ioutil.ReadAll(r.Body)
  if err != nil {
    log.Fatal(err)
  }

  var user models.Users

  if err = json.Unmarshal(bodyRequest, &user); err != nil {
    log.Fatal(err)
  }

  db, err := db.Connect()

 if err != nil {
    log.Fatal(err)
 }  
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
