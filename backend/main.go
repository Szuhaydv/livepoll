package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/lib/pq"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filepath.Join("..", "public", "index.html"))
}

type Poll struct {
	Title     string    `json:"title"`
	Duration  *string   `json:"duration,omitempty"`
	Options   []Option  `json:"options"`
	CreatedAt time.Time `json:"created_at"`
}

type Option struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Vote struct {
	VoterID  string    `json:"voter_id"`
	OptionID uuid.UUID `json:"option_id"`
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

	pollID, createErr := insertPoll(poll)
	if createErr != nil {
		http.Error(w, fmt.Sprintf("Failed to insert poll %v", createErr), http.StatusInternalServerError)
		return
	}
	err = insertOptions(pollID, poll.Options)
	if err != nil {
		http.Error(w, "Failed to insert options", http.StatusInternalServerError)
		return
	}

	// return the poll ID
	response := struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: pollID,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode pollID json", http.StatusInternalServerError)
	}
}

func handleGetPoll(w http.ResponseWriter, r *http.Request) {
	pollID := mux.Vars(r)["pollID"]
	if pollID == "" {
		http.Error(w, "Missing path param (pollID) in request", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(pollID)
	if err != nil {
		http.Error(w, "Specified id query paramter not in acceptable uuid format", http.StatusBadRequest)
		return
	}
	poll, selectErr := getPoll(id)
	if selectErr != nil {
		http.Error(w, selectErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(poll)
}

func handleVote(w http.ResponseWriter, r *http.Request) {
	var vote Vote

	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ip := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
	vote.VoterID = ip

	err = updateVotes(vote)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating votes: %v", err), http.StatusInternalServerError)
		return
	}
}

func listenToNotifications(notifyChan chan<- string, pollID string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	channelName := "poll_" + pollID

	listener := pq.NewListener(psqlInfo, 10*time.Second, time.Minute, nil)

	if err := listener.Listen(channelName); err != nil {
		return err
	}

	log.Println("Listening for PostgreSQL notifications...")
	for {
		select {
		case notification := <-listener.Notify:
			if notification != nil {
				notifyChan <- notification.Extra
			}
		case <-time.After(90 * time.Second):
			go listener.Ping()
		}
	}
}

func handleSSE(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Validation
	pollID := mux.Vars(r)["pollID"]
	_, err := uuid.Parse(pollID)
	if err != nil {
		http.Error(w, "Not valid uuid", http.StatusBadRequest)
		return
	}
	err = executeQuery("./queries/select_exists_poll.sql", false, pollID)
	if err != nil {
		http.Error(w, "Poll with specified ID does not exist", http.StatusBadRequest)
		return
	}

	newChan := make(chan string)
	go listenToNotifications(newChan, pollID)

	// Listen for notifications on the notificationChannel and send them to the client
	for notification := range newChan {
		fmt.Fprintf(w, "data: %s\n\n", notification)
		flusher.Flush() // Push data to the client
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user1"
	password = "1234"
	dbname   = "livepoll_test"
)

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

	// creating function and trigger for realtime
	executeQuery("./queries/function_create.sql", true)
	executeQuery("./queries/trigger_create.sql", true)

	// ROUTING
	router := mux.NewRouter()

	// serving assets
	fileServer := http.FileServer(http.Dir(filepath.Join("..", "public")))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	// cathcall route serves Svelte SPA frontend

	router.HandleFunc("/create-poll", handlePollCreate).Methods("POST")
	router.HandleFunc("/polls/{pollID}", handleGetPoll).Methods("GET")
	router.HandleFunc("/vote", handleVote).Methods("POST")
	router.HandleFunc("/results/{pollID}", handleSSE).Methods("GET")
	router.PathPrefix("/").HandlerFunc(testHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:7777",
	}
	fmt.Println("Running server...")
	log.Fatal(srv.ListenAndServe())
}
