package controllers

import (
	"net/http"

	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository repositories.ProductRepository
}

// @BasePath /
// List Products
// @Summary List all products
// @Description List all products
// @Produce json
// @Success 200 {array} views.Product
// @Router /products [get]
func (controller ProductController) List(ctx *gin.Context) {
	products, err := services.ProductsList(controller.productRepository)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Unable to list products")
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// @BasePath /
// List Products
// @Summary Create Product
// @Description Create a new product
// @Param request body services.CreateProductInput true  "New product input"
// @Produce json
// @Success 200
// @Router /products [post]
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
