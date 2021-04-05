package models

// AuthData - designed to contain both the authentication token and the authenticated user ID as parsed data.
type AuthData struct {
	ID string `json:"id"`
	Token string `json:"token"`
}