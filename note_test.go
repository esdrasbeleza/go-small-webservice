package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyNote(t *testing.T) {
	note := NewEmptyNote()

	assert.Equal(t, len(note.Title), 0, "Title should be empty")
	assert.Equal(t, len(note.Text), 0, "Text should be empty")
	assert.NotNil(t, note.Id)
}

func TestNewNote(t *testing.T) {
	title := "My title"
	text := "My text!"
	note := NewNote(title, text)

	assert.Equal(t, title, note.Title, "Title is wrong")
	assert.Equal(t, text, note.Text, "Text is wrong")
	assert.NotNil(t, note.Id)
}
