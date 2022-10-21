package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonyhuang07/web-server/go-web/clase04/pkg/store"
)

func TestUpdateName(t *testing.T) {
	// Arrange
	mockStorage := &store.MockStorage{}
	repo := NewRepository(mockStorage)
	_, err := repo.Store(1, "before test", 1.0, 1, true)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Act
	product, err := repo.UpdateName(1, "after test")
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "after test", product.Name)
}
