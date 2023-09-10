package services

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/views"
)

func ProductsList(repository repositories.ProductRepository) ([]views.Product, error) {
	products, err := repository.List()

	return views.DbProductsToView(products), err
}

type CreateProductInput = repositories.CreateProductInput

func ProductCreate(repository repositories.ProductRepository, input CreateProductInput) error {
	error := repository.Create(input)

	return error
}
