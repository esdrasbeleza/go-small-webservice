package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Note struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `bson:"title" json:"title"`
	Text    string        `bson:"text" json:"text"`
	Added   time.Time     `bson:"added" json:"added"`
	Updated time.Time     `bson:"updated" json:"updated"`
}

func NewEmptyNote() *Note {
	return &Note{Added: time.Now(), Updated: time.Now(), Id: bson.NewObjectId()}
}

func NewNote(title, text string) *Note {
	return &Note{Title: title, Text: text, Added: time.Now(), Updated: time.Now(), Id: bson.NewObjectId()}
}
