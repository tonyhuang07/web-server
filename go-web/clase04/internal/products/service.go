package products

type Service interface {
	Store(name string, price float64, quality int, published bool) (Product, error)
	GetAll() ([]Product, error)
	Update(id int, name string, price float64, quality int, published bool) (Product, error)
	UpdatePrice(id int, price float64) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Store(name string, price float64, quality int, published bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++
	product, err := s.repository.Store(lastID, name, price, quality, published)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Update(id int, name string, price float64, quality int, published bool) (Product, error) {
	return s.repository.Update(id, name, price, quality, published)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdatePrice(id int, price float64) (Product, error) {
	return s.repository.UpdatePrice(id, price)
}
