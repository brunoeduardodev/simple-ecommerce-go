package user_gorm

import (
	"gorm.io/gorm"
)

type UserGormModel struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique,index:user_email_idx"`
	Password string
}
