package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique; not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	UserID   int    `json:"user_id"`
}
