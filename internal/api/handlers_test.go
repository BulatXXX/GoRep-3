package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"singularity.com/pz3-http/internal/storage"
)

// helper to build mux with our routes (so r.PathValue works for {id})
func buildMux(h *Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", h.GetTask)
	return mux
}

func TestCreateTask_OK(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)
	mux := buildMux(h)

	body := []byte(`{"title":"Test task"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d, body=%s", rr.Code, rr.Body.String())
	}

	var got storage.Task
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.Title != "Test task" || got.ID == 0 || got.Done {
		t.Errorf("unexpected task: %+v", got)
	}
}

func TestGetTask_OK(t *testing.T) {
	store := storage.NewMemoryStore()
	// pre-create a task in store
	created := store.Create("Read book")

	h := NewHandlers(store)
	mux := buildMux(h)

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, body=%s", rr.Code, rr.Body.String())
	}

	var got storage.Task
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.ID != created.ID || got.Title != created.Title || got.Done != created.Done {
		t.Errorf("unexpected task, want=%+v got=%+v", *created, got)
	}
}
