package main

import "testing"

func TestEmptyNote(t *testing.T) {
	note := NewEmptyNote()

	if len(note.Title) > 0 {
		t.Error("Text should be empty")
	} else if len(note.Text) > 0 {
		t.Error("Text should be empty")
	}
}
