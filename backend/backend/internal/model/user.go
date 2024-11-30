package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserID   uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" valid:"required" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" valid:"required,email" gorm:"uniqueIndex;not null"`
	Password string `json:"password,omitempty" valid:"required,length(6|20)"`
	Phone    string `json:"phone" valid:"required,numeric,matches(^[0-9]{11}$)"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
	IsAdmin  bool   `json:"isAdmin" gorm:"default:false"`
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
