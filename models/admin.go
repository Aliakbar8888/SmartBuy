package models

type Admin struct {
	ID           uint `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

