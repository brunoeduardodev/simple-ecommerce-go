package product_gorm

import (
	"gorm.io/gorm"
)

type ProductGormModel struct {
	gorm.Model

	Name        string
	Description string
	Price       uint
}
