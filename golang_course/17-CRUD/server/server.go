package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"crud/db"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser insert user into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error to read body"))
		return
	}

	var user user

	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error convert user to  struct"))
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error to connect database"))
		return
	}
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error to create statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Error to create exec"))
		return
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Error to get insert id"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created!!! ID: %d", idInsert)))
}

// FindUser search all users
func FindUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error to connect database"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Error to find users"))
		return
	}
	defer row.Close()

	var users []user

	for row.Next() {
		var user user

		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro to scan user"))
			return

		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to convert users to json"))
		return
	}
}

// FindUsers by ID
func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error to convert params to integer"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error to connect database"))
		return
	}

	row, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Error to find user"))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error to scan user"))
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro to convert user to json"))
		return
	}
}

// UpdateUser method to update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to convert param for interger"))
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error to read body request"))
		return
	}
	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error to convert user for struct"))
		return
	}
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error to connect database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")

	if err != nil {
		w.Write([]byte("Error to create statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Error to update user data "))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
