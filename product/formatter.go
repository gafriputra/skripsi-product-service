package product

import (
	"skripsi-product-service/document"
	"skripsi-product-service/gallery"
	"time"
)

type ProductFormatter struct {
	ID          int                          `json:"id"`
	CaregoryID  int                          `json:"user_id"`
	Name        string                       `json:"name"`
	Slug        string                       `json:"slug"`
	Description string                       `json:"description"`
	Price       int                          `json:"price"`
	Status      bool                         `json:"status"`
	UpdatedAt   time.Time                    `json:"updated_at"`
	CreatedAt   time.Time                    `json:"created_at"`
	DeletedAt   time.Time                    `json:"deleted_at"`
	Galleries   []gallery.GalleryFormatter   `json:"galleries"`
	Documents   []document.DocumentFormatter `json:"documents"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
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

	Galleries := []gallery.GalleryFormatter{}
	for _, img := range product.Galleries {
		Galleries = append(Galleries, gallery.GalleryFormatter{
			Image:     img.Image,
			IsDefault: img.IsDefault,
		})
	}
	formatter.Galleries = Galleries

	documents := []document.DocumentFormatter{}
	for _, doc := range product.Documents {
		documents = append(documents, document.DocumentFormatter{
			Name:         doc.Name,
			DocumentLink: doc.DocumentLink,
			Type:         doc.Type,
		})
	}
	formatter.Documents = documents

	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productsFormatter = append(productsFormatter, FormatProduct(product))
	}

	return productsFormatter
}
