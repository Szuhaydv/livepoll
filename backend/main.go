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
	fsBuild := http.FileServer(http.Dir(filepath.Join("..", "public", "build")))
	http.Handle("/build/", http.StripPrefix("/build/", fsBuild))

	fsAssets := http.FileServer(http.Dir(filepath.Join("..", "public", "assets")))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))

	http.HandleFunc("/", testHandler)

	fmt.Println("Running server...")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("Error running server")
	}
}
