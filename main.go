package main

import (
	"fmt"

	"github.com/brunoeduardodev/simple-ecommerce-go/docs"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/controllers"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories/implementations/product_gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Simple Ecommerce Go
//	@version		1.0
//	@description	Project to study golang
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Bruno
//	@contact.url	https://github.com/brunoeduardodev

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:4000

//	@securityDefinitions.apiKey
//  @in header
//  @name Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	fmt.Println("Starting server...")

	db, err := gorm.Open(sqlite.Open("local.db"))
	db.AutoMigrate(&product_gorm.ProductGormModel{})

	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	//
	productRepository := product_gorm.NewProductRepositoryGormIplementation(db)
	productController := controllers.NewProductController(productRepository)

	productsRouter := router.Group("/products")
	{
		productsRouter.GET("", productController.List)
		productsRouter.POST("", productController.Create)

	}

	router.GET("/health", controllers.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":4000")

}
