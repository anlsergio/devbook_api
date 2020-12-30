package models

import (
	"api/src/security"
	"errors"
	"regexp"
	"strings"
	"time"
)

// User represents a user signed up for the social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare - calls the validation and format functions
func (user *User) Prepare(isNew bool) error {
	if err := user.validate(isNew); err != nil {
		return err
	}

	if err := user.format(isNew); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(isNew bool) error {
	if user.Name == "" {
		return errors.New("A value for 'Name' must be provided")
	}

	if user.Username == "" {
		return errors.New("A value for 'Username' must be provided")
	}

	if user.Email == "" {
		return errors.New("A value for 'Email' must be provided")
	}

	emailRegex := regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
	if ! emailRegex.MatchString(user.Email) {
		return errors.New("The value for 'Email' provided is not valid")
	}

	if isNew && user.Password == "" {
		return errors.New("A value for 'Password' must be provided")
	}

	return nil
}

func (user *User) format(isNew bool) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)

	if isNew {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}