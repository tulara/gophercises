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

	// explicit server so we can run it on its own routine
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// CLAUDE -

	// // Channel to listen for OS signals
	// sigChan := make(chan os.Signal, 1)
	// // Notify sigChan when SIGINT (Ctrl+C) or SIGTERM is received
	// signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// // Run server in a goroutine so it doesn't block
	// go func() {
	// 	fmt.Println("Server listening on :8080")
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("Server failed to start: %v", err)
	// 	}
	// }()

	// // Block here until signal is received
	// sig := <-sigChan
	// fmt.Printf("\nReceived signal: %v. Shutting down gracefully...\n", sig)

	// // Create context with timeout for shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// // Attempt graceful shutdown
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server forced to shutdown: %v", err)
	// }

	// fmt.Println("Server stopped gracefully")

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
