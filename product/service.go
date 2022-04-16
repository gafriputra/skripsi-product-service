package product

type Service interface {
	GetProducts(limit int, offset int) ([]Product, error)
	GetProduct(input GetProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(limit int, offset int) ([]Product, error) {
	return s.repository.GetProducts(limit, offset)
}

func (s *service) GetProduct(input GetProductDetailInput) (Product, error) {
	return s.repository.GetProduct(input.Slug)
}
