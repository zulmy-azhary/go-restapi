package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email     string         `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" validate:"required,min=3"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string    `json:"token"`
	User  *UserInfo `json:"user"`
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "users"
}
