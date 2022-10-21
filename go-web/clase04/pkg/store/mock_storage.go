package store

import (
	"fmt"

	"github.com/tonyhuang07/web-server/go-web/clase04/internal/domain"
)

type MockStorage struct {
	dataMock []domain.Product
	errRead  string
	errWrite string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	products := data.(*[]domain.Product)
	*products = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	products := data.(*[]domain.Product)
	m.dataMock = *products
	return nil
}
