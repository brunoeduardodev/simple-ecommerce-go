package views

import (
	"time"

	"github.com/brunoeduardodev/simple-ecommerce-go/models"
)

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	CreatedAt   string `json:"string"`
}

func DbProductToView(product models.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
	}
}

func DbProductsToView(dbProducts []models.Product) []Product {
	products := []Product{}

	for i := 0; i < len(dbProducts); i++ {
		products = append(products, DbProductToView(dbProducts[i]))
	}

	return products
}
