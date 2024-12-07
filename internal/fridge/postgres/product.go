package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"musicadviser/internal/fridge"
	"musicadviser/internal/oops"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    sql.NullString
	Name  sql.NullString
	Count sql.NullInt64
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) LoadProducts(ctx context.Context) ([]fridge.Product, error) {
	// данные, которые будем получать, будем складировать в Product
	// потом валидировать и переносить в то, что ожидает бизнес.

	// Product -> fridgle.Product происходит на уровне работы БД.

	//
	// TODO: написать перекладывание Product -> fridgle.Product с валидацией
	//

	return nil, fmt.Errorf("postgres.LoadProducts() error: %w", oops.ErrNoData)
}

func (s *Storage) SaveProduct(ctx context.Context, product fridge.Product) (id string, err error) {
	return "", nil
}
