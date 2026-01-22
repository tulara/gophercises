package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/tulara/coffeeshop/domain"
	"github.com/tulara/coffeeshop/handler"
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
			ID:   1,
		}

		reqBody, err := json.Marshal(cafe)
		if err != nil {
			t.Fatalf("Failed to marshal cafe: %v", err)
		}

		createResp := DoPUTRequest(t, client, http.MethodPut, server.URL+"/cafes", reqBody)

		if createResp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, createResp.StatusCode)
		}

		// Step 2: Retrieve the cafe
		getReq, err := http.NewRequest(http.MethodGet, server.URL+"/cafes/1", nil)
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
		if retrievedCafe.ID != 1 {
			t.Errorf("Expected retrieved cafe ID 'stumptown-1', got '%d'", retrievedCafe.ID)
		}
	})

	t.Run("create multiple cafes and retrieve each", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}

		cafes := []domain.Cafe{
			{Name: "Blue Bottle", ID: 1},
			{Name: "Intelligentsia", ID: 2},
			{Name: "Counter Culture", ID: 3},
		}

		// Create all cafes
		for _, cafe := range cafes {
			reqBody, err := json.Marshal(cafe)
			if err != nil {
				t.Fatalf("Failed to marshal cafe %d: %v", cafe.ID, err)
			}

			resp := DoPUTRequest(t, client, http.MethodPut, server.URL+"/cafes", reqBody)

			if resp.StatusCode != http.StatusCreated {
				t.Errorf("Failed to create cafe %d: got status %d", cafe.ID, resp.StatusCode)
			}
		}

		// Retrieve each cafe and verify
		for _, expectedCafe := range cafes {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cafes/%s", server.URL, strconv.Itoa(expectedCafe.ID)), nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Failed to retrieve cafe %d: got status %d", expectedCafe.ID, resp.StatusCode)
				resp.Body.Close()
				continue
			}

			var retrievedCafe domain.Cafe
			if err := json.NewDecoder(resp.Body).Decode(&retrievedCafe); err != nil {
				t.Errorf("Failed to unmarshal cafe %d: %v", expectedCafe.ID, err)
				resp.Body.Close()
				continue
			}
			resp.Body.Close()

			if retrievedCafe.Name != expectedCafe.Name {
				t.Errorf("Expected cafe name '%s', got '%s'", expectedCafe.Name, retrievedCafe.Name)
			}
			if retrievedCafe.ID != expectedCafe.ID {
				t.Errorf("Expected cafe ID '%d', got '%d'", expectedCafe.ID, retrievedCafe.ID)
			}
		}
	})

	t.Run("create multiple cafes and paginate", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}

		cafes := []domain.Cafe{
			{Name: "Blue Bottle", ID: 1},
			{Name: "Intelligentsia", ID: 2},
			{Name: "Counter Culture", ID: 3},
		}

		// Create all cafes
		for _, cafe := range cafes {
			reqBody, err := json.Marshal(cafe)
			if err != nil {
				t.Fatalf("Failed to marshal cafe %d: %v", cafe.ID, err)
			}

			resp := DoPUTRequest(t, client, http.MethodPut, server.URL+"/cafes", reqBody)

			if resp.StatusCode != http.StatusCreated {
				t.Errorf("Failed to create cafe %d: got status %d", cafe.ID, resp.StatusCode)
			}
			resp.Body.Close()
		}

		// list 2 from beginning
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cafes?page_size=2", server.URL), nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Failed to retrieve cafes, got status %d", resp.StatusCode)
			resp.Body.Close()
		}

		var retrievedCafes *handler.GetCafesDTO
		if err := json.NewDecoder(resp.Body).Decode(&retrievedCafes); err != nil {
			t.Errorf("Failed to unmarshal cafes: %v", err)
			resp.Body.Close()
		}
		resp.Body.Close()

		if len(retrievedCafes.Data) != 2 {
			t.Fatalf("Fetched %d instead of 2 cafes", len(retrievedCafes.Data))
		}

		// use cursor from previous response to request next page, and list one remaining cafe
		// list 2 from beginning
		id := retrievedCafes.Data[1].ID
		req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cafes?page_size=2&cursor=%d", server.URL, id), nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Failed to retrieve cafes, got status %d", resp.StatusCode)
			resp.Body.Close()
		}

		if err := json.NewDecoder(resp.Body).Decode(&retrievedCafes); err != nil {
			t.Errorf("Failed to unmarshal cafes: %v", err)
			resp.Body.Close()
		}
		resp.Body.Close()

		if len(retrievedCafes.Data) != 1 {
			t.Error("Incorrect number of cafes fetched")
			resp.Body.Close()
		}
	})

	t.Run("test cafe not found", func(t *testing.T) {
		server := setupTestServer()
		defer server.Close()

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, server.URL+"/cafes/999", nil)
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

// NOTE closes body before deserialisation
func DoPUTRequest(t *testing.T, client *http.Client, method string, path string, reqBody []byte) *http.Response {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	resp.Body.Close()
	return resp
}

// setupTestServer creates a test HTTP server with the same routes as main()
func setupTestServer() *httptest.Server {
	s := store.NewMemoryStore()
	mux := setupRoutes(s)
	server := httptest.NewServer(mux)
	return server
}
