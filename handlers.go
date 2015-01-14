package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	found := false
	for _, note := range notes {
		if strconv.Itoa(note.Id) == noteId {
			json.NewEncoder(w).Encode(note)
			found = true
		}
	}

	if !found {
		w.WriteHeader(404)
	}
}
