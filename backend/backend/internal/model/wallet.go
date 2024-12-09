package model

import "time"

type Wallet struct {
	WalletID    int       `json:"id" gorm:"primary_key"`
	UserID      int       `json:"user_id"`
	Balance     float64   `json:"balance" gorm:"default:0"`
	LastUpdated time.Time `json:"last_updated" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
