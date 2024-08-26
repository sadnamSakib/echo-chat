package models

// ChatMessage represents a chat message
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   string `json:"userId"`
}
