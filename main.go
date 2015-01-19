package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func main() {
	defer session.Close()

	router := setupRouter()
	setupDatabase()

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setupRouter() *mux.Router {
	// FIXME: routes are NOT in the right REST format
	dao := CreateNotesMongoDao()
	handler := CreateHandler(dao)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/note", handler.ListNotes)
	router.HandleFunc("/note/id/{noteId}", handler.GetNote)
	router.HandleFunc("/note/create", handler.RegisterNote).Methods("POST")
	return router
}

func setupDatabase() {
	session = createDatabaseSession()
	collection = session.DB("notes").C("notes")
	resetDatabase()
}
