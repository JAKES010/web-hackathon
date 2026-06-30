package main

import (
	"fmt"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.HandleFunc("/method-inspector", InspectorHandler)
	mainMux.HandleFunc("/echo", EchoHandler)
	mainMux.HandleFunc("/headers", HeadersHandler)
	mainMux.HandleFunc("/form", FormHandler)
	mainMux.HandleFunc("/status", StatusHandler)

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	apiMux.HandleFunc("/v1/ping", PingHandler)
	apiMux.HandleFunc("/v1/greet", GreetHandler)

	fmt.Println("server is running on port http://localhost:8080 ..............")
	http.ListenAndServe(":8080", mainMux)
}
