package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository Users) Get(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // %nameOrUsername%

	rows, err := repository.db.Query(
		"SELECT id, name, username, email, createdAt FROM users WHERE name LIKE ? or username LIKE ?",
		nameOrUsername, nameOrUsername,
	)

	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}