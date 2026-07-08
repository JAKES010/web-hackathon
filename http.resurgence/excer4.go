package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalHandler(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	if operation == "" || a == "" || b == "" {
		http.Error(w, "missing query parameters", http.StatusBadRequest)
		return
	}

	num1, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	var result int

	switch operation {
	case "addition", "add":

		result = add(num1, num2)
	case "subtraction", "sub":

		result = subract(num1, num2)
	case "multiply":

		result = multiply(num1, num2)

	default:
		http.Error(w, "invalid operations", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Result: %d", result)
}
func add(a int, b int) int {
	return a + b
}

func subract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
