package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int      `json:"id" gorm:"primaryKey"`
	CategoryID  int      `json:"categoryId"`
	Category    Category `json:"-" gorm:"forignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"productName"`
	Image       string   `json:"image"`
	Stock       int      `json:"stock"`
	Price       int      `json:"price"`
}

type Category struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
	Image    string `json:"categoryImage"`
}
