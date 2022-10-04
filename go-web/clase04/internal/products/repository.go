package products

type Product struct {
	ID        int     `json:"id" `
	Name      string  `json:"name" `
	Price     float64 `json:"price" `
	Quality   int     `json:"quality"`
	Published bool    `json:"published"`
}

var products []Product
var lastID int

type Repository interface {
	Store(id int, name string, price float64, quality int, published bool) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) Store(id int, name string, price float64, quality int, published bool) (Product, error) {
	product := Product{
		ID:        lastID,
		Name:      name,
		Price:     price,
		Quality:   quality,
		Published: published,
	}
	products = append(products, product)
	lastID = product.ID
	return product, nil
}

func (repo *repository) LastID() (int, error) {
	return lastID, nil
}
