package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represent user repository
type Users struct {
  db *sql.DB
}

// Create new User repository
func NewUserRepository(db *sql.DB) *Users {
  return &Users{db}
}

//Create insert user on database
func (u Users) Create(user models.Users) (uint64, error) {
  return 0, nil
}
