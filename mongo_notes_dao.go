package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoNotesDAO struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func CreateNotesMongoDao() MongoNotesDAO {
	session := createDatabaseSession()
	collection := session.DB("notes").C("notes")
	return MongoNotesDAO{session: session, collection: collection}
}

func createDatabaseSession() *mgo.Session {
	session, error := mgo.Dial("127.0.0.1")

	if error != nil {
		panic(error)
	}

	return session
}

func (n MongoNotesDAO) GetAllNotes() []Note {
	result := []Note{}
	n.collection.Find(nil).All(&result)
	return result
}

func (n MongoNotesDAO) GetNoteById(noteId string) (Note, error) {
	query := bson.M{"_id": bson.ObjectIdHex(noteId)}
	result := Note{}
	err := n.collection.Find(query).One(&result)

	return result, err
}

func (n MongoNotesDAO) StoreNote(note Note) error {
	return n.collection.Insert(note)
}

func (n MongoNotesDAO) Close() {
	n.session.Close()
}

func (n MongoNotesDAO) resetDatabase() {
	n.collection.RemoveAll(nil)

	notes := []Note{
		NewNote("Things to buy", "Eggs, ham, cheese, beer"),
		NewNote("Important URL", "http://www.esdrasbeleza.com"),
	}

	for _, note := range notes {
		log.Printf("Inserting note %s\n", note.Id)
		err := n.collection.Insert(note)

		if err != nil {
			log.Println(err)
		}
	}
}
