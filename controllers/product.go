package controllers

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/repositories"
	"github.com/brunoeduardodev/simple-ecommerce-go/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository repositories.ProductRepository
}

func (controller ProductController) List(ctx *gin.Context) {
	products, err := services.ProductsList(controller.productRepository)

	if err != nil {
		sendError(ctx, 503, "Unable to list products")
		return
	}

	ctx.JSON(200, products)
}

func (controller ProductController) Create(ctx *gin.Context) {
	var input services.CreateProductInput

	ctx.ShouldBind(&input)

	err := services.ProductCreate(controller.productRepository, input)

	if err != nil {
		sendError(ctx, 503, "Unable to create product")
		return
	}

	ctx.Status(200)
}

func NewProductController(repository repositories.ProductRepository) ProductController {
	return ProductController{productRepository: repository}
}
