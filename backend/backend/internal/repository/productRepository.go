package repository

import (
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

//TODO: Что надо делать с продуктом в бд... :
//TODO:
