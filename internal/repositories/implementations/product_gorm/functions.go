package product_gorm

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/models"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories"
	"gorm.io/gorm"
)

type ProductRepositoryGormImplementation struct {
	db *gorm.DB
}

func (r ProductRepositoryGormImplementation) Create(input repositories.CreateProductInput) error {
	product := ProductGormModel{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	result := r.db.Create(&product)

	return result.Error
}

func (r ProductRepositoryGormImplementation) List() ([]models.Product, error) {
	var dbProducts []ProductGormModel
	result := r.db.Find(&dbProducts)

	return dbProductsToModel(dbProducts), result.Error

}

func NewProductRepositoryGormIplementation(db *gorm.DB) ProductRepositoryGormImplementation {
	return ProductRepositoryGormImplementation{db: db}

}
