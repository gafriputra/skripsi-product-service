package main

import (
	"fmt"
	"log"
	"skripsi-product-service/handler"
	"skripsi-product-service/helper"
	"skripsi-product-service/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		helper.Env("DB_HOST"), helper.Env("DB_USER"), helper.Env("DB_PASSWORD"), helper.Env("DB_NAME"), helper.Env("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1/products")

	api.GET("/", productHandler.GetProducts)
	api.GET("/:slug", productHandler.GetProduct)
	api.GET("/benchmark", productHandler.DummyBenchmark)

	router.Run("0.0.0.0:" + helper.Env("APP_PORT"))
}
