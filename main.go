package main

import (
	"fmt"

	"github.com/brunoeduardodev/simple-ecommerce-go/controllers"
	"github.com/brunoeduardodev/simple-ecommerce-go/models"
	"github.com/brunoeduardodev/simple-ecommerce-go/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting server...")

	db, err := gorm.Open(sqlite.Open("local.db"))
	db.AutoMigrate(&models.Product{})

	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	productRepository := repositories.NewProductRepositoryGormIplementation(db)
	productController := controllers.NewProductController(productRepository)

	router.GET("/health", controllers.Health)
	router.GET("/products", productController.List)
	router.POST("/products", productController.Create)

	router.Run(":4000")

}
