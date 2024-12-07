package mock

import (
	"context"
	"musicadviser/internal/fridge"

	"github.com/google/uuid"
)

type Mock struct{}

func NewFridge() *Mock {
	return &Mock{}
}

func (m *Mock) Products(ctx context.Context) ([]fridge.Product, error) {
	products := []fridge.Product{
		{
			ID:    uuid.New().String(),
			Name:  "Test name",
			Count: 17,
		},
	}

	return products, nil
}

func (m *Mock) Place(ctx context.Context, product fridge.Product) (id string, err error) {
	return uuid.New().String(), nil
}
