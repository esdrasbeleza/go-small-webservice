package main

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoNotesDAO struct {
	collection *mgo.Collection
}

func CreateNotesMongoDao(collection *mgo.Collection) *MongoNotesDAO {
	return &MongoNotesDAO{collection: collection}
}

func (n MongoNotesDAO) GetAllNotes() []Note {
	result := []Note{}
	n.collection.Find(nil).All(&result)
	return result
}

func (n MongoNotesDAO) GetNoteById(noteId string) (Note, error) {
	result := Note{}
	var err error

	if bson.IsObjectIdHex(noteId) {
		bsonId := bson.ObjectIdHex(noteId)
		err = n.collection.FindId(bsonId).One(&result)
	} else {
		err = errors.New("not found")
	}
	return result, err
}

func (n MongoNotesDAO) StoreNote(note *Note) error {
	return n.collection.Insert(note)
}

func (n MongoNotesDAO) resetDatabase() {
	n.collection.RemoveAll(nil)

	notes := []*Note{
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
