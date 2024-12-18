package model

type Wallet struct {
	WalletID int `gorm:"primarykey"`
	Amount   int `json:"amount" gorm:"default:0"`
}
