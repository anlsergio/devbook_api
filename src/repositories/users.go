package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users - represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository - creates a repository of Users
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create - adds a user into the database
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

// Get - is responsible  for fetching all users that meet the name or username filtering
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

// GetByID - is responsible for fetching a single user from the database by ID
func (repository Users) GetByID(userID uint64) (models.User, error) {
	row, err := repository.db.Query("SELECT id, name, username, email, createdAt FROM users WHERE id = ?", userID)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()
	
	var user models.User
	
	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update - changes a user information based on a given ID
func (repository Users) Update(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Username, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// Delete - removes a user from the database based on the ID
func (repository Users) Delete(userID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// GetByEmail - is responsible for fetching a user's ID and Password from the database filtered by its email address
func (repository Users) GetByEmail(email string) (models.User, error) {
	row, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow - Allows a user to follow another
func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow - Allows a user to unfollow another
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// GetFollowers - Fetches all followers from a given user
func (repository Users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.createdAt FROM users u
		INNER JOIN followers f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
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

// GetFollowing - Gets all users that a given user is currently following
func (repository Users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.createdAt FROM users u
		INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, userID)
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