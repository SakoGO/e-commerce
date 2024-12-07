package model

type Category struct {
	CategoryID         int       `json:"category_id" gorm:"primaryKey"`
	CategoryName       string    `json:"category_name"`
	ParentCategoryID   int       `json:"parent_category_id,omitempty"`
	ParentCategoryName *Category `json:"parent_category_name,omitempty" gorm:"foreignKey:ParentCategoryID"`
	Image              string    `json:"categoryImage"`
}
