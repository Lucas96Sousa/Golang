package repositories

import (
	"database/sql"
	"fmt"

	"api/src/models"
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

// Find -> search users  method by name or nick in database
func (repository Users) Find(nameOrNick string) ([]models.Users, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%
	rows, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// FindByID - search by id
func (repository Users) FindByID(ID uint64) (models.Users, error) {
	rows, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.Users{}, err
	}
	defer rows.Close()

	var user models.Users

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.Users{}, err
		}
	}
	return user, nil
}

// UpdateUser -> update user informartion
func (repository Users) UpdateUser(ID uint64, user models.Users) error {
	statement, err := repository.db.Prepare("update users set name =? , email = ?, nick = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, user.Nick, ID); err != nil {
		return err
	}

	return nil
}

// DeleteUser -> remove user from table
func (repository Users) DeleteUser(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return nil
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
