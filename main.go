package main

import (
	"log" //
	"net/http"
	"orness/api"

	"github.com/gorilla/mux"
)

func main() {
	req := mux.NewRouter()
	req.HandleFunc("/notes", api.GetNotes).Methods("GET") // Get method for retreiving the data from the memory
	req.HandleFunc("/notes", api.AddNote).Methods("POST") // Post method for writing the data in to the memory
	log.Fatal(http.ListenAndServe(":4000", req))          // Writes out Log to console if gets any error
}
