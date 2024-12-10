package music_test

import (
	"context"
	"musicadviser/internal/music"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type memoryStore struct {
	products []music.Product
}

func (m *memoryStore) LoadProducts(ctx context.Context) ([]music.Product, error) {
	return m.products, nil
}

func (m *memoryStore) SaveProduct(ctx context.Context, product music.Product) (string, error) {
	m.products = append(m.products, product)
	return "test-id", nil
}

func (m *memoryStore) GetAllUserBands(ctx context.Context) (music.UserBandsResponse, error) {
	response := make(music.UserBandsResponse)
	for _, p := range m.products {
		response[p.UserID] = append(response[p.UserID], p.BandName)
	}
	return response, nil
}

func TestAppService_GetAllUserBands(t *testing.T) {
	// Initialize test data
	store := &memoryStore{
		products: []music.Product{
			{ID: "1", UserID: "user1", BandName: "Band A"},
			{ID: "2", UserID: "user1", BandName: "Band B"},
			{ID: "3", UserID: "user2", BandName: "Band C"},
		},
	}

	service := music.NewAppService(store)

	// Test getting all user bands
	got, err := service.GetAllUserBands(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := music.UserBandsResponse{
		"user1": []string{"Band A", "Band B"},
		"user2": []string{"Band C"},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetAllUserBands() mismatch (-want +got):\n%s", diff)
	}
}

func TestAppService_Place(t *testing.T) {
	store := &memoryStore{}
	service := music.NewAppService(store)

	testProduct := music.Product{
		UserID:   "user1",
		BandName: "Band A",
	}

	id, err := service.Place(context.Background(), testProduct)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != "test-id" {
		t.Errorf("expected id 'test-id', got %s", id)
	}

	// Verify the product was stored
	bands, err := service.GetAllUserBands(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(bands["user1"]) != 1 || bands["user1"][0] != "Band A" {
		t.Errorf("expected band 'Band A' for user1, got %v", bands["user1"])
	}
}
