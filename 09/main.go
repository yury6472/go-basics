package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Привет Интернет")
	})
	// Use port 9090
	http.ListenAndServe(":9090", nil)
}
