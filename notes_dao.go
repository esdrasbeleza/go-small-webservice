package main

type NotesDao interface {
	GetAllNotes() []Note

	GetNoteById(noteId string) (Note, error)

	StoreNote(note Note) error
}
