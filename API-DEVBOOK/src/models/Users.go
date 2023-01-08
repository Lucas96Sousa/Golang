package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Represent users
type Users struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Prepare calls methods validate and lint for users recieved
func (user *Users) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.lint(stage); err != nil {
		return err
	}
	return nil
}

func (user *Users) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is obrigatory and not empty")
	}

	if user.Nick == "" {
		return errors.New("Nick is obrigatory and not empty")
	}

	if user.Email == "" {
		return errors.New("Email is obrigatory and not empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("Password is obrigatory and not empty")
	}

	return nil
}

func (user *Users) lint(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
