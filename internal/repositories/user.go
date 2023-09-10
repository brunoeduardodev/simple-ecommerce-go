package repositories

import "github.com/brunoeduardodev/simple-ecommerce-go/internal/models"

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRepository interface {
	Create(input CreateUserInput) error
	FindByEmail(email string) (*models.User, error)
}
