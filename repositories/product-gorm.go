package repositories

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/models"

	"gorm.io/gorm"
)

type ProductRepositoryGormImplementation struct {
	db *gorm.DB
}

func (r ProductRepositoryGormImplementation) Create(input CreateProductInput) error {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	result := r.db.Create(&product)

	return result.Error
}

func (r ProductRepositoryGormImplementation) List() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)

	return products, result.Error

}

func NewProductRepositoryGormIplementation(db *gorm.DB) ProductRepositoryGormImplementation {
	return ProductRepositoryGormImplementation{db: db}

}
