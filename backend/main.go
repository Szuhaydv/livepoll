package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"path/filepath"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filepath.Join("..", "public", "index.html"))
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user1"
	password = "1234"
	dbname   = "livepoll_test"
)

func main() {
	// creating DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	fs := http.FileServer(http.Dir(filepath.Join("..", "public")))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Serve the root route with testHandler
	http.HandleFunc("/", testHandler)

	fmt.Println("Running server...")
	err = http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("Error running server")
	}
}
