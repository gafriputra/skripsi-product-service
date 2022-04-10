package product

type Service interface {
	GetProducts() ([]Product, error)
	GetProduct(input GetProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts() ([]Product, error) {
	return s.repository.GetProducts()
}

func (s *service) GetProduct(input GetProductDetailInput) (Product, error) {
	return s.repository.GetProduct(input.Slug)
}
