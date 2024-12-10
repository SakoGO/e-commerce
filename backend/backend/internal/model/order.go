package model

import "time"

type Order struct {
	OrderID       int       `json:"order_id" gorm:"primaryKey"`
	UserID        int       `json:"user_id" gorm:"foreignKey:UserID"`
	WalletID      int       `json:"wallet_id" gorm:"foreignKey:WalletID"`
	OrderDate     time.Time `json:"order_date"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`         // pending, completed, canceled
	PaymentStatus string    `json:"payment_status"` // unpaid, paid
	PaymentMethod string    `json:"payment_method"`

	User   User   `json:"user"`
	Wallet Wallet `json:"wallet"`
	Cart   []Cart `json:"cart"`
}
