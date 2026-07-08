package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/count", CountHandler)
	http.HandleFunc("/calculate", CalHandler)
	http.HandleFunc("/agent", AgentHandler)
	http.HandleFunc("/dashboard", DashboardHandler)
	http.HandleFunc("/legacy", HandleLagacy)
	http.HandleFunc("/v2", v2)
	fmt.Println("Server is running on port 8080.........")
	http.ListenAndServe(":8080", nil)
}
