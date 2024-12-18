package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductID   int             `json:"id" gorm:"primaryKey"`
	ShopID      int             `json:"shop_id" gorm:"index"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       string          `json:"price"`
	Stock       string          `json:"stock"`
	Image       string          `json:"image"`
	CategoryID  int             `json:"category_id"`
	Category    Category        `gorm:"foreignkey:CategoryID"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}
