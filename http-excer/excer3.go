package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Token received:%s", token)
	}
	ContentType := r.Header.Get("Content-Type")
	if ContentType == "" {
		http.Error(w, "Content-Type not provided", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Content-Type: %s", ContentType)
	}
}
