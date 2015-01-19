package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	dao     MongoNotesDAO
	session *mgo.Session
)

func main() {
	defer session.Close()

	router := setupRouter()

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setupRouter() *mux.Router {
	session := createDatabaseSession()
	collection := session.DB("notes").C("notes")

	dao = CreateNotesMongoDao(collection)
	dao.resetDatabase()
	handler := CreateHandler(dao)

	// FIXME: routes are NOT in the right REST format
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/note", handler.ListNotes)
	router.HandleFunc("/note/id/{noteId}", handler.GetNote)
	router.HandleFunc("/note/create", handler.RegisterNote).Methods("POST")
	return router
}

func createDatabaseSession() *mgo.Session {
	session, error := mgo.Dial("127.0.0.1")

	if error != nil {
		panic(error)
	}

	return session
}
