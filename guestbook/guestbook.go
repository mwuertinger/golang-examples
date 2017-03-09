package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type entry struct {
	Author  string
	Message string
}

var (
	entries = []entry{}
	mu      sync.Mutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/entry", handleEntryGet).Methods("GET")
	r.HandleFunc("/entry", handleEntryPost).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("www")))

	panic(http.ListenAndServe(":8080", r))
}

func handleEntryGet(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	out, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		log.Printf("unable to marshall json: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func handleEntryPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entry := entry{}
	err := decoder.Decode(&entry)
	if err != nil {
		log.Printf("unable to decode json: %v", err)
		w.WriteHeader(500)
		return
	}

	log.Printf("entry = %v", entry)

	mu.Lock()
	entries = append(entries, entry)
	defer mu.Unlock()

	w.WriteHeader(200)
}
