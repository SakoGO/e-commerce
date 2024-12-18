package model

type Category struct {
	CategoryID int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
}
