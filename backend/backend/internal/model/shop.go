package model

type Shop struct {
	ShopID          int       `json:"id" gorm:"primary_key"`
	ShopName        string    `json:"name" gorm:"unique,not null"`
	ShopDescription string    `json:"shop_description"`
	OwnerID         int       `json:"ownerID"`
	Products        []Product `json:"products" gorm:"foreignKey:ShopID"`

	Owner *User `json:"owner" gorm:"foreignKey:OwnerID"`
}
