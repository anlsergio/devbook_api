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

// GetPostbyID - Get a specific Post from the database based on the ID
func (repository Posts) GetPostbyID(postID uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.username FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// GetPosts - Get all posts from users that a given user is following along with its own posts
func (repository Posts) GetPosts(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.username FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}