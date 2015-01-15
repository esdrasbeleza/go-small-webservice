package main

import "time"

type Note struct {
	Id      int       `bson:"id" json:"id"`
	Title   string    `bson:"title" json:"title"`
	Text    string    `bson:"text" json:"text"`
	Added   time.Time `bson:"added" json:"added"`
	Updated time.Time `bson:"updated" json:"updated"`
}

var lastId int = 0

func CreateNote(title, text string) Note {
	// OMG an anonymous function to increment id after we return a new note
	defer func() { lastId++ }()

	return Note{Title: title, Text: text, Added: time.Now(), Updated: time.Now(), Id: lastId}
}
