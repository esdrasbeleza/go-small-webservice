package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao MongoNotesDAO

func main() {
	defer dao.Close()

	router := setupRouter()

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setupRouter() *mux.Router {
	dao = CreateNotesMongoDao()
	dao.resetDatabase()
	handler := CreateHandler(dao)

	// FIXME: routes are NOT in the right REST format
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/note", handler.ListNotes)
	router.HandleFunc("/note/id/{noteId}", handler.GetNote)
	router.HandleFunc("/note/create", handler.RegisterNote).Methods("POST")
	return router
}
