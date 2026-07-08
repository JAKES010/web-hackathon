package main

import (
	"fmt"
	"net/http"
)

func HandleLagacy(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func v2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to version 2")
}
