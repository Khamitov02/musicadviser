package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"musicadviser/internal/music"

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

func (s *Storage) LoadProducts(ctx context.Context) ([]music.Product, error) {
	// данные, которые будем получать, будем складировать в Product
	// потом валидировать и переносить в то, что ожидает бизнес.

	// Product -> fridgle.Product происходит на уровне работы БД.

	//
	// TODO: написать перекладывание Product -> fridgle.Product с валидацией
	//
	var dbProducts []Product
	query := "SELECT ID, Name, Count FROM products" // create query string

	// Execute query
	err := s.db.SelectContext(ctx, &dbProducts, query)
	if err != nil {
		return nil, fmt.Errorf("failed to select products: %w", err)
	}

	var products []music.Product
	for _, dbProduct := range dbProducts {
		if !dbProduct.ID.Valid {
			return nil, fmt.Errorf("product ID is NULL")
		}
		if !dbProduct.Name.Valid {
			return nil, fmt.Errorf("product Name is NULL")
		}
		if !dbProduct.Count.Valid {
			return nil, fmt.Errorf("product Count is NULL")
		}

		product := music.Product{
			ID:    dbProduct.ID.String,
			Name:  dbProduct.Name.String,
			Count: uint(dbProduct.Count.Int64),
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) SaveProduct(ctx context.Context, product music.Product) (id string, err error) {
	query := `
		INSERT INTO products (Name, Count)
		VALUES ($1, $2)
		RETURNING ID
	`

	var newID string
	err = s.db.QueryRowContext(ctx, query, product.Name, product.Count).Scan(&newID)
	if err != nil {
		return "", fmt.Errorf("failed to insert product: %w", err)
	}

	return newID, nil
}
