package music

import (
	"context"
	"errors"
	"fmt"
	"musicadviser/internal/oops"
)

type AppService struct {
	store Store
}

func NewAppService(s Store) *AppService {
	return &AppService{
		store: s,
	}
}

func (s *AppService) Products(ctx context.Context) ([]Product, error) {
	products, err := s.store.LoadProducts(ctx)
	if err != nil {
		if errors.Is(err, oops.ErrNoData) {
			//..
		}
		return nil, fmt.Errorf("firdge.Products() error: %w", err)
	}

	return products, nil
}

func (s *AppService) Place(ctx context.Context, product Product) (id string, err error) {
	id, err = s.store.SaveProduct(ctx, product)
	if err != nil {
		return "", err
	}

	return id, nil
}


func (s *AppService) GetAllUserBands(ctx context.Context) (UserBandsResponse, error) {
	bands, err := s.store.GetAllUserBands(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user bands: %w", err)
	}
	return bands, nil
}
