package main

import (
	"fmt"
	"io"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "badRequest", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "%s", content)

}
