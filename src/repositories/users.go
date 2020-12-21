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
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, username, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}