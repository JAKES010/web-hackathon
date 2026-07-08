package main

import (
	"fmt"
	"io"
	"net/http"
)

func CountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Send a POST request with text to count words")
	case http.MethodPost:
		contents, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "badRequest", http.StatusBadRequest)
			return
		}
		text := string(contents)
		fmt.Fprintf(w, "%d", len(text))
	default:
		http.Error(w, "Method NOt Allowed", http.StatusMethodNotAllowed)
		return
	}

}
