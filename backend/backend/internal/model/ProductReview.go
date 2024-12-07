package model

import "time"

type ProductReview struct {
	ReviewID   int       `json:"review_id" gorm:"primaryKey"`
	ProductID  int       `json:"product_id" gorm:"foreignKey:ProductID"`
	UserID     int       `json:"user_id" gorm:"foreignKey:UserID"`
	Rating     float64   `json:"rating"`
	ReviewText string    `json:"review_text"`
	ReviewDate time.Time `json:"review_date"`

	Product Product `json:"product"`
	User    User    `json:"user"`
}
