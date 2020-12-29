package models

import "time"

// Post - represents a post to the social media feed
type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorID,omitempty"`
	AuthorUserName string    `json:"authorUserName,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
