package model

import (
	"time"
)

type User struct {
	UserID    int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" valid:"required" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email     string    `json:"email" valid:"required,email" gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	Password  string    `json:"password,omitempty" valid:"type:varchar(255);required,length(6|20)"`
	Phone     string    `json:"phone" valid:"required,matches(^[0-9]{11}$)" gorm:"type:varchar(15);not null;uniqueIndex"`
	Blocked   bool      `json:"blocked" gorm:"default:false"`
	Role      string    `json:"role" gorm:"varchar(255);default:customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`

	Wallet *Wallet `json:"wallet" gorm:"foreignKey:UserID;references:UserID"`
}
