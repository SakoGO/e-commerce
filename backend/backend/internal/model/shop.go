package model

import (
	"gorm.io/gorm"
	"time"
)

type Shop struct {
	ShopID      int             `gorm:"primarykey"`
	Name        string          `json:"name" valid:"required" gorm:"unique;type:varchar(255)"`
	Description string          `json:"description" valid:"required" gorm:"type:varchar(1000)"`
	Email       string          `json:"email" valid:"required;email" gorm:"type:varchar(256);unique;not null"`
	OwnerID     int             `json:"owner_id"`
	Products    []*Product      `gorm:"foreignkey:ProductID"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}
