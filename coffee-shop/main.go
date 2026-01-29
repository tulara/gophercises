package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tulara/coffeeshop/handler"
	"github.com/tulara/coffeeshop/middleware"
	"github.com/tulara/coffeeshop/store"
)

func main() {
	store := store.NewMemoryStore()
	mux := setupRoutes(store)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	fmt.Println("Listen on :8080")
}

// setupRoutes creates a new mux and registers all routes
func setupRoutes(store store.Store) *http.ServeMux {
	h := handler.NewHandler(store)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleCafe)
	mux.HandleFunc("PUT /cafes", middleware.WithAuth(h.HandleCreateCafe))
	mux.HandleFunc("GET /cafes", h.HandleGetCafes)
	mux.HandleFunc("GET /cafes/{id}", h.HandleGetCafe)

	mux.HandleFunc("POST /register", h.HandleRegisterUser)
	mux.HandleFunc("POST /login", h.HandleLogin)
	return mux
}

func handleCafe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're looking for coffee ☕ \n")
}
