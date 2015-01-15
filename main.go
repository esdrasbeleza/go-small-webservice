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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/note", ListNotes)
	router.HandleFunc("/note/{noteId}", ShowNote)

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

	notes := []Note{
		CreateNote("Things to buy", "Eggs, ham, cheese, beer"),
		CreateNote("Important URL", "http://www.esdrasbeleza.com"),
	}

	for index, note := range notes {
		log.Printf("Inserting note %d\n", index)
		collection.Insert(&note)
	}
}
