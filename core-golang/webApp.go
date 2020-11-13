package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go is cool")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("localhost:8000", nil)
}
