package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID    int             `json:"id" gorm:"primaryKey"`
	Username  string          `json:"username" valid:"required" gorm:"type:varchar(256);unique;not null"`
	Email     string          `json:"email" valid:"required;email" gorm:"type:varchar(256);unique;not null"`
	Password  string          `json:"password,omitempty" valid:"type:varchar(255);required,length(6|20)"`
	Phone     string          `json:"phone" valid:"required,matches(^[0-9]{11}$)" gorm:"type:varchar(256);not null;unique"`
	Role      string          `json:"role" gorm:"varchar(255);default:customer"`
	Blocked   bool            `json:"blocked" gorm:"default:false"`
	WalletID  int             `json:"wallet_id"`
	Wallet    Wallet          `gorm:"foreignkey:WalletID"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
