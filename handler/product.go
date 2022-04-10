package handler

import (
	"net/http"
	"skripsi-product-service/helper"
	"skripsi-product-service/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts()
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
		response := helper.APIResponse("Campaign Not Found", http.StatusNotFound, "success", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success get campaign detail", http.StatusOK, "success", product.FormatProduct(productDetail))
	c.JSON(http.StatusOK, response)
}
