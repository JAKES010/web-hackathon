package main

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
		fmt.Fprintf(w, "Greetings, [%s]!", name)
	} else {
		fmt.Fprintf(w, "Greetings, [%s]!", name)
	}
}
