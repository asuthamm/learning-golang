// basic-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// this is a middleware
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		log.Println(f)
		log.Println(bar)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {

	http.HandleFunc("/foo", foo) // bypass middleware
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
