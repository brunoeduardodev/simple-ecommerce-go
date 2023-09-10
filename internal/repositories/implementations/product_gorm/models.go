package product_gorm

import (
	"time"

	"gorm.io/gorm"
)

type ProductGormModel struct {
	gorm.Model

	Name        string
	Description string
	Price       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
