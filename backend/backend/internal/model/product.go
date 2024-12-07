package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID       int     `json:"id" gorm:"primaryKey"`
	ProductName     string  `json:"product_name"`
	Description     string  `json:"description"`
	Image           string  `json:"image"`
	Stock           int     `json:"stock"`
	Price           float64 `json:"price"`
	QuantityInStock int     `json:"quantity_in_stock"`
	CategoryID      int     `json:"category_id" gorm:"foreignKey:CategoryID"`
	ShopID          int     `json:"shop_id" gorm:"foreignKey:ShopID"`

	Category Category `json:"category"`
	Shop     Shop     `json:"shop"`
}
