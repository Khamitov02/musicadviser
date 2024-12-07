package sqlite

import (
	"context"
	"sandbox/internal/fridge"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) LoadProducts(ctx context.Context) ([]fridge.Product, error) {
	return nil, nil
}

func (s *Storage) SaveProduct(ctx context.Context, product fridge.Product) (id string, err error) {
	return "", nil
}
