package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filepath.Join("..", "public", "index.html"))
}

type Poll struct {
	Title    string   `json:"title"`
	Duration *string  `json:"duration,omitempty"`
	Options  []Option `json:"options"`
}

type Option struct {
	Name string `json:"name"`
}

func handlePollCreate(w http.ResponseWriter, r *http.Request) {
	var poll Poll

	err := json.NewDecoder(r.Body).Decode(&poll)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(poll.Options) < 2 {
		http.Error(w, "Insufficient amount of options specified (min. 2)", http.StatusBadRequest)
		return
	}
	if poll.Duration == nil {
		defaultDuration := "120 second"
		poll.Duration = &defaultDuration
	}

	pollID, createErr := insertPoll(db, poll)
	if createErr != nil {
		http.Error(w, fmt.Sprintf("Failed to insert poll %v", createErr), http.StatusInternalServerError)
		return
	}
	err = insertOptions(db, pollID, poll.Options)
	if err != nil {
		http.Error(w, "Failed to insert options", http.StatusInternalServerError)
		return
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user1"
	password = "1234"
	dbname   = "livepoll_test"
)

func executeQuery(filepath string) error {
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

var db *sql.DB

func main() {
	// creating DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
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
	err = tablesInit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized tables")

	// ROUTING
	router := mux.NewRouter()

	// serving assets
	fileServer := http.FileServer(http.Dir(filepath.Join("..", "public")))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	// cathcall route serves Svelte SPA frontend

	router.HandleFunc("/create-poll", handlePollCreate).Methods("POST")
	router.PathPrefix("/").HandlerFunc(testHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:7777",
	}
	fmt.Println("Running server...")
	log.Fatal(srv.ListenAndServe())
}
