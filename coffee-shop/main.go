package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleCafe)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleCafe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're looking for coffee ☕ \n")
}
