package models

import (
	"time"
)

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
