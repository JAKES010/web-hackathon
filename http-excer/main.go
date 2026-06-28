package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/method-inspector", InspectorHandler)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/headers", headers)
	fmt.Println("server is running on port http://localhost:8080 ..............")
	http.ListenAndServe(":8080", nil)
}
