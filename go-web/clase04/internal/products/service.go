package products

import (
	"github.com/tonyhuang07/web-server/go-web/clase04/internal/domain"
)

type Service interface {
	Store(name string, price float64, quality int, published bool) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Update(id int, name string, price float64, quality int, published bool) (*domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Store(name string, price float64, quality int, published bool) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}
	lastID++
	product, err := s.repository.Store(lastID, name, price, quality, published)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Update(id int, name string, price float64, quality int, published bool) (*domain.Product, error) {
	return s.repository.Update(id, name, price, quality, published)
}

func (s *service) UpdateName(id int, name string) (domain.Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
