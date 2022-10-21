package products

import (
	"fmt"

	"github.com/tonyhuang07/web-server/go-web/clase04/internal/domain"
	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/store"
)

type Repository interface {
	Store(id int, name string, price float64, quality int, published bool) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	LastID() (int, error)
	Update(id int, name string, price float64, quality int, published bool) (*domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Store(id int, name string, price float64, quality int, published bool) (domain.Product, error) {
	var products []domain.Product
	err := repo.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}
	product := domain.Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Quality:   quality,
		Published: published,
	}

	products = append(products, product)

	err = repo.db.Write(&products)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repo *repository) LastID() (int, error) {
	var products []domain.Product
	err := repo.db.Read(&products)
	if err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].ID, nil
}

func (repo *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := repo.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) UpdateName(id int, nombre string) (domain.Product, error) {
	var updated bool = false
	var producp []domain.Product
	if err := r.db.Read(&producp); err != nil {
		return domain.Product{}, err
	}

	var product *domain.Product
	for _, value := range producp {
		if value.ID == id {
			value.Name = nombre
			product = &value
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("domain.product id %d not found", id)
	}

	if err := r.db.Write(&producp); err != nil {
		return domain.Product{}, err
	}

	return *product, nil
}

func (r *repository) Update(id int, name string, price float64, quality int, published bool) (*domain.Product, error) {
	var updated bool = false
	var producp []*domain.Product
	if err := r.db.Read(&producp); err != nil {
		return nil, err
	}

	var product *domain.Product
	for _, value := range producp {
		if value.ID == id {
			value.Name = name
			value.Price = price
			value.Quality = quality
			value.Published = published
			product = value
			updated = true
		}
	}

	if !updated {
		return nil, fmt.Errorf("domain.product id %d not found", id)
	}

	if err := r.db.Write(&producp); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var indice int

	var p []*domain.Product
	if err := r.db.Read(&p); err != nil {
		return err
	}
	for value := range p {
		if p[value].ID == id {
			indice = value
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("domain.product id %d not exist", id)
	}

	p = append(p[:indice], p[indice+1:]...)
	if err := r.db.Write(&p); err != nil {
		return err
	}

	return nil
}
