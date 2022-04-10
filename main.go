package main

import (
	"log"
	"skripsi-product-service/handler"
	"skripsi-product-service/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=gafriputra dbname=skripsi port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1/product")

	api.GET("/", productHandler.GetProducts)
	api.GET("/:slug", productHandler.GetProduct)

	router.Run(":7001")
}
