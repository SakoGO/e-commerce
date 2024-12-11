package model

type Product struct {
	ProductID int      `json:"id" gorm:"primaryKey"`
	Category  Category `gorm:"foreignkey:CategoryID"`
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	Stock     int      `json:"stock"`
	Price     int      `json:"price"`
}

type Category struct {
	CategoryID int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"category"`
	Image      string `json:"categoryImage"`
}
