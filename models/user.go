package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Nik       string    `json:"nik"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Nik      string `json:"Nik" binding:"required"`
	Password string `json:"password" binding:"required,min=3"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
