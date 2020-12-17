package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository creates a repository of Users
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create adds a user into the database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}