package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../public/build")))

	fmt.Println("Running server...")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("Error running server")
	}
}
