package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func createDatabaseSession() *mgo.Session {
	session, error := mgo.Dial("127.0.0.1")

	if error != nil {
		panic(error)
	}

	return session
}

func resetDatabase() {
	collection.RemoveAll(nil)

	notes := []Note{
		NewNote("Things to buy", "Eggs, ham, cheese, beer"),
		NewNote("Important URL", "http://www.esdrasbeleza.com"),
	}

	for _, note := range notes {
		log.Printf("Inserting note %s\n", note.Id)
		err := collection.Insert(note)

		if err != nil {
			log.Println(err)
		}
	}
}
