package controllers

import "net/http"

// CreateUser -> Insert a new user
func CreateUser(w http.ResponseWriter, r *http.Response) {
  w.Write([]byte("Create user"))
}

// FindUsers -> Find all user in database
func FindUsers(w http.ResponseWriter, r *http.Response) {
  w.Write([]byte("Find all users"))
}

//FindUserbyId -> Find specific user by id
func FindUserbyId(w http.ResponseWriter, r *http.Response) {
  w.Write([]byte("Find user by id"))
}

//UpdateUser -> Update user data
func UpdateUser(w http.ResponseWriter, r *http.Response) {
 w.Write([]byte("Upadate user data"))
}

//DeleteUser -> Delete specific user
func DeleteUser(w http.ResponseWriter, r *http.Response) {
  w.Write([]byte("Delete user data"))
}
