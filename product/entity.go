package product

import (
	"skripsi-product-service/document"
	"skripsi-product-service/gallery"
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
	Galleries   []gallery.Gallery
	Documents   []document.Document
}
