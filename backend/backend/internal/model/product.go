package model

type Product struct {
	ProductID int `json:"id" gorm:"primaryKey"`
	//	CategoryID  int      `json:"category_id" gorm:"index"`
	Category    Category `gorm:"foreignkey:CategoryID"`
	ShopID      int      `json:"shopID" gorm:"index"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Stock       string   `json:"stock"`
	Image       string   `json:"image"`
}

type Category struct {
	CategoryID int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"category"`
	Image      string `json:"categoryImage"`
}
