package mock

import (
	"context"
	"musicadviser/internal/music"
)

type Mock struct{}

func NewFridge() *Mock {
	return &Mock{}
}

func (m *Mock) GetAllUserBands(ctx context.Context) (music.UserBandsResponse, error) {
	return music.UserBandsResponse{
		"user1": []string{"Band A", "Band B"},
		"user2": []string{"Band C"},
	}, nil
}

func (m *Mock) Place(ctx context.Context, product music.Product) (id string, err error) {
	return "mock-id", nil
}

func (m *Mock) Products(ctx context.Context) ([]music.Product, error) {
	return []music.Product{
		{
			ID:       "1",
			UserID:   "user1",
			BandName: "Band A",
		},
	}, nil
}
