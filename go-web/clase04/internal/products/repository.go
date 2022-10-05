package products

import (
	"fmt"

	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/store"
)

type Product struct {
	ID        int     `json:"id" `
	Name      string  `json:"name" `
	Price     float64 `json:"price" `
	Quality   int     `json:"quality"`
	Published bool    `json:"published"`
}

var products []Product

type Repository interface {
	Store(id int, name string, price float64, quality int, published bool) (Product, error)
	GetAll() ([]Product, error)
	LastID() (int, error)
	Update(id int, name string, price float64, quality int, published bool) (Product, error)
	UpdatePrice(id int, price float64) (Product, error)
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

func (repo *repository) Store(id int, name string, price float64, quality int, published bool) (Product, error) {
	var products []Product
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	product := Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Quality:   quality,
		Published: published,
	}

	products = append(products, product)
	if err := repo.db.Write(products); err != nil {
		return Product{}, err
	}

	return product, nil
}

func (repo *repository) LastID() (int, error) {
	var products []Product
	err := repo.db.Read(&products)
	if err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].ID, nil
}

func (repo *repository) GetAll() ([]Product, error) {
	var products []Product
	err := repo.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *repository) Update(id int, name string, price float64, quality int, published bool) (Product, error) {
	product := Product{
		Name:      name,
		Price:     price,
		Quality:   quality,
		Published: published,
	}

	for i := range products {
		if products[i].ID == id {
			product.ID = id
			products[i] = product
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product with id %d not found", id)
}

func (repo *repository) Delete(id int) error {
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with id %d not found", id)
}

func (repo *repository) UpdatePrice(id int, price float64) (Product, error) {
	for i := range products {
		if products[i].ID == id {
			products[i].Price = price
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("product with id %d not found", id)
}
