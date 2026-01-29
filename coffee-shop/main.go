package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/tulara/coffeeshop/handler"
	"github.com/tulara/coffeeshop/middleware"
	"github.com/tulara/coffeeshop/store"
)

func main() {
	store := store.NewMemoryStore()
	mux := setupRoutes(store)

	// explicit server so we can run it on its own routine
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//difference between channel and context?
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		// what happens if we don't filter out http.ErrServerClosed
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	fmt.Println("Server started on :8080")
	fmt.Println("Press Ctrl+C or send SIGTERM to shut down")

	<-ctx.Done()
	fmt.Println("Shutting down server gracefully...")

	shutdownCtx, timeout := context.WithTimeout(context.Background(), 30*time.Second)
	defer timeout()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error shutting down: %v", err)
	}
	fmt.Println("Server stopped gracefully")
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
