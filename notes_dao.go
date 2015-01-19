package main

import (
	"gopkg.in/mgo.v2/bson"
)

type NotesDao interface {
	GetAllNotes() []Note

	GetNoteById(noteId string) (Note, error)

	StoreNote(note Note) error
}

type MongoNotesDAO struct {
}

func CreateNotesMongoDao() MongoNotesDAO {
	return MongoNotesDAO{}
}

func (n MongoNotesDAO) GetAllNotes() []Note {
	result := []Note{}
	collection.Find(nil).All(&result)
	return result
}

func (n MongoNotesDAO) GetNoteById(noteId string) (Note, error) {
	query := bson.M{"_id": bson.ObjectIdHex(noteId)}
	result := Note{}
	err := collection.Find(query).One(&result)

	return result, err
}

func (n MongoNotesDAO) StoreNote(note Note) error {
	return collection.Insert(note)
}
