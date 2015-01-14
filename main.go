package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/note", ListNotes)
	router.HandleFunc("/note/{noteId}", ShowNote)

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
