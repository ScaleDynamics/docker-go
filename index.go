package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "World"
	}
	version := runtime.Version()[2:]
	fmt.Fprintf(w, "Hello %s from Go v%s", name, version)
}

func main() {
	port := 8080
	fmt.Printf("HTTP server listening on port %d", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
