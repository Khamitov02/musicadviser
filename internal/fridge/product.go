package fridge

import (
	"context"
)

type Product struct {
	ID    string
	Name  string
	Count uint
}

type Service interface {
	Products(ctx context.Context) ([]Product, error)
	Place(ctx context.Context, product Product) (id string, err error)
}

type Store interface {
	LoadProducts(ctx context.Context) ([]Product, error)
	SaveProduct(ctx context.Context, product Product) (id string, err error)
}
