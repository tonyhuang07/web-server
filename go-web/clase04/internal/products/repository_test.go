package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	Products []Product
}

func (s *StubStore) Read(v interface{}) error {
	*v.(*[]Product) = s.Products
	return nil
}

func (s *StubStore) Write(v interface{}) error {
	s.Products = v.([]Product)
	return nil
}

type MockStore struct {
	Products         []Product
	UpdateNameCalled bool
}

func (m *MockStore) Read(v interface{}) error {
	m.UpdateNameCalled = true
	*v.(*[]Product) = m.Products
	return nil
}

func (m *MockStore) Write(v interface{}) error {
	m.Products = v.([]Product)
	return nil
}

func TestGetAll(t *testing.T) {
	products := []Product{
		{
			ID:        1,
			Name:      "test",
			Price:     1.0,
			Quality:   1,
			Published: true,
		},

		{
			ID:        2,
			Name:      "test",
			Price:     1.0,
			Quality:   1,
			Published: true,
		},
	}
	stubStore := &StubStore{Products: products}
	repo := NewRepository(stubStore)

	got, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, products, got)
}

func TestUpdateName(t *testing.T) {
	products := []Product{
		{
			ID:        1,
			Name:      "Before Update",
			Price:     1.0,
			Quality:   1,
			Published: true,
		},
	}

	mockStore := &MockStore{Products: products}
	repo := NewRepository(mockStore)

	var productsBefore []Product
	mockStore.Read(&productsBefore)
	assert.Equal(t, "Before Update", productsBefore[0].Name)

	repo.UpdateName(1, "After Update")

	var productsAfter []Product
	mockStore.Read(&productsAfter)
	assert.Equal(t, "After Update", productsAfter[0].Name)
	assert.True(t, mockStore.UpdateNameCalled)

}
