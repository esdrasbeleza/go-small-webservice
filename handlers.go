package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func ShowNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	fmt.Fprintf(w, "Some note %s", noteId)
}
