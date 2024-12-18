package model

import "time"

type Payment struct {
	PaymentID     int       `json:"payment_id" gorm:"primaryKey"`
	OrderID       int       `json:"order_id" gorm:"foreignKey"`
	WalletID      int       `json:"wallet_id" gorm:"foreignKey:WalletID"`
	Amount        float64   `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	TransactionID string    `json:"transaction_id"`

	Order  Order  `json:"order"`
	Wallet Wallet `json:"wallet"`
}
