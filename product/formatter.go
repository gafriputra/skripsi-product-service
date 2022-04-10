package product

import "time"

type ProductFormatter struct {
	ID          int       `json:"id"`
	CaregoryID  int       `json:"user_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Status      bool      `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func FormatProduct(product Product) ProductFormatter {
	return ProductFormatter{
		ID:          product.ID,
		CaregoryID:  product.CaregoryID,
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       product.Price,
		Status:      product.Status,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productsFormatter = append(productsFormatter, FormatProduct(product))
	}

	return productsFormatter
}