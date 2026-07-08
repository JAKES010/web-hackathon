package main

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-API-Key")

	if token == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if token != "secret123" {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome to your dashboard")
}
