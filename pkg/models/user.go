package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	PassHash []byte `json:"password,omitempty"`
}