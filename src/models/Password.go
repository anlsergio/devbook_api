package models

// Password - represents the structure of a update password request
type Password struct {
	New string `json:"new"`
	Current string `json:"current"`
}