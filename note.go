package main

import "time"

type Note struct {
	Id      int       `json: "id"`
	Title   string    `json: "title"`
	Text    string    `json: "text"`
	Added   time.Time `json: "added"`
	Updated time.Time `json: "updated"`
}

var lastId int = 0

func CreateNote(title, text string) Note {
	// OMG an anonymous function to increment id after we return a new note
	defer func() { lastId++ }()

	return Note{Title: title, Text: text, Added: time.Now(), Updated: time.Now(), Id: lastId}
}

// Create some notes
var notes []Note = []Note{
	CreateNote("Things to buy", "Eggs, ham, cheese, beer"),
	CreateNote("Important URL", "http://www.esdrasbeleza.com"),
}
