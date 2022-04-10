package product

import (
	"skripsi-product-service/image"
	"time"
)

type Product struct {
	ID          int
	CaregoryID  int
	Name        string
	Slug        string
	Description string
	Price       int
	Status      bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Images      []image.Image
}
