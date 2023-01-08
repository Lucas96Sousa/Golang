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

// Create insert user on database
func (repository Users) Create(user models.Users) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil

}
