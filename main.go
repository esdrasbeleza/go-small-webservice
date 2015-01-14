package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/notes", ListNotes)
	router.HandleFunc("/note/{noteId}", ShowNote)

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! :D")
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List notes!")
}

func ShowNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	fmt.Fprintf(w, "Some note %s", noteId)
}
