package music_test

import (
	"bytes"
	"encoding/json"
	"musicadviser/internal/music"
	"musicadviser/internal/music/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/go-cmp/cmp"
)

func TestHandler_GetMusic(t *testing.T) {
	service := mock.NewFridge()
	router := chi.NewRouter()
	h := music.NewHandler(router, service)
	h.Register()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/getMusic", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("body", func(t *testing.T) {
		var got music.UserBandsResponse
		err := json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		want := music.UserBandsResponse{
			"user1": []string{"Band A", "Band B"},
			"user2": []string{"Band C"},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("GET /api/v1/getMusic mismatch: (-want +got)\n%s", diff)
		}
	})
}

func TestHandler_PutMusic(t *testing.T) {
	service := mock.NewFridge()
	router := chi.NewRouter()
	h := music.NewHandler(router, service)
	h.Register()

	tests := []struct {
		name       string
		request    music.MusicRequest
		wantStatus int
	}{
		{
			name: "valid request",
			request: music.MusicRequest{
				UserID: "user1",
				Bands:  []string{"Band A", "Band B"},
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "empty bands list",
			request: music.MusicRequest{
				UserID: "user1",
				Bands:  []string{},
			},
			wantStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/v1/putMusic", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("handler returned wrong status code: want %d, got %d", tt.wantStatus, rr.Code)
			}
		})
	}
}
