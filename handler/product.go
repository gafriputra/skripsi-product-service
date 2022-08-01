package handler

import (
	"net/http"
	"skripsi-product-service/helper"
	"skripsi-product-service/product"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit
	products, err := h.service.GetProducts(limit, offset)
	if err != nil {
		response := helper.APIResponse("Failed get products", http.StatusBadGateway, "error", err.Error())
		c.JSON(http.StatusBadGateway, response)
		return
	}

	response := helper.APIResponse("Success get products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	var input product.GetProductDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.service.GetProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if productDetail.ID == 0 {
		response := helper.APIResponse("Product Not Found", http.StatusNotFound, "success", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success get product detail", http.StatusOK, "success", product.FormatProduct(productDetail))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DummyBenchmark(c *gin.Context) {
	var input product.DummyBenchmark
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	benchmark := input.Benchmark
	sleep := 0
	switch benchmark {
	case "monolithic":
		time.Sleep(1500 * time.Millisecond)
		sleep = 1500
	case "microservice":
		time.Sleep(1000 * time.Millisecond)
		sleep = 1000
	case "microservice-scale":
		time.Sleep(500 * time.Millisecond)
		sleep = 500
	}
	c.JSON(http.StatusOK, sleep)
}
