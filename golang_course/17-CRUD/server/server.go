package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type user struct {
  ID uint32 `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
}

// CreateUser insert user into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      w.Write([]byte("Failure to read body"))
      return
    }
    
    var user user
  
    if err = json.Unmarshal(body, &user); err != nil {
      w.Write([]byte("Failure convert user to  struct"))
    }

    db, err := db.Connect()
    if err != nil {
      w.Write([]byte("Failure to connect database"))
      return
    }
    statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
    if err != nil {
      w.Write([]byte("Failure to create statement"))
      return
    }
    defer statement.Close()

    insert, err := statement.Exec(user.Name, user.Email)

    if err != nil {
      w.Write([]byte("Failure to create exec"))
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
  db, err := db.Conect()
    if err != nil {
      w.Write([]byte("Failure to connect database"))
      return
    }
    defer db.Close()
  
  row, erro := db.Query("select * from users")
    if err != nil {
      w.Write([]byte("Error to find users"))
      return
  }
  defer row.Close()

  var users []user

  for row.Next() {
    var user user

    if err != row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
      w.Write([]byte("Erro to scan user"))
      return

    }
    users = append(users, user)
  }

  w.WriteHeader(http.StatusOK)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {}
