package models

import (
	"errors"
	"strings"
	"time"
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

	user.lint()
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

	if stage == "register" && user.Password == "" {
		return errors.New("Password is obrigatory and not empty")
	}

	return nil
}

func (user *Users) lint() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
