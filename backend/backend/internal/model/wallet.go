package model

type Wallet struct {
	ID     int  `json:"id" gorm:"primary_key"`
	UserID int  `json:"user_id"`
	User   User `json:"-" gorm:"forignkey:UserID"`
	Amount int  `json:"amount" gorm:"default:0"`
}
