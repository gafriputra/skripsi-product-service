package product

import "gorm.io/gorm"

type Repository interface {
	GetProducts(limit int, offset int) ([]Product, error)
	GetProduct(slug string) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts(limit int, offset int) ([]Product, error) {
	var products []Product
	err := r.db.Preload("Images").Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *repository) GetProduct(slug string) (Product, error) {
	var product Product
	err := r.db.Preload("Images").Preload("Documents").Where("slug = ?", slug).Find(&product).Error
	return product, err
}
