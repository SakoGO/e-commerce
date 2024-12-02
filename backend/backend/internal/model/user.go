package model

import "time"

type User struct {
	UserID    int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" valid:"required" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email     string    `json:"email" valid:"required,email" gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	Password  string    `json:"password,omitempty" valid:"type:varchar(255);required,length(6|20)"`
	Phone     string    `json:"phone" valid:"required,matches(^[0-9]{11}$)" gorm:"type:varchar(15);not null;uniqueIndex"`
	Blocked   bool      `json:"blocked" gorm:"default:false"`
	IsAdmin   bool      `json:"isAdmin" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

/*
type Address struct {
	*gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      int    `json:"userId"`
	AddressName string `json:"name" valid:"required"`
	HouseNum    uint   `json:"houseId" valid:"required, numeric"`
	Street      string `json:"street" valid:"required"`
	City        string `json:"city" valid:"required"`
	Region      string `json:"region" valid:"required"`
	Phone       string `json:"phone" valid:"required,numeric,matches(^[0-9]{11}$)"`
	Default     bool   `json:"default" gorm:"default:false"`
}
*/
