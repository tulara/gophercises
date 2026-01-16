package main

import (
	"fmt"
	"net/http"

	"github.com/tulara/coffeeshop/handler"
	"github.com/tulara/coffeeshop/store"
)

func main() {
	store := store.NewMemoryStore()
	handler := handler.NewHandler(store)

	http.HandleFunc("/", handleCafe)
	http.HandleFunc("/cafes", handler.HandleCreateCafe)
	http.HandleFunc("/cafes/{id}", handler.HandleGetCafe)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleCafe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're looking for coffee ☕ \n")
}
