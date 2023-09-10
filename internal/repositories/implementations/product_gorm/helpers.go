package product_gorm

import "github.com/brunoeduardodev/simple-ecommerce-go/internal/models"

func dbProductToModel(dbProduct ProductGormModel) models.Product {
	return models.Product{
		ID:          dbProduct.ID,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Price:       dbProduct.Price,
		CreatedAt:   dbProduct.CreatedAt,
		UpdatedAt:   dbProduct.UpdatedAt,
	}
}

func dbProductsToModel(dbProducts []ProductGormModel) []models.Product {
	var products []models.Product
	for i := 0; i < len(dbProducts); i++ {
		products = append(products, dbProductToModel(dbProducts[i]))
	}

	return products
}
