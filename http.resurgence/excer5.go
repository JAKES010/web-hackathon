package main

import (
	"fmt"
	"net/http"
)

func AgentHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")

	if user == "" {
		return
	}

	fmt.Fprintf(w, "You are visiting us using: %s", user)
}
