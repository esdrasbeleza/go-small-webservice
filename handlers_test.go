package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var handler Handler

type FakeDao struct {
	notes []Note
}

func (f *FakeDao) createNotes() {
	emptyNote := NewEmptyNote()
	newNote := NewNote("My new note", "Some text")
	f.notes = []Note{emptyNote, newNote}
}

func (f FakeDao) GetAllNotes() []Note {
	f.createNotes()
	return f.notes
}

func (f FakeDao) GetNoteById(noteId string) (Note, error) {
	f.createNotes()
	switch noteId {
	case "0":
		return f.notes[0], nil
	case "1":
		return f.notes[1], nil
	default:
		return Note{}, errors.New("Not found")
	}
}

func (f FakeDao) StoreNote(note Note) error {
	return nil
}

func createTestHandler() {
	dao := FakeDao{}
	handler = CreateHandler(dao)
}

func TestListNotes(t *testing.T) {
	createTestHandler()

	req, _ := http.NewRequest("GET", "addressDoesNotMatter", nil)
	w := httptest.NewRecorder()
	handler.ListNotes(w, req)

	assert.NotEqual(t, 0, len(w.Body.String()), "Body must not be empty")
	assert.Equal(t, 200, w.Code, "Code must be 200 but was "+strconv.Itoa(w.Code))
}
