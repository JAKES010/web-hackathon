package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	context, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}

	if context < 100 || context > 599 {
		http.Error(w, "code must be a valid HTTP status code (100 to 599)", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(context)
		fmt.Fprintf(w, "Responding with status [%d]", context)
	}

}
