package models

import (
	"errors"
	"strings"
	"time"
)

// Post - represents a post to the social media feed
type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorID,omitempty"`
	AuthorUsername string    `json:"authorUsername,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

// Prepare - calls the validation and format functions
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("The post's title cannot be blank")
	}

	if post.Content == "" {
		return errors.New("The post needs a content")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}