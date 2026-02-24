package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"` // Don't return password in JSON
	AvatarURL string `json:"avatar_url"`
	Provider string `json:"provider"` // "email" or "github"
}
