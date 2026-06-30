package main

import (
	"fmt"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	ContentType := r.Header.Get("Content-Type")

	if ContentType != "application/x-www-form-urlencoded" {
		http.Error(w, "only application/x-www-form-urlencoded is supported", http.StatusUnsupportedMediaType)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "malformed form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("username")
	lang := r.FormValue("language")

	if name == "" || lang == "" {
		http.Error(w, "missing username", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Hello [%s], you are coding in [%s]!", name, lang)
	}

}
