package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tulara/coffeeshop/domain"
	"github.com/tulara/coffeeshop/store"
)

func TestHandleCreateCafe(t *testing.T) {
	t.Run("test idempotent PUT - create same cafe twice", func(t *testing.T) {
		s := store.NewMemoryStore()
		h := NewHandler(s)

		cafe := domain.Cafe{
			Name: "Idempotent Cafe",
			ID:   "idempotent-1",
		}

		reqBody, err := json.Marshal(cafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		// First PUT - should return 201 Created
		req1 := httptest.NewRequest(http.MethodPut, "/cafes", bytes.NewBuffer(reqBody))
		req1.Header.Set("Content-Type", "application/json")
		w1 := httptest.NewRecorder()

		h.HandleCreateCafe(w1, req1)

		if w1.Code != http.StatusCreated {
			t.Errorf("First PUT: Expected status %d, got %d", http.StatusCreated, w1.Code)
		}

		// Second PUT with same data - should return 200 OK (overwrites, but idempotent)
		req2 := httptest.NewRequest(http.MethodPut, "/cafes", bytes.NewBuffer(reqBody))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()

		h.HandleCreateCafe(w2, req2)

		if w2.Code != http.StatusOK {
			t.Errorf("Second PUT (idempotent override): Expected status %d, got %d", http.StatusOK, w2.Code)
		}

		// Verify cafe still exists and is unchanged (idempotent - same data = same result)
		retrievedCafe := s.GetCafe("idempotent-1")
		if retrievedCafe == nil {
			t.Fatal("Expected cafe to exist, but it was nil")
		}
		if retrievedCafe.Name != "Idempotent Cafe" {
			t.Errorf("Expected cafe name 'Idempotent Cafe', got '%s'", retrievedCafe.Name)
		}
		if retrievedCafe.ID != "idempotent-1" {
			t.Errorf("Expected cafe ID 'idempotent-1', got '%s'", retrievedCafe.ID)
		}
	})

	t.Run("test PUT override - update cafe with different data", func(t *testing.T) {
		s := store.NewMemoryStore()
		h := NewHandler(s)

		// Create initial cafe
		originalCafe := domain.Cafe{
			Name: "Original Name",
			ID:   "update-test-1",
		}

		reqBody1, err := json.Marshal(originalCafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		req1 := httptest.NewRequest(http.MethodPut, "/cafes", bytes.NewBuffer(reqBody1))
		req1.Header.Set("Content-Type", "application/json")
		w1 := httptest.NewRecorder()

		h.HandleCreateCafe(w1, req1)

		if w1.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, w1.Code)
		}

		// Update with different data
		updatedCafe := domain.Cafe{
			Name: "Updated Name",
			ID:   "update-test-1", // Same ID
		}

		reqBody2, err := json.Marshal(updatedCafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		req2 := httptest.NewRequest(http.MethodPut, "/cafes", bytes.NewBuffer(reqBody2))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()

		h.HandleCreateCafe(w2, req2)

		if w2.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w2.Code)
		}

		// Verify cafe was updated
		retrievedCafe := s.GetCafe("update-test-1")
		if retrievedCafe == nil {
			t.Fatal("Expected cafe to exist, but it was nil")
		}
		if retrievedCafe.Name != "Updated Name" {
			t.Errorf("Expected cafe name 'Updated Name', got '%s'", retrievedCafe.Name)
		}
	})

	t.Run("test PUT with missing ID", func(t *testing.T) {
		s := store.NewMemoryStore()
		h := NewHandler(s)

		cafe := domain.Cafe{
			Name: "Cafe Without ID",
			ID:   "", // Missing ID
		}

		reqBody, err := json.Marshal(cafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		req := httptest.NewRequest(http.MethodPut, "/cafes", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.HandleCreateCafe(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}

		body := w.Body.String()
		if !strings.Contains(body, "Cafe ID is required") {
			t.Errorf("Expected error message about missing ID, got: %s", body)
		}
	})
}
