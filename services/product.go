package services

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/repositories"
	"github.com/brunoeduardodev/simple-ecommerce-go/views"
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
