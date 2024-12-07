package model

import "time"

type Wallet struct {
	WalletID    int       `json:"id" gorm:"primary_key"`
	UserID      int       `json:"user_id" gorm:"foreignKey:UserID"`
	Balance     int       `json:"balance" gorm:"default:0"`
	LastUpdated time.Time `json:"last_updated"`

	User User `json:"user"`
}
