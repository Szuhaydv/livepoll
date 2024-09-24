package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filepath.Join("..", "public", "index.html"))
}

func main() {
	fs := http.FileServer(http.Dir(filepath.Join("..", "public")))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Serve the root route with testHandler
	http.HandleFunc("/", testHandler)

	fmt.Println("Running server...")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("Error running server")
	}
}
