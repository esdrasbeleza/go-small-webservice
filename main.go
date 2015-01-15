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

	// FIXME: routes are NOT in the right REST format
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/note", ListNotes)
	router.HandleFunc("/note/id/{noteId}", ShowNote)
	router.HandleFunc("/note/create", AddNote).Methods("POST")

	CreateSession()
	InsertNotesIntoDatabase()

	log.Println("Starting server!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func CreateSession() {
	session = CreateDatabaseSession()
}

func InsertNotesIntoDatabase() {
	collection = session.DB("notes").C("notes")
	collection.RemoveAll(nil)

	notes := []Note{
		CreateNote("Things to buy", "Eggs, ham, cheese, beer"),
		CreateNote("Important URL", "http://www.esdrasbeleza.com"),
	}

	for _, note := range notes {
		log.Printf("Inserting note %s\n", note.Id)
		err := collection.Insert(note)
		if err != nil {
			log.Println(err)
		}
	}
}
