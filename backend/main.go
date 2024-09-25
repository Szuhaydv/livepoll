package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
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

func executeQuery(db *sql.DB, filepath string) error {
	query, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

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

	// initializing tables
	err = tablesInit(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized tables")

	// TEST INSERT
	err = insertPoll(db, "150 second", "Test poll insert")
	if err != nil {
		fmt.Println("Error doing test insert")
	}

	// ROUTING
	router := mux.NewRouter()

	// serving assets
	fileServer := http.FileServer(http.Dir(filepath.Join("..", "public")))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	// cathcall route serves Svelte SPA frontend

	router.PathPrefix("/").HandlerFunc(testHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:7777",
	}
	fmt.Println("Running server...")
	log.Fatal(srv.ListenAndServe())
}
