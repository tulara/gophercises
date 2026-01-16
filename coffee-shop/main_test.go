package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tulara/coffeeshop/domain"
	"github.com/tulara/coffeeshop/store"
)

func TestAPIEndToEnd(t *testing.T) {
	t.Run("create and retrieve cafe", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}

		// Step 1: Create a cafe
		cafe := domain.Cafe{
			Name: "Stumptown Coffee",
			ID:   "stumptown-1",
		}

		reqBody, err := json.Marshal(cafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		createReq, err := http.NewRequest(http.MethodPut, server.URL+"/cafes", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		createReq.Header.Set("Content-Type", "application/json")

		createResp, err := client.Do(createReq)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer createResp.Body.Close()

		if createResp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, createResp.StatusCode)
		}

		// Step 2: Retrieve the cafe
		getReq, err := http.NewRequest(http.MethodGet, server.URL+"/cafes/stumptown-1", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		getResp, err := client.Do(getReq)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer getResp.Body.Close()

		if getResp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, getResp.StatusCode)
		}

		var retrievedCafe domain.Cafe
		if err := json.NewDecoder(getResp.Body).Decode(&retrievedCafe); err != nil {
			t.Fatalf("Failed to unmarshal retrieved cafe: %v", err)
		}

		if retrievedCafe.Name != "Stumptown Coffee" {
			t.Errorf("Expected retrieved cafe name 'Stumptown Coffee', got '%s'", retrievedCafe.Name)
		}
		if retrievedCafe.ID != "stumptown-1" {
			t.Errorf("Expected retrieved cafe ID 'stumptown-1', got '%s'", retrievedCafe.ID)
		}
	})

	t.Run("create multiple cafes and retrieve each", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}

		cafes := []domain.Cafe{
			{Name: "Blue Bottle", ID: "blue-bottle-1"},
			{Name: "Intelligentsia", ID: "intelligentsia-1"},
			{Name: "Counter Culture", ID: "counter-culture-1"},
		}

		// Create all cafes
		for _, cafe := range cafes {
			reqBody, err := json.Marshal(cafe)
			if err != nil {
				t.Fatalf("Failed to marshal cafe %s: %v", cafe.ID, err)
			}

			req, err := http.NewRequest(http.MethodPut, server.URL+"/cafes", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				t.Errorf("Failed to create cafe %s: got status %d", cafe.ID, resp.StatusCode)
			}
		}

		// Retrieve each cafe and verify
		for _, expectedCafe := range cafes {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cafes/%s", server.URL, expectedCafe.ID), nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Failed to retrieve cafe %s: got status %d", expectedCafe.ID, resp.StatusCode)
				resp.Body.Close()
				continue
			}

			var retrievedCafe domain.Cafe
			if err := json.NewDecoder(resp.Body).Decode(&retrievedCafe); err != nil {
				t.Errorf("Failed to unmarshal cafe %s: %v", expectedCafe.ID, err)
				resp.Body.Close()
				continue
			}
			resp.Body.Close()

			if retrievedCafe.Name != expectedCafe.Name {
				t.Errorf("Expected cafe name '%s', got '%s'", expectedCafe.Name, retrievedCafe.Name)
			}
			if retrievedCafe.ID != expectedCafe.ID {
				t.Errorf("Expected cafe ID '%s', got '%s'", expectedCafe.ID, retrievedCafe.ID)
			}
		}
	})

	t.Run("test cafe not found", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, server.URL+"/cafes/non-existent", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}

// setupTestServer creates a test HTTP server with the same routes as main()
func setupTestServer() *httptest.Server {
	s := store.NewMemoryStore()
	mux := setupRoutes(s)
	server := httptest.NewServer(mux)
	return server
}
