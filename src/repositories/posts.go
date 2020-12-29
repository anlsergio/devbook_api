package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts - represents a repository of posts
type Posts struct {
	db *sql.DB
}

// NewPostsRepository - creates a repository of posts
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create - saves a post into the database
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(postID), nil
}