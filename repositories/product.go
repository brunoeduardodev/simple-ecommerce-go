package repositories

import "github.com/brunoeduardodev/simple-ecommerce-go/models"

type CreateProductInput struct {
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
}

type ProductRepository interface {
	Create(input CreateProductInput) error
	List() ([]models.Product, error)
}
